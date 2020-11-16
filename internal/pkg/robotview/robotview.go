package robotview

import (
	"container/list"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
)

// RobotView type
type RobotView struct {
	safe int
	land *mapland.MapLand
}

// SumDigits sum the absolute decimal digits in a number
func sumDigits(num int) int {
	res := 0
	if num < 0 {
		num = -num
	}
	if num == 0 {
		return res
	} else {
		for ; num > 0; num = num / 10 {
			digit := num % 10
			res += digit
		}
	}
	return res
}

// isSafe checks whether a coordinate is safe
func (rv *RobotView) isSafe(x int, y int) bool {
	val := sumDigits(x) + sumDigits(y)
	return val <= rv.safe
}

// New RobotView
func New(trigger int, land *mapland.MapLand) *RobotView {
	return &RobotView{safe: trigger, land: land}
}

// robot storing its coodinates
type robot struct {
	x int
	y int
}

// MarkMines mark all the mines in the land
func (rv *RobotView) MarkMines() {
	x1, y1, x2, y2 := rv.land.GetStartEnd()
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			if !rv.isSafe(x, y) {
				rv.land.SetMine(x, y)
			} else {
				rv.land.SetClear(x, y)
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

// returns the coords of this robot
// used in function parameters
func (r robot) coords() (int, int) {
	return r.x, r.y
}

type robotQueue struct {
	queue *list.List
}

func (q *robotQueue) size() int {
	return q.queue.Len()
}

func (q *robotQueue) push(r robot) {
	q.queue.PushBack(r)
}

func (q *robotQueue) pop() robot {
	e := q.queue.Front() // first element
	r := e.Value.(robot)
	q.queue.Remove(e)
	return r
}

func newRobotQueue() *robotQueue {
	q := list.New()
	return &robotQueue{queue: q}
}

// FloodFill marks the available non mined space as accessable
// by the robot and returns the number of coordinates it contains
func (rv *RobotView) FloodFill(x int, y int) (int, error) {
	queue := newRobotQueue()
	count := int(0)
	if !rv.land.IsClear(x, y) {
		return 0, nil
	}
	rv.land.SetAccessable(x, y)
	count++
	queue.push(robot{
		x: x,
		y: y,
	})
	for queue.size() > 0 {
		r := queue.pop()

		if w := r.west(); rv.land.IsClear(w.coords()) {
			rv.land.SetAccessable(w.coords())
			count++
			queue.push(w)
		}
		if e := r.east(); rv.land.IsClear(e.coords()) {
			rv.land.SetAccessable(e.coords())
			count++
			queue.push(e)
		}
		if n := r.north(); rv.land.IsClear(n.coords()) {
			rv.land.SetAccessable(n.coords())
			count++
			queue.push(n)
		}
		if s := r.south(); rv.land.IsClear(s.coords()) {
			rv.land.SetAccessable(s.coords())
			count++
			queue.push(s)
		}
	}
	return count, nil
}
