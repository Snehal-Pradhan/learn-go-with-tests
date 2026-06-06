package main


import "math"

type Shape interface {
	Area() float64
}



type Rectangle struct {
	height float64
	width float64
}


type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return float64(r.width)*float64(r.height)
}

func (c Circle) Area() float64 {
	return math.Pi*float64(c.radius)*float64(c.radius)
}


func Perimeter (r Rectangle) float64 {
	return 2*(r.height+r.width)
}