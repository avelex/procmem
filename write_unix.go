//go:build linux || darwin
// +build linux darwin

package procmem

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/sys/unix"
)

func Write(pid int, address uintptr, data []byte) error {
	len := len(data)

	localIovec := unix.Iovec{
		Base: &data[0],
		Len:  uint64(len),
	}
	remoteIovec := unix.RemoteIovec{
		Base: address,
		Len:  len,
	}

	n, err := unix.ProcessVMWritev(pid, []unix.Iovec{localIovec}, []unix.RemoteIovec{remoteIovec}, 0)
	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("pid[%v] address[0x%x] 0 bytes wrote", pid, address)
	}

	return nil
}

func WriteUint32(pid int, address uintptr, value uint32) error {
	buff := make([]byte, UINT32_LEN)
	binary.LittleEndian.PutUint32(buff, value)

	return Write(pid, address, buff)
}

func WriteUint64(pid int, address uintptr, value uint64) error {
	buff := make([]byte, UINT64_LEN)
	binary.LittleEndian.PutUint64(buff, value)

	return Write(pid, address, buff)
}

// if your dist is golang program, it will be work if string it's array of char
func WriteString(pid int, address uintptr, value string) error {
	return Write(pid, uintptr(address), []byte(value))
}
