package robotview

import (
	"errors"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
)

// RobotView type
type RobotView struct {
	land *mapland.MapLand
}

const (
	// The safe threshold
	safe = 23
)

// Get array of digits in a positive or negative decimal number
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

// sum an array of ints
func sum(arr []int) int {
	res := int(0)
	for _, v := range arr {
		res += v
	}
	return res
}

// isSafe checks whether a coordinate is safe
func isSafe(x int, y int) bool {
	digits := getDigits(x)
	digits = append(digits, getDigits(y)...)
	return (sum(digits) <= safe)
}

// New RobotView
func New(land *mapland.MapLand) *RobotView {
	return &RobotView{land: land}
}

// robot storing its coodinates
type robot struct {
	x int
	y int
}

// Mark all the mines in the land
func (r *RobotView) MarkMines() {
	x1, y1, x2, y2 := r.land.GetStartEnd()
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

// east makes a new eastwards robot
func (r robot) east() robot {
	return robot{r.x + 1, r.y}
}

// west makes a new westwards robot
func (r robot) west() robot {
	return robot{r.x - 1, r.y}
}

// south makes a new southwards robot
func (r robot) south() robot {
	return robot{r.x, r.y - 1}
}

// north makes a new northwards robot
func (r robot) north() robot {
	return robot{r.x, r.y + 1}
}

// FloodFill marks the available non mined space as accesable
// by the robot and returns the number of coordinates it contains
func (rv *RobotView) FloodFill(x int, y int) (int, error) {
	var queue []robot
	x1, y1, x2, y2 := rv.land.GetStartEnd()
	count := int(0)
	if !rv.land.IsClear(x, y) {
		return 0, nil
	}
	rv.land.SetAccessable(x, y)
	count++
	queue = append(queue, robot{
		x: x,
		y: y,
	})
	for len(queue) > 0 {
		r := queue[0]     // first element
		queue = queue[1:] //Dequeue
		if (r.x) == x1 || r.x == x2-1 || r.y == y1 || r.y == y2-1 {
			return 0, errors.New("FloodFill has reached the edge of the world, the world needs to be bigger")
		}
		if w := r.west(); rv.land.IsClear(w.x, w.y) {
			rv.land.SetAccessable(w.x, w.y)
			count++
			queue = append(queue, w)
		}
		if e := r.east(); rv.land.IsClear(e.x, e.y) {
			rv.land.SetAccessable(e.x, e.y)
			count++
			queue = append(queue, e)
		}
		if n := r.north(); rv.land.IsClear(n.x, n.y) {
			rv.land.SetAccessable(n.x, n.y)
			count++
			queue = append(queue, n)
		}
		if s := r.south(); rv.land.IsClear(s.x, s.y) {
			rv.land.SetAccessable(s.x, s.y)
			count++
			queue = append(queue, s)
		}
	}
	return count, nil
}
