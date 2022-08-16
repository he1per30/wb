package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func pointIntersection(p1 Point, p2 Point) float64 {
	k1 := p1.x - p2.x
	k2 := p1.y - p2.y
	return math.Sqrt(k1*k1 + k2*k2)
}

func main24() {
	p1 := NewPoint(10, 12)
	p2 := NewPoint(2, 4)
	fmt.Println(pointIntersection(*p1, *p2))
}
