package robotview

import (
	"github.com/matthewvcarey1/calmrobot/pkg/internal/pkg/mapland"
)

type RobotView struct {
	land *mapland.MapLand
}

const (
	safe = 23
)

func getDigits(num int) []int {
	var res []int
	if num < 0 {
		num = -num
	}
	if num == 0 {
		res = append(res, 0)
	} else {
		for ; num > 0; num = num / 10 {
			digit := num % 10
			res = append(res, digit)
		}
	}
	return res
}

func sum(arr []int) int {
	res := int(0)
	for _, v := range arr {
		res += v
	}
	return res
}

func isSafe(x int, y int) bool {
	digits := getDigits(x)
	digits = append(digits, getDigits(y)...)
	return (sum(digits) <= safe)
}

func New(land *mapland.MapLand) *RobotViewView {
	return &RobotView{land: land}
}

type robot struct {
	x int
	y int
}

func (r *RobotView) MarkMines() {
	x1, y1, x2, y2 := r.land.GetMinMax()
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			if !isSafe(x, y) {
				r.land.SetMine(x, y)
			} else {
				r.land.SetClear(x, y)
			}
		}
	}
}

func (r robot) east() robot {
	return robot{r.x + 1, r.y}
}

func (r robot) west() robot {
	return robot{r.x - 1, r.y}
}

func (r robot) south() robot {
	return robot{r.x, r.y - 1}
}

func (r robot) north() robot {
	return robot{r.x, r.y + 1}
}

func (rv *RobotView) FloodFill(x int, y int) int {
	var queue []robot
	count := int(0)
	if rv.land.Get(x, y) != 0x20 {
		return 0
	}
	rv.land.Set(x, y)
	count++
	queue = append(queue, robot{
		x: x,
		y: y,
	})
	for len(queue) > 0 {
		r := queue[0]     // first element
		queue = queue[1:] //Dequeue
		if w := r.west(); rv.land.Get(w.x, w.y) == 0x20 {
			rv.land.Set(w.x, w.y)
			count++
			queue = append(queue, w)
		}
		if e := r.east(); rv.land.Get(e.x, e.y) == 0x20 {
			rv.land.Set(e.x, e.y)
			count++
			queue = append(queue, e)
		}
		if n := r.north(); rv.land.Get(n.x, n.y) == 0x20 {
			rv.land.Set(n.x, n.y)
			count++
			queue = append(queue, n)
		}
		if s := r.south(); rv.land.Get(s.x, s.y) == 0x20 {
			rv.land.Set(s.x, s.y)
			count++
			queue = append(queue, s)
		}
	}
	return count
}
