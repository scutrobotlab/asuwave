package serial

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.bug.st/serial"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/variable"
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

var Chch = make(chan string) // 新图表Json

var chOp = make(chan bool)       // 敞开心扉
var chEd = make(chan bool)       // 沉默不语
var chRx = make(chan []byte, 10) // 来信收讫
var chTx = make(chan []byte, 10) // 去信已至

const testPortName = "Test port"

// Find ports
func Find() []string {
	var ports []string
	ports = append(ports, testPortName)

	tmp, err := serial.GetPortsList()
	if err != nil {
		glog.Errorln("Serial ports errors: ", err.Error())
	}
	if len(tmp) == 0 {
		glog.Infoln("No serial ports found!")
	}
	for _, port := range tmp {
		if strings.Contains(port, "USB") || strings.Contains(port, "ACM") || strings.Contains(port, "COM") || strings.Contains(port, "tty.usb") {
			ports = append(ports, port)
		}
	}
	return ports
}

// Open serial port
func Open(name string, baud int) error {
	SerialCur.Name = name
	SerialCur.Mode.BaudRate = baud

	if name == testPortName {
		SerialCur.Port = newTestPort()
		chOp <- true
		return nil
	}

	var err error
	SerialCur.Port, err = serial.Open(SerialCur.Name, &SerialCur.Mode)
	if err != nil {
		SerialCur.Name = ""
		return err
	}
	glog.Infoln(SerialCur.Name, "Opened.")
	chOp <- true
	return nil
}

// Close serial port
func Close() error {
	if SerialCur.Name == "" {
		return errors.New("serial port had closed")
	}

	err := SerialCur.Port.Close()
	if err != nil {
		return err
	}
	glog.Infoln(SerialCur.Name, "Closed.")
	SerialCur.Name = ""
	chEd <- true
	return nil
}

//Transmit data
func Transmit(data []byte) error {
	glog.V(3).Infoln("serial port write: ", data)
	_, err := SerialCur.Port.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//Receive data
func Receive(buff []byte) ([]byte, error) {
	n, err := SerialCur.Port.Read(buff)
	glog.V(5).Infoln("serial port read: ", n)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return buff[0:0], nil
	}
	return buff[:n], nil
}

var adding = map[variable.CmdT]time.Time{}
var deling = map[variable.CmdT]time.Time{}

func SendWriteCmd(v variable.T) error {
	if SerialCur.Port == nil || SerialCur.Name == "" {
		return errors.New("no serial port")
	}

	glog.Infoln("Send write cmd", v)
	data := variable.MakeWriteCmd(v)
	chTx <- data
	return nil
}

func SendCmd(act variable.ActMode, v variable.CmdT) error {
	if SerialCur.Port == nil || SerialCur.Name == "" {
		return errors.New("no serial port")
	}

	if act == variable.Subscribe {
		if t, ok := adding[v]; ok {
			if time.Since(t) < time.Second {
				glog.V(2).Infoln("Has sent subscribe cmd recently", v)
				return nil
			}
		}
		adding[v] = time.Now()
	} else if act == variable.Unsubscribe {
		if t, ok := deling[v]; ok {
			if time.Since(t) < time.Second {
				glog.V(2).Infoln("Has sent unsubscribe cmd recently", v)
				return nil
			}
		}
		deling[v] = time.Now()
	}

	glog.Infoln("Send cmd", act, v)
	data := variable.MakeCmd(act, v)
	chTx <- data
	return nil
}

func GrReceive() {
	buff := make([]byte, 200)
	for {
		<-chOp
		glog.V(4).Infoln("chOp...")
	Loop:
		for {
			select {
			case <-chEd:
				glog.V(4).Infoln("GrReceive: got chEd...")
				break Loop
			default:
				glog.V(4).Infoln("GrReceive: default...")
				b, err := Receive(buff)
				if err != nil {
					glog.Errorln("GrReceive error:", err)
				}
				glog.V(4).Infoln("GrReceive b: ", b)
				chRx <- b
				glog.V(4).Infoln("GrReceive: send chRx...")
				time.Sleep(5 * time.Millisecond)
			}
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func GrTransmit() {
	for {
		glog.V(4).Infoln("GrTransmit: ")
		b := <-chTx
		glog.V(4).Infoln("GrTransmit: got chTx...")
		err := Transmit(b)
		if err != nil {
			glog.Errorln("GrTransmit error: ", err)
		}
		time.Sleep(3 * time.Millisecond)
	}
}

func GrRxPrase() {
	var rxBuff []byte
	for {
		select {
		case rx := <-chRx: // 收到你的来信
			glog.V(4).Infoln("GrRxPrase: got chRx...")

			glog.V(4).Infoln("had buff: ", rxBuff)

			rxBuff = append(rxBuff, rx...) // 深藏我的心底

			glog.V(4).Infoln("got buff: ", rxBuff)

			// 解开长情的信笺
			// 残余的信亦不能忘却
			var vars []variable.CmdT
			vars, rxBuff = variable.Unpack(rxBuff)

			// 所有的酸甜苦辣都值得铭记
			glog.V(4).Infoln("left buff: ", rxBuff)
			glog.V(4).Infof("got vars: %v\n", vars)

			if len(vars) > 10 {
				glog.Fatalln("Too many vars")
			}

			// 拼凑出变量的清单
			chart, add, del := variable.Filt(vars)
			if len(chart) != 0 {
				b, _ := json.Marshal(chart)
				Chch <- string(b)
			}

			glog.V(3).Infoln("len(chart): ", len(chart))
			if glog.V(2) && len(add) > 0 || len(del) > 0 {
				glog.Infof("add: %v, del: %v\n", add, del)
			}

			// 挂念的变量，还望顺问近祺
			go func() {
				for _, v := range add {
					err := SendCmd(variable.Subscribe, v)
					if err != nil {
						glog.Errorln("SendCmd error:", err)
					}
				}
			}()

			// 无缘的变量，就请随风逝去
			go func() {
				for _, v := range del {
					err := SendCmd(variable.Unsubscribe, v)
					if err != nil {
						glog.Errorln("SendCmd error:", err)
					}
				}
			}()

		case <-time.After(200 * time.Millisecond): // 200ms不见
			if SerialCur.Port == nil || SerialCur.Name == "" {
				break
			}
			glog.V(4).Infoln("GrRxPrase: time after 200ms...")
			// 甚是想念
			_, add, _ := variable.Filt([]variable.CmdT{})
			glog.V(3).Infoln("add: ", add)
			for _, v := range add {
				err := SendCmd(variable.Subscribe, v)
				if err != nil {
					glog.Errorln("SendCmd error:", err)
				}
			}
		}
	}
}
