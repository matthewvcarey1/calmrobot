package main

import (
	"fmt"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
	"github.com/matthewvcarey1/calmrobot/internal/pkg/robot"
)

func main() {
	// 789 is chosen as 7+8+9=24
	// Every coordinate with 789 in will be a mine
	land := mapland.New(789*2 + 1)
	robot := robot.New(land)
	robot.MarkMines()
	land.Draw()
	// Flood fill the from the centre of the map
	count := robot.FloodFill(0, 0)
	fmt.Println("Count", count)
	land.Draw()
	//fmt.Println("Count", land.Count())
	fmt.Println("The robot can access", count, "coordinates")

}
