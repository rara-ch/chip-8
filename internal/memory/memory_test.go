package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemory(t *testing.T) {
	// Valid Test: Load program into memory
	program := []byte{0x11, 0xE3}
	memory := NewMemory(program)
	assert.Equal(t, memory.data[0x200], byte(0x11))
	assert.Equal(t, memory.data[0x201], byte(0xE3))
	assert.Equal(t, memory.data[0x202], byte(0x00))

	// Valid Test: Write value
	memory = NewMemory([]byte{})
	err := memory.Write(0x0100, 0x11)
	require.NoError(t, err)
	assert.Equal(t, memory.data[0x100], byte(0x11))

	// Valid Test: Read value
	memory = NewMemory([]byte{})
	err = memory.Write(0x0100, 0x11)
	require.NoError(t, err)
	value, err := memory.Read(0x0100)
	assert.Equal(t, value, byte(0x11))

	// Invalid Test: Write value to an address thats too high
	memory = NewMemory([]byte{})
	err = memory.Write(0x1000, 0x11)
	require.Error(t, err)

	// Invalid Test: Read from an address thats too high
	memory = NewMemory([]byte{})
	_, err = memory.Read(0x1000)
	require.Error(t, err)
}
