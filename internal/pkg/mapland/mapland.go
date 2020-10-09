package mapland

import "fmt"

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
		for iq := range land[i] {
			land[i][iq] = 0x20
		}
	}
	return &MapLand{size: len(land[0]), half: len(land[0])>>1, land: land}
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

// Set mark a coordinate as visitable
func (m *MapLand) Set(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = 0x4F
}

// SetMine mark a coordinate as a mine
func (m *MapLand) SetMine(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = 0x58
}

// SetClear mark a coordinate as not a mine
func (m *MapLand) SetClear(x int, y int) {
	x += m.half
	y += m.half
	m.land[x][y] = 0x20
}

// Count the robot visitable area
func (m *MapLand) Count() int {
	count := 0
	for _, x := range m.land {
		for _, y := range x {
			if y == 0x4F {
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
