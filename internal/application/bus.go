package application

import "github.com/broose-goose/Get-This-Pi-On-The-WiFi/internal/infrastructure/ble"

type applicationBus interface {
	Shutdown()
	ble.Bus
}
