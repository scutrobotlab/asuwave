package serial

import (
	"errors"
	"fmt"
	"math"
	"time"

	"go.bug.st/serial"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/variable"
	"github.com/scutrobotlab/asuwave/pkg/slip"
)

/**
0x20000000 静止，可写变量
0x30000000 动态变化，double
0x40000000 动态变化，float
0x50000000 动态变化，int8
0x60000000 动态变化，int16
0x70000000 动态变化，int32
0x80000000 动态变化，int64
**/
var vartypeMap = map[uint32]string{
	3: "double",
	4: "float",
	5: "int8_t",
	6: "int16_t",
	7: "int32_t",
	8: "int64_t",
}

var (
	chAddr                         = make(chan bool, 10) // 修改通知
	addresses    map[uint32]bool   = map[uint32]bool{}   // 观察的地址
	writeData    map[uint32][]byte = map[uint32][]byte{} // 直接写入数据
	BoardSysTime time.Time         = time.Now()          // 虚拟电路板的系统时间
)

type testPort struct{}

func newTestPort() serial.Port {
	glog.Infoln("TestPort open at: ", BoardSysTime)
	return &testPort{}
}

func (tp *testPort) SetMode(mode *serial.Mode) error { return nil }

func testValue(x float64, addr uint32) []byte {
	if data, ok := writeData[addr]; ok {
		return data
	}

	ratio := float64(addr&0xF) * 8.0
	phase := float64((addr >> 4) & 0xF)
	freq := (float64((addr>>8)&0xFF) - 0x80) / 64.0
	freq = math.Exp(freq)
	amplitude := float64((addr>>16)&0xFF) / 16.0
	amplitude = math.Exp(amplitude)
	waveform := (addr >> 24) & 0xF
	vartype := (addr >> 28) & 0xF
	currentPhase := x*freq + phase
	dist := currentPhase - math.Floor(currentPhase)

	glog.V(5).Info("ratio=", ratio)
	glog.V(5).Info("phase=", phase)
	glog.V(5).Info("freq=", freq)
	glog.V(5).Info("amplitude=", amplitude)
	glog.V(5).Info("currentPhase=", currentPhase)
	glog.V(5).Info("dist=", dist)
	glog.V(5).Infof("vartype=%d(%s)", vartype, vartypeMap[vartype])

	var scale float64
	switch waveform {
	case 0: // square
		glog.V(5).Info("waveform=square")
		if dist < 0.5 {
			scale = 1.0
		} else {
			scale = -1.0
		}
	case 1: // triangle
		glog.V(5).Info("waveform=triangle")
		scale = 4 * (math.Abs(dist-0.5) - 0.25)
	case 2: // scan
		glog.V(5).Info("waveform=scan")
		scale = 2 * (dist - 0.5)
	case 3: //exsin
		glog.V(5).Info("waveform=exsin")
		scale = math.Exp(x/ratio) * math.Sin(currentPhase*2*math.Pi)
	case 4: //e-xsin
		glog.V(5).Info("waveform=e-xsin")
		scale = math.Exp(-x/ratio) * math.Sin(currentPhase*2*math.Pi)
	default: // sin
		glog.V(5).Info("waveform=sin")
		scale = math.Sin(currentPhase * 2 * math.Pi)
	}
	y := scale * amplitude

	if t, ok := vartypeMap[vartype]; ok {
		data := variable.SpecToBytes(t, y)
		return data
	}

	return make([]byte, 8)
}

func (tp *testPort) Read(p []byte) (n int, err error) {
	for len(addresses) == 0 {
		if _, ok := <-chAddr; !ok {
			return 0, nil
		}
	}
	data := make([]byte, 0, len(addresses)*40)

	i := 0
	for addr := range addresses {
		var pdu [20]byte
		pdu[0] = 1                                // 单片机代号 board
		pdu[1] = 2                                // 响应或错误代号 act (0x02 = 订阅的正常返回)
		pdu[2] = 8                                // 数据长度 length
		copy(pdu[3:7], variable.AnyToBytes(addr)) // 单片机地址
		t := time.Since(BoardSysTime)
		x := t.Seconds()
		u := t.Milliseconds()
		y := testValue(x, addr)
		copy(pdu[7:15], y)                               // 数据
		copy(pdu[15:19], variable.AnyToBytes(uint32(u))) // 时间戳
		pdu[19] = '\n'                                   // 尾部固定为0x0a
		i++
		sdu := slip.Pack(pdu[:])
		data = append(data, sdu...)
	}

	return copy(p, data), nil
}

func (tp *testPort) Write(p []byte) (n int, err error) {
	if p[len(p)-1] != '\n' {
		return 0, errors.New("invalid package")
	}
	board := p[0]
	glog.Infoln("Got write: board = ", board)
	act := variable.ActMode(p[1])
	glog.Infoln("Got write: act = ", act)
	length := p[2]
	glog.Infoln("Got write: length = ", length)
	address := variable.BytesToUint32(p[3:7])
	glog.Infoln("Got write: address = ", address)
	data := p[7:15]

	switch act {
	case variable.Subscribe:
		go time.AfterFunc(500*time.Millisecond, func() {
			addresses[address] = true
			chAddr <- true
			glog.Infof("Adding address: %08X\n", address)
		})

	case variable.Unsubscribe:
		go time.AfterFunc(500*time.Millisecond, func() {
			delete(addresses, address)
			chAddr <- true
			glog.Infof("Deleting address: %08X\n", address)
		})

	case variable.Write:
		go time.AfterFunc(500*time.Millisecond, func() {
			writeData[address] = data
			glog.Infof("Writing address: %08X = %v\n", address, data)
		})

	default:
		return 0, errors.New(fmt.Sprint("invalid act: ", act))
	}

	return 16, nil
}

func (tp *testPort) ResetInputBuffer() error { return nil }

func (tp *testPort) ResetOutputBuffer() error { return nil }

func (tp *testPort) SetDTR(dtr bool) error { return errors.New("not supported") }

func (tp *testPort) SetRTS(rts bool) error { return errors.New("not supported") }

func (tp *testPort) SetReadTimeout(timeout time.Duration) error { return nil }

func (tp *testPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return nil, errors.New("not supported")
}

func (tp *testPort) Close() error {
	close(chAddr)
	return nil
}
