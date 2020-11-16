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

var verbose = flag.Bool("verbose", false, "prints large maps of the land before and after the flood fill, best pipe to a file")
var images = flag.Bool("images", false, "generates png files")
var safe = flag.Int("safe", 23, "safe value for robot")

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Now().Sub(t1))
	}()
	flag.Parse()
	// 1400x1400 is the smallest world that will work (699 * 2 + 2)
	// 6 + 9 + 9 = 24 even one coordinate will cause a mine
	land := mapland.New(*safe)
	rv := robotview.New(*safe, land)
	// Mark all the mines on the map
	rv.MarkMines()
	if *verbose {
		land.Draw()
	}
	if *images {
		land.DrawImage("mines.png")
	}
	// Flood fill the from the centre of the map
	// the accessable coordinates
	count, err := rv.FloodFill(0, 0)
	if err != nil {
		// An error about the flood reaching the edge of the world
		println("Error:", err.Error())
		return
	}
	axisfill := mapland.CalcMaplandSize(*safe) - 1
	count = (count-axisfill)*quadrants + origin
	if *verbose {
		fmt.Println()
		land.Draw()
	}
	if *images {
		land.DrawImage("robot.png")
	}
	fmt.Println("The robot can access", count, "coordinates")

}
