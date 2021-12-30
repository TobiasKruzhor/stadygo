package main

import (
	"fmt"
	"strings"
)

func contains() {
	fmt.Println("You find yaourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	// 3-1
	// var wearShades bool
	// var read = strings.Contains(command, "read")
}
