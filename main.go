package main

import (
	"fmt"
	"github.com/HandsomeJeff/aesthetically-pleasing-trees/calculations"
)

func main() {
	trees := []int{3,4,5,3,7}
	solution := calculations.Solution(trees)

	fmt.Println(solution)
}
