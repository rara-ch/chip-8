package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rara-ch/chip-8/internal/display"
	"github.com/rara-ch/chip-8/internal/memory"
)

type stack [12]uint16
type registers [16]byte

// const pixel = 10

type chip8 struct {
	memory    memory.Memory
	display   display.Display
	stack     stack
	registers registers
	pc        int
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("please include a path to the .ch8 file you would like to run")
	}

	file, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatalf("could not read file: %v", err)
	}

	chip8 := &chip8{
		memory:    memory.NewMemory(file),
		display:   display.Display{},
		stack:     stack{},
		registers: registers{},
		pc:        0x200,
	}
	fmt.Println(chip8)
}
