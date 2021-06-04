package serial

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.bug.st/serial"

	"github.com/scutrobotlab/asuwave/datautil"
	"github.com/scutrobotlab/asuwave/logger"
	"github.com/scutrobotlab/asuwave/variable"
)

type T struct {
	Name string
	Mode serial.Mode
	Port serial.Port
}

var SerialCur = T{
	Name: "",
	Mode: serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	},
	Port: nil,
}

var chOp = make(chan bool)
var chEd = make(chan bool)
var chRx = make(chan []byte, 10)
var chTx = make(chan []byte, 10)

const testPortName = "Test port"

// Find ports
func Find() []string {
	var ports []string
	ports = append(ports, testPortName)

	tmp, err := serial.GetPortsList()
	if err != nil {
		logger.Log.Println("Serial ports errors!")
	}
	if len(tmp) == 0 {
		logger.Log.Println("No serial ports found!")
	}
	for _, port := range tmp {
		if strings.Contains(port, "USB") || strings.Contains(port, "ACM") || strings.Contains(port, "COM") || strings.Contains(port, "tty.usb") {
			ports = append(ports, port)
		}
	}
	return ports
}

// Open serial port
func Open(name string) error {
	SerialCur.Name = name

	if name == testPortName {
		SerialCur.Port = newTestPort()
		chOp <- true
		return nil
	}

	var err error
	SerialCur.Port, err = serial.Open(name, &SerialCur.Mode)
	if err != nil {
		SerialCur.Name = ""
		return err
	}
	chOp <- true
	return nil
}

// Close serial port
func Close() error {
	if SerialCur.Name == "" {
		return errors.New("empty serial port")
	}

	err := SerialCur.Port.Close()
	if err != nil {
		return err
	}
	SerialCur.Name = ""
	chEd <- true
	return nil
}

//Transmit data
func Transmit(data []byte) error {
	_, err := SerialCur.Port.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//Receive data
func Receive(buff []byte) ([]byte, error) {
	n, err := SerialCur.Port.Read(buff)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return buff[0:0], nil
	}
	return buff[:n], nil
}

func SendCmd(act uint8, v variable.T) error {
	if SerialCur.Port == nil || SerialCur.Name == "" {
		return errors.New("no serial port")
	}

	data := datautil.MakeCmd(act, &v)

	chTx <- data
	return nil
}

func GrReceive() {
	buff := make([]byte, 200)
	for {
		<-chOp
		for _, v := range variable.ToRead.Variables {
			SendCmd(datautil.ActModeSubscribe, v)
			time.Sleep(10 * time.Millisecond)
		}
	Loop:
		for {
			select {
			case <-chEd:
				break Loop
			default:
				b, err := Receive(buff)
				if err != nil {
					logger.Log.Println("GrReceive error:", err)
				}
				chRx <- b
				time.Sleep(5 * time.Millisecond)
			}
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func GrTransmit() {
	for {
		b := <-chTx
		err := Transmit(b)
		if err != nil {
			logger.Log.Println("GrTransmit error: ", err)
		}
		time.Sleep(3 * time.Millisecond)
	}
}

func GrRxPrase(c chan string) {
	var rxBuff []byte
	var x variable.ListChartT
	for {
		rx := <-chRx
		rxBuff = append(rxBuff, rx...)

		startIdx, endIdx := datautil.FindValidPart(rxBuff)

		logger.Log.Printf("rxBuff: %#v\n", rxBuff)
		logger.Log.Printf("startIdx: %d, endIdx: %d\n", startIdx, endIdx)

		buff := rxBuff[startIdx:endIdx]

		x.Variables = nil
		datautil.MakeChartPack(&x, &variable.ToRead, buff)
		if len(x.Variables) != 0 {
			b, _ := json.Marshal(x)
			c <- string(b)
		}

		if endIdx >= len(rxBuff) {
			rxBuff = nil
		} else {
			rxBuff = rxBuff[endIdx:]
		}
	}
}
