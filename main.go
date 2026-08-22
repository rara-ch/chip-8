package main

import (
	"fmt"

	"github.com/rara-ch/chip-8/internal/memory"
)

func main() {
	memory := memory.NewMemory()
	fmt.Println("Memory: ", memory)
}
