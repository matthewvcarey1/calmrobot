package mapland

type MapLand struct {
	size int
	land [][]bool
}

func New(size uint) *MapLand {
	land := make([][]bool, size)
	for i := range land {
		land[i] = make([]bool, size)
	}
	return &MapLand{size: len(land[0]), land: land}
}

func (m *MapLand) GetMinMax() (int, int, int, int) {
	return (m.size / 2) - m.size, (m.size / 2) - m.size, m.size / 2, m.size / 2
}

func (m *MapLand) Set(x int, y int) {
	x += m.size / 2
	y += m.size / 2
	m.land[x][y] = true
}

func (m *MapLand) Count() int {
	count := 0
	for _, x := range m.land {
		for _, y := range x {
			if y {
				count++
			}
		}
	}
	return count
}
