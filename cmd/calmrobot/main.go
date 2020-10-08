package main

import (
	"fmt"

	"github.com/matthewvcarey1/calmrobot/internal/pkg/mapland"
	"github.com/matthewvcarey1/calmrobot/internal/pkg/robot"
)

func main() {
	fmt.Println("Hello Robot")
	land := mapland.New(100)
	robot = robot.New(land)
	robot.DrawMapLand()
	fmt.Println(land)
}
