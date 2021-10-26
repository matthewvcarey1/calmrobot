package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
	"github.com/matthewvcarey1/calmrobot/internal/pkg/robotview"
)

const (
	origin    = 1
	quadrants = 4
)

var images = flag.Bool("images", false, "generates png files")
var safe = flag.Int("safe", 23, "safe value for robot")

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Now().Sub(t1))
	}()
	flag.Parse()
	// 1400x1400 is the smallest world that will work for a safe value of 23 (699 * 2 + 2)
	// 6 + 9 + 9 = 24 means that even one coordinate will cause a mine
	land := mapland.New(*safe)
	rv := robotview.New(*safe, land)
	// Mark all the mines on the map
	rv.MarkMines()

	if *images {
		land.DrawImage("mines.png")
	}
	// Flood fill the from the centre of the map
	// the accessable coordinates
	countInQuadrant, err := rv.FloodFill(0, 0)
	if err != nil {
		// An error about the flood reaching the edge of the world
		println("Error:", err.Error())
		return
	}
	// mapland.CalcMaplandSize(num) adds a padding value of 1
	// that represents the edge of the map. So we remove it.
	// Thus axisfill represents all the accessable coords in one axis.

	// We only caclulate one quadrant as they are all rotations
	// of one another and then we multiply by 4 but we must first
	// strip one axis from each quadrants results.

	// Each quadrant will have two axises but 4 X 2 = 8
	// Which is 4 too many so we have to remove the 4 surplus axises
	// And then add the shared origin back at the end.

	axisfill := mapland.CalcMaplandSize(*safe) - 1
	count := (countInQuadrant-axisfill)*quadrants + origin

	if *images {
		land.DrawImage("robot.png")
	}
	fmt.Println("The robot can access", count, "coordinates")
}
