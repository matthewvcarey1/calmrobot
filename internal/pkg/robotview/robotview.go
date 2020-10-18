package robotview

import (
	"container/list"
	"errors"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
)

// RobotView type
type RobotView struct {
	safe int
	land *mapland.MapLand
}

//const (
// The safe threshold
//	safe = 23
//)

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
func (rv *RobotView) isSafe(x int, y int) bool {
	digits := getDigits(x)
	digits = append(digits, getDigits(y)...)
	return (sum(digits) <= rv.safe)
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

// Mark all the mines in the land
func (r *RobotView) MarkMines() {
	x1, y1, x2, y2 := r.land.GetStartEnd()
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			if !r.isSafe(x, y) {
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

// returns the coords of this robot
// used in function parameters
func (r robot) coords() (int, int) {
	return r.x, r.y
}

/*
type robotQueue struct {
	queue []robot
}

func (q *robotQueue) size() int {
	return len(q.queue)
}

func (q *robotQueue) push(r robot) {
	q.queue = append(q.queue, r)
}

func (q *robotQueue) pop() robot {
	r := q.queue[0]       // first element
	q.queue = q.queue[1:] //Dequeue
	return r
}

func newRobotQueue() *robotQueue {
	return &robotQueue{queue: []robot{}}
}
*/

type robotQueue struct {
	queue *list.List
}

func (q robotQueue) size() int {
	return q.queue.Len()
}

func (q robotQueue) push(r robot) {
	q.queue.PushBack(r)
}

func (q robotQueue) pop() robot {
	e := q.queue.Front() // first element
	r := e.Value.(robot)
	q.queue.Remove(e)
	return r
}

func newRobotQueue() robotQueue {
	q := list.New()
	return robotQueue{queue: q}
}

// FloodFill marks the available non mined space as accessable
// by the robot and returns the number of coordinates it contains
func (rv *RobotView) FloodFill(x int, y int) (int, error) {
	queue := newRobotQueue()
	x1, y1, x2, y2 := rv.land.GetStartEnd()
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
		if r.x == x1 || r.x == x2-1 || r.y == y1 || r.y == y2-1 {
			return 0, errors.New("FloodFill has reached the edge of the world, the world needs to be bigger")
		}
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
