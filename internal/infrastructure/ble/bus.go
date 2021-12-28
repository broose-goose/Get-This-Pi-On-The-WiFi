package ble

type Bus interface {
	HandleBleConnectionOpen() (txChan chan []byte)
	HandleBleMessageReceived(by []byte)
	HandleBleTryClose()
}
