package robot

import (
	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
)

type Robot struct {
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

func New(land *mapland.MapLand) *Robot {
	return &Robot{land: land}
}

type pair struct {
	x int
	y int
}

func (r *Robot) MarkMines() {
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

func east(p pair)pair{
	return pair{p.x+1, p.y}
}

func west(p pair) pair{
	return  pair{p.x-1, p.y}
}

func south(p pair)pair{
	return pair{p.x, p.y-1}
}

func north(p pair)pair{
	return pair{p.x, p.y+1}
}


func (r *Robot) FloodFill(x int, y int) int {
	var queue []pair
	count := int(0)
	if r.land.Get(x, y) != 0x20 {
		return 0
	}
	r.land.Set(x, y)
	count++
	queue = append(queue, pair{
		x: x,
		y: y,
	})
	for len(queue) > 0 {
		p := queue[0]     // first element
		queue = queue[1:] //Dequeue
		if w:=west(p); r.land.Get(w.x,w.y) == 0x20 {
			r.land.Set(w.x, w.y)
			count++
			queue = append(queue, w)
		}
		if e:=east(p); r.land.Get(e.x, e.y) == 0x20 {
			r.land.Set(e.x, e.y)
			count++
			queue = append(queue, e)
		}
		if n:=north(p); r.land.Get(n.x, n.y) == 0x20 {
			r.land.Set(n.x, n.y)
			count++
			queue = append(queue, n)
		}
		if s:=south(p); r.land.Get(s.x, s.y) == 0x20 {
			r.land.Set(s.x, s.y)
			count++
			queue = append(queue, s)
		}
	}
	return count
}
