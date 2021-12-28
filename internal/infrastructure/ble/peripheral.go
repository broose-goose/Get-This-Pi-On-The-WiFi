package ble

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type peripheral struct {
	shutdown bool
	shutdownLock *sync.Mutex
	shutdowns chan struct{}
	wg *sync.WaitGroup
}

func NewPeripheral() (*peripheral, error) {

	log.Println("ble, NewPeripheral: creating")

	shutdowns := make(chan struct{})
	shutdownLock := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	log.Println("ble, NewPeripheral: created successfully")

	return &peripheral{
		shutdown: true,
		shutdowns: shutdowns,
		shutdownLock: shutdownLock,
		wg: wg,
	}, nil
}

func (p *peripheral) Startup() error {

	log.Println("ble, Startup: starting")

	p.shutdownLock.Lock()
	defer p.shutdownLock.Unlock()

	if p.shutdown == false {
		return errors.New("ble, Startup: Tried to startup server twice")
	}

	go func () {
		var err error = nil
		if err != nil {
			log.Println(fmt.Sprintf("ble: reported err %s", err))
			// turn off bluetooth
		}
	}()

	p.shutdown = false

	log.Println("ble, Startup: started")

	return nil
}

func (p *peripheral) Shutdown() {

	log.Printf("ble, Shutdown: Shutting down")

	p.shutdownLock.Lock()
	defer p.shutdownLock.Unlock()
	if p.shutdown {
		return
	}
	p.shutdown = true

	close(p.shutdowns)
	p.wg.Wait()
	// maybe closeout peripheral here?
	var err error = nil
	if err != nil {
		log.Println(fmt.Sprintf("ble, Shutdown: error closing server, %s", err.Error()))
	}

	log.Printf("ble, Shutdown: probably done")
}
