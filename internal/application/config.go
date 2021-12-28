package application

type Config struct {
	BlePeripheralName            string
	BlePeripheralNameIsNotUnique bool
}

func (conf *Config) GetBlePeripheralName() string {
	if conf.BlePeripheralName == "" {
		return "Brooses-Awesome-Peripheral"
	} else {
		return conf.BlePeripheralName
	}
}

func (conf *Config) GetBlePeripheralNameIsNotUnique() bool {
	return conf.BlePeripheralNameIsNotUnique
}
