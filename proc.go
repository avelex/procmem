package procmem

type Process interface {
	Read(dst []byte, address uintptr, size uint64) (n int, err error)
	ReadUint32(address uintptr) (uint32, error)
	ReadUint64(address uintptr) (uint64, error)
	ReadString(address uintptr, size uint64) (string, error)
	ReadGoString(address uintptr, size uint64) (string, error)
	Write(address uintptr, data []byte) error
	WriteUint32(address uintptr, value uint32) error
	WriteUint64(address uintptr, value uint64) error
	WriteString(address uintptr, value string) error
	// Do not close the certain process
	Close() error
}
