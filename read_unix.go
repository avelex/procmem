//go:build linux || darwin
// +build linux darwin

package procmem

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/unix"
)

func Read(pid int, dst []byte, address uintptr, size uint64) (n int, err error) {
	localIovec := unix.Iovec{
		Base: &dst[0],
		Len:  size,
	}
	remoteIovec := unix.RemoteIovec{
		Base: address,
		Len:  int(size),
	}

	n, err = unix.ProcessVMReadv(pid, []unix.Iovec{localIovec}, []unix.RemoteIovec{remoteIovec}, 0)

	return
}

func ReadUint32(pid int, address uintptr) (uint32, error) {
	buff := make([]byte, UINT32_LEN)

	n, err := Read(pid, buff, address, UINT32_LEN)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buff[:n]), nil
}

func ReadUint64(pid int, address uintptr) (uint64, error) {
	buff := make([]byte, UINT64_LEN)

	n, err := Read(pid, buff, address, UINT64_LEN)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buff[:n]), nil
}

func ReadString(pid int, address uintptr, size uint64) (string, error) {
	buff := make([]byte, size)

	n, err := Read(pid, buff, uintptr(address), size)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(buff), n), nil
}

func ReadGoString(pid int, address uintptr, size uint64) (string, error) {
	ptr, err := ReadUint64(pid, address)
	if err != nil {
		return "", err
	}

	return ReadString(pid, uintptr(ptr), size)
}

func ReadFloat32(pid int, addres uintptr) (float32, error) {
	buff := make([]byte, FLOAT32_LEN)

	n, err := Read(pid, buff, addres, FLOAT32_LEN)
	if err != nil {
		return 0, err
	}

	return BytesToFloat32(buff[:n]), nil
}

func ReadFloat64(pid int, addres uintptr) (float64, error) {
	buff := make([]byte, FLOAT64_LEN)

	n, err := Read(pid, buff, addres, FLOAT64_LEN)
	if err != nil {
		return 0, err
	}

	return BytesToFloat64(buff[:n]), nil
}
