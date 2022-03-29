package main

import "testing"

func TestTriangleCalculator(t *testing.T) {

	// get := calculateSide(10, 10, 45)
	// res := 9.74349024921019 // answer is wrong shd be 7.653668647301795
	get := calculateSide(3, 4, 90)
	res := 5.0
	if get != res {
		t.Errorf("CalculateSide() = %f Want %f Got %f", get, res, get)
	}
}
