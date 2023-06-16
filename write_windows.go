//go:build windows
// +build windows

package procmem

import (
	"encoding/binary"

	"golang.org/x/sys/windows"
)

func Write(pid int, address uintptr, data []byte) error {
	h, err := windows.OpenProcess(windows.PROCESS_VM_WRITE|windows.PROCESS_VM_OPERATION, false, uint32(pid))
	if err != nil {
		return err
	}

	defer windows.CloseHandle(h)

	return write(h, address, data)
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

func WriteString(pid int, address uintptr, value string) error {
	return Write(pid, address, []byte(value))
}

func write(handle windows.Handle, address uintptr, data []byte) error {
	return windows.WriteProcessMemory(handle, address, &data[0], uintptr(len(data)), nil)
}
