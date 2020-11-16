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

// CalcMaplandSize calculates length of sides needed for a given safe value
func CalcMaplandSize(trigger int) int {
	trigger++
	nines := trigger / 9
	extra := trigger % 9
	result := extra
	for n := 0; n < nines; n++ {
		result = result * 10
		result += 9
	}
	return (result + 1)
}

// New MapLand
func New(trigger int) *MapLand {
	size := CalcMaplandSize(trigger)
	fmt.Println("Size", size)
	land := make([][]byte, size)
	for i := range land {
		land[i] = make([]byte, size)
	}
	return &MapLand{size: len(land[0]), half: len(land[0]) >> 1, land: land}
}

// GetStartEnd the start and ends of the land in robotview coordinates
func (m *MapLand) GetStartEnd() (int, int, int, int) {
	return 0, 0, m.size, m.size
}

// Get coordinate contents
func (m *MapLand) Get(x int, y int) byte {
	return m.land[x][y]
}

// IsClear returns true if coordinate is clear
func (m *MapLand) IsClear(x int, y int) bool {
	if y < 0 || x < 0 {
		return false
	}
	return m.Get(x, y) == clear
}

// SetAccessable mark a coordinate as visitable
func (m *MapLand) SetAccessable(x int, y int) {
	m.land[x][y] = accessable
}

// SetMine mark a coordinate as a mine
func (m *MapLand) SetMine(x int, y int) {
	m.land[x][y] = mine
}

// SetClear mark a coordinate as not a mine
func (m *MapLand) SetClear(x int, y int) {
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

// DrawImage the whole map as a png file using the one quadrant that we have calculated
func (m *MapLand) DrawImage(fname string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{m.size * 2, m.size * 2}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	mineColour := color.RGBA{0xff, 0x00, 0x00, 0xff}
	clearColour := color.White
	accessableColour := color.RGBA{0x1d, 0xd8, 0x80, 0xff}
	for x, c := range m.land {
		for y, v := range c {
			px := m.size - x
			py := m.size - y
			qx := m.size + x
			qy := m.size - y
			rx := m.size - x
			ry := m.size + y
			sx := m.size + x
			sy := m.size + y
			switch v {
			case clear:
				img.Set(px, py, clearColour)
				img.Set(qx, qy, clearColour)
				img.Set(rx, ry, clearColour)
				img.Set(sx, sy, clearColour)
			case mine:
				img.Set(px, py, mineColour)
				img.Set(qx, qy, mineColour)
				img.Set(rx, ry, mineColour)
				img.Set(sx, sy, mineColour)
			case accessable:
				img.Set(px, py, accessableColour)
				img.Set(qx, qy, accessableColour)
				img.Set(rx, ry, accessableColour)
				img.Set(sx, sy, accessableColour)
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
