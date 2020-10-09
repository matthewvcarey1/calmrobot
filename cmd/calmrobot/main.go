package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
	"github.com/matthewvcarey1/calmrobot/internal/pkg/robotview"
)

var verbose = flag.Bool("verbose", false, "prints large maps of the land before and after the flood fill, best pipe to a file")

func main() {
	t1 := time.Now()
	defer func() {
		fmt.Println("Time taken:", time.Now().Sub(t1))
	}()
	flag.Parse()
	// 1400x1400 is the smallest world that will work (by trial and error)
	land := mapland.New(1400)
	rv := robotview.New(land)
	// Mark all the mines on the map
	rv.MarkMines()
	if *verbose {
		land.Draw()
	}
	// Flood fill the from the centre of the map
	// the accessable coordinates
	count, err := rv.FloodFill(0, 0)
	if err != nil {
		// An error about the flood reaching the edge of the world
		println("Error:", err.Error())
		return
	}
	if *verbose {
		fmt.Println()
		land.Draw()
	}
	fmt.Println("The robot can access", count, "coordinates")

}
