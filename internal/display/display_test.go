package display

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDisplay(t *testing.T) {
	// Valid Test: Clear display
	display := Display{}
	display[0x00] = true
	display[0x01] = true
	display.Clear()
	assert.False(t, display[0x00])
	assert.False(t, display[0x01])

	// Valid Test: Draw a 5px line in the top right corner
	display = Display{}
	err := display.Draw(0x00, 0x00, []bool{true, true, true, true, true})
	require.NoError(t, err)
	assert.Equal(t, display, Display{true, true, true, true, true})
	assert.False(t, display[0x06])

	// Valid Test: Draw a 2px vertical line in the top right corner
	display = Display{}
	err = display.Draw(0x00, 0x00, []bool{true, false, false, false, false, true, false, false, false, false})
	require.NoError(t, err)
	assert.True(t, display[0x00])
	assert.True(t, display[0x05])
	assert.False(t, display[0x01])
	assert.False(t, display[0x04])

	// Invalid Test: Sprite does not have 5px width
	display = Display{}
	err = display.Draw(0x00, 0x00, []bool{true})
	require.Error(t, err)

	// Invalid Test: xPosition is too high
	display = Display{}
	err = display.Draw(0x50, 0x00, []bool{true, true, true, true, true})
	require.Error(t, err)

	// Invalid Test: yPosition is too high
	display = Display{}
	err = display.Draw(0x00, 0x50, []bool{true, true, true, true, true})
	require.Error(t, err)
}
