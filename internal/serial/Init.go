package serial

func init() {
	go GrReceive()
	go GrTransmit()
	go GrRxPrase()
}
