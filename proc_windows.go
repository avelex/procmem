//go:build windows
// +build windows

package procmem

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/windows"
)

/*
Windows Process Handler

Don't forget to close process!
*/
type process struct {
	handle windows.Handle
}

func NewProcess(pid int) (Process, error) {
	h, err := windows.OpenProcess(
		windows.PROCESS_VM_READ|windows.PROCESS_VM_WRITE|windows.PROCESS_VM_OPERATION, false, uint32(pid))
	if err != nil {
		return nil, err
	}

	return &process{
		handle: h,
	}, nil
}

func (p *process) Read(dst []byte, address uintptr, size uint64) (n int, err error) {
	return read(p.handle, dst, address, size)
}

func (p *process) ReadUint32(address uintptr) (uint32, error) {
	buff := make([]byte, UINT32_LEN)

	n, err := p.Read(buff, address, UINT32_LEN)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(buff[:n]), nil
}

func (p *process) ReadUint64(address uintptr) (uint64, error) {
	buff := make([]byte, UINT64_LEN)

	n, err := p.Read(buff, address, UINT64_LEN)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(buff[:n]), nil
}

func (p *process) ReadString(address uintptr, size uint64) (string, error) {
	buff := make([]byte, size)

	n, err := p.Read(buff, address, size)
	if err != nil {
		return "", err
	}

	return unsafe.String(unsafe.SliceData(buff), n), nil
}

func (p *process) ReadGoString(address uintptr, size uint64) (string, error) {
	ptr, err := p.ReadUint64(address)
	if err != nil {
		return "", err
	}

	return p.ReadString(uintptr(ptr), size)
}

func (p *process) Write(address uintptr, data []byte) error {
	return write(p.handle, address, data)
}

func (p *process) WriteUint32(address uintptr, value uint32) error {
	buff := make([]byte, UINT32_LEN)
	binary.LittleEndian.PutUint32(buff, value)

	return p.Write(address, buff)
}

func (p *process) WriteUint64(address uintptr, value uint64) error {
	buff := make([]byte, UINT64_LEN)
	binary.LittleEndian.PutUint64(buff, value)

	return p.Write(address, buff)
}

func (p *process) WriteString(address uintptr, value string) error {
	return p.Write(address, []byte(value))
}

func (p *process) WriteFloat32(addres uintptr, value float32) error {
	return p.Write(address, Float32ToBytes(value))
}

func (p *process) WriteFloat64(addres uintptr, value float64) error {
	return p.Write(address, Float64ToBytes(value))
}

func (p *process) Close() error {
	return windows.CloseHandle(p.handle)
}
