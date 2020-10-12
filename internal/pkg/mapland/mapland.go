package mapland

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

const (
	clear      = 0x20
	mine       = 0x58
	accessable = 0x4F
)

// MapLand type
type MapLand struct {
	size int
	half int
	land [][]byte
}

// New MapLand
func New(size uint) *MapLand {
	if size%2 > 0 {
		size++
	}

	land := make([][]byte, size)
	for i := range land {
		land[i] = make([]byte, size)
	}
	return &MapLand{size: len(land[0]), half: len(land[0]) >> 1, land: land}
}

// GetStartEnd the start and ends of the land in robotview coordinates
func (m *MapLand) GetStartEnd() (int, int, int, int) {
	return (m.half) - m.size, m.half - m.size, m.half, m.half
}

// Get coordinate contents
func (m *MapLand) Get(x int, y int) byte {
	x += m.half
	y += m.half
	return m.land[x][y]
}

func (m *MapLand) IsClear(x int, y int) bool {
	return m.Get(x, y) == clear
}

// Set mark a coordinate as visitable
func (m *MapLand) SetAccessable(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = accessable
}

// SetMine mark a coordinate as a mine
func (m *MapLand) SetMine(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = mine
}

// SetClear mark a coordinate as not a mine
func (m *MapLand) SetClear(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = clear
}

// This is not used but could be used for testing.
// Count the robot visitable area
func (m *MapLand) Count() int {
	count := 0
	for _, x := range m.land {
		for _, y := range x {
			if y == accessable {
				count++
			}
		}
	}
	return count
}

// Draw the whole map
func (m *MapLand) Draw() {
	for _, x := range m.land {
		fmt.Println(string(x))
	}
}

// Draw the whole map as a png file
func (m *MapLand) DrawImage(fname string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{m.size, m.size}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	Red := color.RGBA{0xff, 0x00, 0x00, 0xff}
	accessableColour := color.RGBA{0x1d, 0xd8, 0x80, 0xff}
	for x, c := range m.land {
		for y, v := range c {
			px := m.size - x
			py := m.size - y
			switch v {
			case clear:
				img.Set(px, py, color.White)
			case mine:
				img.Set(px, py, Red)
			case accessable:
				img.Set(px, py, accessableColour)
			default:
				// Use zero value.
			}
		}
	}
	f, err := os.Create(fname)
	if err != nil {
		println("Could not create output file:", fname)
		return
	}
	err = png.Encode(f, img)
	if err != nil {
		println("Could not write output file:", fname)
	}
}
