package display

import "fmt"

const Height = 64
const Width = 32

type Display [Height * Width]bool

// Opcode 00E0
func (d *Display) Clear() {
	for i := range d {
		d[i] = false
	}
}

// Opcode
func (d *Display) Draw(positionX, positionY byte, sprite []bool) error {
	spriteWidth := len(sprite) % 5
	if spriteWidth != 0 {
		return fmt.Errorf("all sprites must have a width of five pixels: %v", spriteWidth)
	}

	if positionX > 0x20 {
		return fmt.Errorf("positionX must be less than the width of the display (0x20): %v", positionX)
	}

	if positionY > 0x40 {
		return fmt.Errorf("positionX must be less than the height of the display (0x40): %v", positionX)
	}

	start := (positionY * byte(Height)) + positionX

	for i, value := range sprite {
		d[start+byte(i)] = value
	}

	return nil
}
