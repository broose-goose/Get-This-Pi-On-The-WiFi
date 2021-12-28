package bus

import "log"

func (b *bus) HandleBleConnectionOpen() (txChan chan []byte) {
	return make(chan []byte, 10)
}

func (b *bus) HandleBleMessageReceived(by []byte) {
	log.Println(string(by))
}

func (b *bus) HandleBleTryClose() {
	log.Println("closing")
}
