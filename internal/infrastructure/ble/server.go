package ble

import (
	"errors"
	"fmt"
	"sync"
	"tinygo.org/x/bluetooth"
)

type server struct {
	shutdowns chan struct{}
	wg *sync.WaitGroup
	adapter bluetooth.Adapter
}

func newServer(shutdowns chan struct{}, wg *sync.WaitGroup) (*server, error) {
	bleAdapter := bluetooth.DefaultAdapter
	err := bleAdapter.Enable()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("ble, newServer: failed to enable ble; %s", err.Error()))
	}
}