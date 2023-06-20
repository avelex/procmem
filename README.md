# Process Memory Read-Writer on Unix and Windows

## Usage

```go
package main

import (
	"github.com/avelex/procmem"
)

func main() {
    pid := 164433
    address := 0xc00012c008
    proc := procmem.NewProcess(pid) 

    want := "Hello"

    got, err := proc.ReadString(uintptr(address), len(want))
    if err != nil {
        ...
    }

    fmt.Println(want == got) // true

    newString := "World"

    if err := proc.WriteString(uintptr(ptr), newString); err != nil {
        ...
    }

    got, err = proc.ReadString(uintptr(address), len(newString))
    if err != nil {
        ...
    }

    fmt.Println(newString == got) // true
}
```
