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

	float32Var := float32(3.14)
	float32VarPtr := &float32Var

	float64Var := float64(0.0000000001)
	float64VarPtr := &float64Var

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
		fmt.Printf("Float32Var: %v\n", float32Var)
		fmt.Printf("Float32VarPtr: %v\n", float32VarPtr)
		fmt.Println()
		fmt.Printf("Float64Var: %v\n", float64Var)
		fmt.Printf("Float64VarPtr: %v\n", float64VarPtr)
		fmt.Println()
	}
}
