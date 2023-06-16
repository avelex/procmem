//go:build windows
// +build windows

package procmem

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/windows"
)

func Read(pid int, dst []byte, address uintptr, size uint64) (n int, err error) {
	h, err := windows.OpenProcess(windows.PROCESS_VM_READ, false, uint32(pid))
	if err != nil {
		return 0, err
	}

	return read(h, dst, address, size)
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

	n, err := Read(pid, buff, address, size)
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

func read(handle windows.Handle, dst []byte, address uintptr, size uint64) (n int, err error) {
	ptr := uintptr(n)

	if err := windows.ReadProcessMemory(handle, address, &dst[0], uintptr(size), &ptr); err != nil {
		return 0, err
	}

	n = int(ptr)

	return
}
