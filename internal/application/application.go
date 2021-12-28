package application

import (
	"github.com/broose-goose/Get-This-Pi-On-The-WiFi/internal/infrastructure/ble"
	"github.com/broose-goose/Get-This-Pi-On-The-WiFi/internal/infrastructure/bus"
	"log"
	"sync"
)

type Application struct {
	serviceBus    applicationBus
	blePeripheral *ble.Peripheral
	shutdown      bool
	shutdownLock  sync.Mutex
}

func NewApplication(conf *Config) (*Application, error) {
	log.Println("Application, NewApplication: creating")

	/* create application instance */
	app := &Application{
		shutdown: true,
	}

	/* create bus */
	app.serviceBus = bus.NewBus()

	/* create servers */
	blePeripheral, err := ble.NewPeripheral(conf, app.serviceBus)
	if err != nil {
		log.Fatalf("Application, NewApplication: failed to initialize ble peripheral")
		return nil, err
	}
	app.blePeripheral = blePeripheral

	return app, err
}

func (app *Application) Startup() error {

	log.Println("Application, Startup: starting")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown == false {
		return nil
	}

	app.shutdown = false

	err := app.blePeripheral.Startup()
	if err != nil {
		return err
	}

	return nil
}

func (app *Application) Shutdown() error {

	log.Println("Application, Shutdown: shutting down")

	app.shutdownLock.Lock()
	defer app.shutdownLock.Unlock()
	if app.shutdown {
		return nil
	}
	app.shutdown = true

	app.blePeripheral.Shutdown()

	return nil
}
