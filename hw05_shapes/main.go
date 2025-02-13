package main

import (
	"math"
)

type Circle struct { // круг
	R float32 // R - радиус
}
type Triangle struct { // треугольник
	H, L float32 // H - высота, L - длина
}
type Rectangle struct { // прямоугольник
	A, B float32 // A - длина, B - ширина
}

type Area interface {
	CalcArea() float32
}

//**********

func (c Circle) CalcArea() float32 {
	print("Circle Area =\n")
	return (math.Pi) * (c.R) * (c.R)
}

func (t Triangle) CalcArea() float32 {
	print("Triangle Area =\n")
	return t.H * t.L / 2
}

func (r Rectangle) CalcArea() float32 {
	print("Rectangle Area =\n")
	return r.A * r.B
}

//**********

func CallInterface(a Area) {

	print(a.CalcArea())

}

func main() {

	var c Circle
	c.R = 3.
	CallInterface(c)

	var t Triangle
	t.H = 2.
	t.L = 3.
	CallInterface(t)

	var r Rectangle
	r.A = 5.
	r.B = 4.
	CallInterface(r)

}
