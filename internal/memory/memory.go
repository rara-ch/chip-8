package memory

import "fmt"

const size = 4096

type Memory struct {
	data *[size]byte
}

func NewMemory(program []byte) Memory {
	memory := Memory{
		data: &[size]byte{},
	}
	memory.loadFont()
	memory.loadProgram(program)
	return memory
}

func (m *Memory) Read(address uint16) (byte, error) {
	if address >= 0x1000 {
		return 0x00, fmt.Errorf("address must be a 12 bits long: 0x%04X", address)
	}
	return m.data[address], nil
}

func (m *Memory) Write(address uint16, value byte) error {
	if address >= 0x1000 {
		return fmt.Errorf("address must be a 12 bits long: 0x%04X", address)
	}
	m.data[address] = value
	return nil
}

func (m *Memory) loadFont() {
	address := 0x050
	for _, char := range fonts {
		for _, line := range char {
			m.data[address] = line
			address++
		}
	}
}

func (m *Memory) loadProgram(program []byte) {
	copy(m.data[0x200:], program)
}
