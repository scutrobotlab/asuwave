package variable

import (
	"bytes"
	"encoding/binary"
	"math"
)

func AnyToBytes(i interface{}) []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, i)
	return buf.Bytes()
}

func SpecFromBytes(vType string, data []byte) float64 {
	switch vType {
	case "uint8_t":
		return float64(BytesToUint8(data))
	case "uint16_t":
		return float64(BytesToUint16(data))
	case "uint32_t":
		return float64(BytesToUint32(data))
	case "uint64_t":
		return float64(BytesToUint64(data))
	case "int8_t":
		return float64(BytesToInt8(data))
	case "int16_t":
		return float64(BytesToInt16(data))
	case "int32_t", "int":
		return float64(BytesToInt32(data))
	case "int64_t":
		return float64(BytesToInt64(data))
	case "float":
		return float64(BytesToFloat32(data))
	case "double":
		return float64(BytesToFloat64(data))
	default:
		return 0
	}
}
func SpecToBytes(vType string, i float64) []byte {
	switch vType {
	case "uint8_t":
		return AnyToBytes(uint8(i))
	case "uint16_t":
		return AnyToBytes(uint16(i))
	case "uint32_t":
		return AnyToBytes(uint32(i))
	case "uint64_t":
		return AnyToBytes(uint64(i))
	case "int8_t":
		return AnyToBytes(int8(i))
	case "int16_t":
		return AnyToBytes(int16(i))
	case "int32_t", "int":
		return AnyToBytes(int32(i))
	case "int64_t":
		return AnyToBytes(int64(i))
	case "float":
		return AnyToBytes(float32(i))
	case "double":
		return AnyToBytes(float64(i))
	default:
		return AnyToBytes(i)
	}
}

func BytesToInt8(i []byte) int8 {
	var o int8
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt16(i []byte) int16 {
	var o int16
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt32(i []byte) int32 {
	var o int32
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToInt64(i []byte) int64 {
	var o int64
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint8(i []byte) uint8 {
	var o uint8
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint16(i []byte) uint16 {
	var o uint16
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint32(i []byte) uint32 {
	var o uint32
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToUint64(i []byte) uint64 {
	var o uint64
	buf := bytes.NewReader(i)
	binary.Read(buf, binary.LittleEndian, &o)
	return o
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
