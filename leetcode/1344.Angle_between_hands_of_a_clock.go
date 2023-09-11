package main

import (
	"fmt"
	"math"
)

func angleClock(hour int, minutes int) float64 {
	mAngle := float64(minutes * (360 / 60))
	fmt.Println(mAngle)
	hAngle := (float64(hour % 12) + float64(minutes) / 60) * (360 / 12)
	fmt.Println(hAngle)
	diff := math.Abs(mAngle - hAngle)
	return min1344(diff, 360 - diff)
}

func min1344(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func main() {
	hour, minutes := 3, 15
	res := angleClock(hour, minutes)
	fmt.Println(res)
}
