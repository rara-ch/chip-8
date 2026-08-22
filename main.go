package main

import (
	"fmt"

	"github.com/rara-ch/chip-8/internal/memory"
)

type stack [12]uint16
type registers [16]byte

func main() {
	memory := memory.NewMemory()
	stack := stack{}
	registers := registers{}
	pc := 0x200

	fmt.Println("Memory: ", memory)
	fmt.Println("Stack: ", stack)
	fmt.Println("Registers: ", registers)
	fmt.Println("Program Counter: ", pc)
}
