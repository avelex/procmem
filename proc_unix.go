//go:build linux || darwin
// +build linux darwin

package procmem

/*
UNIX Process Handler
*/
type process struct {
	pid int
}

func NewProcess(pid int) Process {
	return &process{
		pid: pid,
	}
}

func (p *process) Read(dst []byte, address uintptr, size uint64) (n int, err error) {
	return Read(p.pid, dst, address, size)
}

func (p *process) ReadUint32(address uintptr) (uint32, error) {
	return ReadUint32(p.pid, address)
}

func (p *process) ReadUint64(address uintptr) (uint64, error) {
	return ReadUint64(p.pid, address)
}

func (p *process) ReadString(address uintptr, size uint64) (string, error) {
	return ReadString(p.pid, address, size)
}

func (p *process) ReadGoString(address uintptr, size uint64) (string, error) {
	return ReadGoString(p.pid, address, size)
}

func (p *process) Write(address uintptr, data []byte) error {
	return Write(p.pid, address, data)
}

func (p *process) WriteUint32(address uintptr, value uint32) error {
	return WriteUint32(p.pid, address, value)
}

func (p *process) WriteUint64(address uintptr, value uint64) error {
	return WriteUint64(p.pid, address, value)
}

func (p *process) WriteString(address uintptr, value string) error {
	return WriteString(p.pid, address, value)
}

func (p *process) Close() error { return nil }
