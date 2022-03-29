package main

import (
	"fmt"
	"math"
)

func convertDegtoRad(angle float64) float64 {
	// 2 pi = 360 deg
	// 1 deg = 2pi/360
	// angle * pi/180
	return angle * (math.Pi / 180)
}
func calculateSide(length1, length2, angle float64) float64 {

	//Insert the code here
	angle = convertDegtoRad(angle)
	length3 := math.Pow(math.Pow(length1, 2)+math.Pow(length2, 2)-2*length1*length2*math.Cos(angle), 0.5)

	//Do not remove this line
	fmt.Println("The 3rd length of the triange is", length3)
	return length3
}

func main() {
	var length1 float64
	var length2 float64
	var angle float64

	fmt.Println("Enter the first length of the triangle: ")
	fmt.Scanln(&length1)
	fmt.Println("Enter the second length of the triangle: ")
	fmt.Scanln(&length2)
	fmt.Println("Enter the angle between the two lengths: ")
	fmt.Scanln(&angle)

	calculateSide(length1, length2, angle)
}
