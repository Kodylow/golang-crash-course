package main

import "fmt"

func main() {
	x, y := 15, 10
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	color := "yellow"

	if color == "yellow" {
		fmt.Printf("Color is %s but not brown", color)
	} else if color == "brown" {
		fmt.Printf("Color is %s but not yellow", color)
	} else {
		fmt.Printf("Color is %s but not yellow or brown", color)
	}

	switch color {
	case "yellow":
		fmt.Printf("Color is %s but not brown", color)
	case "brown":
		fmt.Printf("Color is %s but not yellow", color)
	default:
		fmt.Printf("Color is %s but not yellow or brown", color)

	}
}
