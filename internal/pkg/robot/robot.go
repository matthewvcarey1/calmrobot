package robot

import (
	"fmt"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
)

type Robot struct {
	x    int
	y    int
	land *mapland.Mapland
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
	return &Robot{x: 0, y: 0, land & mapland.MapLand}
}

func (r *Robot) WalkWest() bool {
	if isSafe(r.x-1, r.y) {
		r.x -= 1
		r.land.Set(r.x, r.y)
		return true
	}
	return false
}

func (r *Robot) WalkEast() bool {
	if isSafe(r.x+1, r.y) {
		r.x += 1
		r.land.Set(r.x, r.y)
		return true
	}
	return false
}

func (r *Robot) WalkNorth() bool {
	if isSafe(r.x, r.y+1) {
		r.y += 1
		r.land.Set(r.x, r.y)
		return true
	}
	return false
}

func (r *Robot) WalkSouth() bool {
	if isSafe(r.x, r.y-1) {
		r.y -= 1
		r.land.Set(r.x, r.y)
		return true
	}
	return false
}

func (r *Robot) DrawMapLand() {
	x1, y1, x2, y2 := r.land.GetMinMax()
	for x := x1; x < x2; x++ {
		for y := y1; y < x2; y++ {
			if !isSafe(x, y) {
				fmt.Print("X")
			} else {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}
