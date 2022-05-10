package serial

import (
	"errors"
	"math"
	"time"

	"go.bug.st/serial"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/datautil"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

type testPort struct {
	readingAddresses []uint32
	createdTime      time.Time
}

var BoardSysTime time.Time = time.Now() // 虚拟电路板的系统时间

func newTestPort() serial.Port {
	glog.Infoln("TestPort open at: ", BoardSysTime)
	return &testPort{
		readingAddresses: []uint32{},
		createdTime:      BoardSysTime,
	}
}

func (tp *testPort) SetMode(mode *serial.Mode) error { return nil }

func testValue(x float64, addr uint32) float64 {
	ratio := float64(addr&0xF) * 16.0
	phase := float64((addr >> 4) & 0xF)
	freq := (float64((addr>>8)&0xFF) - 0x80) / 16.0
	freq = math.Exp(freq)
	amplitude := float64((addr>>16)&0xFF) / 16.0
	amplitude = math.Exp(amplitude)
	waveform := (addr >> 24) & 0xF
	currentPhase := x*freq + phase
	dist := currentPhase - math.Floor(currentPhase)
	var scale float64
	switch waveform {
	case 0: // square
		if dist < 0.5 {
			scale = 1.0
		} else {
			scale = -1.0
		}
	case 1: // triangle
		scale = 4 * (math.Abs(dist-0.5) - 0.25)
	case 2: // scan
		scale = 2 * (dist - 0.5)
	case 3: //exsin
		scale = math.Exp(x/ratio) * math.Sin(currentPhase*2*math.Pi)
	case 4: //e-xsin
		scale = math.Exp(-x/ratio) * math.Sin(currentPhase*2*math.Pi)
	default: // sin
		scale = math.Sin(currentPhase * 2 * math.Pi)
	}
	y := scale * amplitude
	glog.V(3).Infoln("Test port: Address: 0x%08X %.5f => %.5f\n", addr, x, y)
	return y
}

func (tp *testPort) Read(p []byte) (n int, err error) {
	addresses := tp.readingAddresses
	maxNumPack := len(p) / 20
	if len(addresses) > maxNumPack {
		addresses = addresses[:maxNumPack]
	}

	for i, addr := range addresses {
		s := p[20*i : 20*(i+1)]
		s[0] = 1                                // 单片机代号 board
		s[1] = 2                                // 响应或错误代号 act (0x02 = 订阅的正常返回)
		s[2] = 8                                // 数据长度 typeLen
		copy(s[3:7], variable.AnyToBytes(addr)) // 单片机地址
		t := time.Since(tp.createdTime)
		x := t.Seconds()
		u := t.Milliseconds()
		y := testValue(x, addr)
		copy(s[7:15], variable.AnyToBytes(y))          // 数据
		copy(s[15:19], variable.AnyToBytes(uint32(u))) // 时间戳
		s[19] = '\n'                                   // 尾部固定为0x0a
	}
	return len(addresses) * 20, nil
}

func (tp *testPort) Write(p []byte) (n int, err error) {
	if len(p) != 16 {
		return 0, errors.New("invalid len")
	}
	if p[len(p)-1] != '\n' {
		return 0, errors.New("invalid package")
	}

	board := p[0]
	if board != 1 {
		return 0, errors.New("invalid board")
	}
	act := datautil.ActMode(p[1])
	typeLen := p[2]
	if typeLen != 8 {
		return 0, errors.New("unsupported typeLen")
	}
	address := variable.BytesToUint32(p[3:7])

	switch act {
	case datautil.Subscribe:
		tp.readingAddresses = append(tp.readingAddresses, address)
		glog.Infoln("Adding address: %08X\n", address)

	case datautil.Unsubscribe:
		var newAddresses []uint32
		for _, addr := range tp.readingAddresses {
			if addr != address {
				newAddresses = append(newAddresses, addr)
			}
		}
		tp.readingAddresses = newAddresses
		glog.Infoln("Deleting address: %08X\n", address)

	default:
		return 0, errors.New("invalid act")
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

func (tp *testPort) Close() error { return nil }
