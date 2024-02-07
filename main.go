package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var thisCube bool
	var thisSquare bool

	fmt.Println("================================================")
	fmt.Println("     Mini Challenge 1 - SquareCube by Denir     ")
	fmt.Println("================================================")

	fmt.Printf("Masukkan Nomor = ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		square := math.Sqrt(float64(i))
		if square == math.Floor(square) {
			thisSquare = true
		} else {
			thisSquare = false
		}

		cube := math.Cbrt(float64(i))
		if cube == math.Floor(cube) {
			thisCube = true
		} else {
			thisCube = false
		}

		switch {
		case thisSquare && thisCube:
			fmt.Println("SquareCube")
		case thisSquare:
			fmt.Println("Square")
		case thisCube:
			fmt.Println("Cube")
		default:
			fmt.Println(i)
		}
	}

}
