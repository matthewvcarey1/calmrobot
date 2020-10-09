package mapland

import "fmt"

type MapLand struct {
	size int
	land [][]byte
}

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
	return &MapLand{size: len(land[0]), land: land}
}

func (m *MapLand) GetMinMax() (int, int, int, int) {
	half := m.size >> 1
	return (half) - m.size, half - m.size, half, half
}

func (m *MapLand) Get(x int, y int) byte {
	half := m.size >> 1
	x += half
	y += half
	return m.land[x][y]
}

func (m *MapLand) Set(x int, y int) {
	half := m.size >> 1
	x += half
	y += half
	m.land[x][y] = 0x4F
}

func (m *MapLand) SetMine(x int, y int) {
	half := m.size >> 1
	x += half
	y += half
	m.land[x][y] = 0x58
}

func (m *MapLand) SetClear(x int, y int) {
	half := m.size >> 1
	x += half
	y += half
	m.land[x][y] = 0x20
}

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

func (m *MapLand) Draw() {
	for _, x := range m.land {
		fmt.Println(string(x))
	}
}
