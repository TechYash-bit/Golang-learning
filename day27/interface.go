package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func (c circle) perimeter() float64{
	return math.Pi*c.radius*2
}

type rectangle struct {
	width, length float64
}

func (r rectangle) area() float64{
	return r.length*r.width
}
func (r rectangle) perimeter() float64{
	return 2*(r.length+r.width)
}


func display(g geometry){
fmt.Println("Area : ",g.area())
fmt.Println("Perimeter : ",g.perimeter())
}

func main() {

	c:=circle{
		radius: 10,
	}
	display(c)
	r:=rectangle{
		width: 10,
		length: 12,
	}
	display(r)
}