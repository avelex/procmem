package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	intVar := 123
	ptr := &intVar

	strArray := []byte{'H', 'e', 'l', 'l', 'o'}
	strArrayPtr := &strArray[0]

	strVar := "avelex"
	strPtr := &strVar

	buff := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Press ENTER to print vars")
		data := make([]byte, 32)
		buff.Read(data)

		pid := os.Getpid()
		fmt.Println("Process ID:", pid)
		fmt.Println("IntVar: ", intVar)
		fmt.Println("IntVar ptr:", ptr)
		fmt.Println()
		fmt.Println("StrArray: ", string(strArray))
		fmt.Println("StrArray ptr:", strArrayPtr)
		fmt.Println()
		fmt.Println("StrVar: ", string(strVar))
		fmt.Println("StrVar ptr:", strPtr)
		fmt.Println()
	}
}
