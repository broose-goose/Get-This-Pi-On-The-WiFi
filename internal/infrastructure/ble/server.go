package ble

import (
	"fmt"
	"log"
	"sync"
	"tinygo.org/x/bluetooth"
)

type server struct {
	adapter       *bluetooth.Adapter
	advertisement *bluetooth.Advertisement
	bus           Bus
	shutdowns     chan struct{}
	wg            *sync.WaitGroup
}

func newServer(cfg Config, bus Bus, shutdowns chan struct{}, wg *sync.WaitGroup) (*server, error) {

	log.Println("ble, newServer: creating")

	bleAdapter := bluetooth.DefaultAdapter
	err := bleAdapter.Enable()
	if err != nil {
		log.Println(fmt.Sprintf("ble, newServer: failed to enable ble; %s", err.Error()))
		return nil, err
	}

	bleAdv := bleAdapter.DefaultAdvertisement()

	advName := cfg.GetBlePeripheralName()
	if cfg.GetBlePeripheralNameIsNotUnique() {
		advName += "-" + randSeq(5)
	}

	err = bleAdv.Configure(bluetooth.AdvertisementOptions{
		LocalName:    advName,
		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDNordicUART},
	})
	if err != nil {
		log.Println(fmt.Sprintf("ble, newServer: failed to setup advertisement; %s", err.Error()))
		return nil, err
	}

	log.Println("ble, newServer: created successfully")

	return &server{
		adapter:       bleAdapter,
		advertisement: bleAdv,
		bus:           bus,
		shutdowns:     shutdowns,
		wg:            wg,
	}, nil
}

func (s *server) startup() error {
	err := s.advertisement.Start()
	if err != nil {
		log.Println(fmt.Sprintf("ble, serverStartup: failed to start advertising; %s", err.Error()))
		return err
	}

	var rxChar bluetooth.Characteristic
	var txChar bluetooth.Characteristic

	var rxBuff []byte

	err = s.adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDNordicUART,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &rxChar,
				UUID:   bluetooth.CharacteristicUUIDUARTRX,
				Flags:  bluetooth.CharacteristicWritePermission | bluetooth.CharacteristicWriteWithoutResponsePermission,
				WriteEvent: func(client bluetooth.Connection, offset int, value []byte) {
					log.Println(value)
					if value[0] == '\n' {
						s.bus.HandleBleMessageReceived(rxBuff)
						rxBuff = []byte{}
					} else {
						rxBuff = append(rxBuff, value...)
					}
				},
			},
			{
				Handle: &txChar,
				UUID:   bluetooth.CharacteristicUUIDUARTTX,
				Flags:  bluetooth.CharacteristicNotifyPermission | bluetooth.CharacteristicReadPermission,
			},
		},
	})
	if err != nil {
		log.Println(fmt.Sprintf("ble, serverStartup: failed to start service; %s", err.Error()))
		return err
	}

	txChan := s.bus.HandleBleConnectionOpen()
	go func(s *server, txChar bluetooth.Characteristic, txChan chan []byte, wg *sync.WaitGroup) {
		for {
			select {
			case msg, ok := <-txChan:
				if !ok {
					log.Println("ble, server: tx chan suddenly closed")
					goto StopTxOut
				}
				for _, c := range msg {
					n, err := txChar.Write([]byte{c})
					if err != nil || n != 0 {
						log.Println("ble, server: unable to write to channel")
						goto StopTxOut
					}
				}
				n, err := txChar.Write([]byte{'\n'})
				if err != nil || n != 0 {
					log.Println("ble, server: unable to write to channel")
					goto StopTxOut
				}
			case _, ok := <-s.shutdowns:
				if !ok {
					goto StopTxOut
				}
			}
		}
	StopTxOut:
		s.bus.HandleBleTryClose()
		wg.Done()
	}(s, txChar, txChan, s.wg)

	s.wg.Add(1)

	return nil
}

func (s *server) shutdown() error {
	// really wish i could disconnect ble here...
	return nil
}
