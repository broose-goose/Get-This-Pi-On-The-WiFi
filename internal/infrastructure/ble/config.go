package ble

type Config interface {
	GetBlePeripheralName() string
	GetBlePeripheralNameIsNotUnique() bool
}
