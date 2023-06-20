package procmem

import (
	"encoding/binary"
	"math"
)

func Float32ToBytes(v float32) []byte {
	buff := make([]byte, FLOAT32_LEN)
	bits := math.Float32bits(v)
	binary.LittleEndian.PutUint32(buff, bits)
	return buff
}

func BytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToBytes(v float64) []byte {
	buff := make([]byte, FLOAT64_LEN)
	bits := math.Float64bits(v)
	binary.LittleEndian.PutUint64(buff, bits)
	return buff
}

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
