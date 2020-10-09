package main

import (
	"flag"
	"fmt"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
	"github.com/matthewvcarey1/calmrobot/internal/pkg/robotview"
)

var verbose = flag.Bool("verbose", false, "prints large maps of the land before and after the flood fill, best pipe to a file")

func main() {
	flag.Parse()
	// 789 is chosen as 7+8+9=24
	// Every coordinate with 789 in will be a mine
	// Thus there will be complete enclosure of mines
	land := mapland.New(789*2 + 1)
	rv := robotview.New(land)
	// Mark all the mines on the map
	rv.MarkMines()
	if *verbose {
		land.Draw()
	}
	// Flood fill the from the centre of the map
	// the accessable coordinates
	count := rv.FloodFill(0, 0)
	if *verbose {
		fmt.Println()
		land.Draw()
	}
	fmt.Println("The robot can access", count, "coordinates")

}
