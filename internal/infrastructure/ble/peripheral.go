package ble

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Peripheral struct {
	server       *server
	shutdown     bool
	shutdownLock *sync.Mutex
	shutdowns    chan struct{}
	wg           *sync.WaitGroup
}

func NewPeripheral(cfg Config, bus Bus) (*Peripheral, error) {

	log.Println("ble, NewPeripheral: creating")

	shutdowns := make(chan struct{})
	shutdownLock := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	srv, err := newServer(cfg, bus, shutdowns, wg)
	if err != nil {
		return nil, err
	}

	log.Println("ble, NewPeripheral: created successfully")

	return &Peripheral{
		server:       srv,
		shutdown:     true,
		shutdowns:    shutdowns,
		shutdownLock: shutdownLock,
		wg:           wg,
	}, nil
}

func (p *Peripheral) Startup() error {

	log.Println("ble, Startup: starting")

	p.shutdownLock.Lock()
	defer p.shutdownLock.Unlock()

	if p.shutdown == false {
		return errors.New("ble, Startup: Tried to startup server twice")
	}

	err := p.server.startup()
	if err != nil {
		return err
	}

	p.shutdown = false

	log.Println("ble, Startup: started")

	return nil
}

func (p *Peripheral) Shutdown() {

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
	err := p.server.shutdown()
	if err != nil {
		log.Println(fmt.Sprintf("ble, Shutdown: error closing server, %s", err.Error()))
	}

	log.Printf("ble, Shutdown: probably done")
}
