# Process Memory Read-Writer on Unix and Windows


```go
    pid := 123
    proc := procmem.NewProcess(pid)

    ptr := 0xc000014350
    want := "avelex"

    got, err := proc.ReadGoString(ptr, len(want))

    fmt.Println(want == got) // true

```