package shapes

import (
	"fmt"
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

func (c Circle) CalcArea() float32 {
	return (math.Pi) * (c.R) * (c.R)
}

func (t Triangle) CalcArea() float32 {
	return t.H * t.L / 2
}

func (r Rectangle) CalcArea() float32 {
	return r.A * r.B
}

func CallInterface(a Area) {
	print(a.CalcArea())
}

func init() {
	fmt.Println("Hello from init hw05_shapes!")
}

func New(text string) error {
	return &errorString{text}
   }
   type errorString struct {
	s string
   }
   func (e *errorString) Error() string {
	return e.s
   }

func main() {

	
	var c Circle
	c.R = 3.
	print("Circle Area = ")
	CallInterface(c)

	var t Triangle
	t.H = 2.
	t.L = 3.
	print("\nTriangle Area = ")
	CallInterface(t)

	var r Rectangle
	r.A = 5.
	r.B = 4.
	print("\nRectangle Area = ")
	CallInterface(r)

	shapes := []Area{c, t, r}
	for _, shape := range shapes {
		switch v := shape.(type) {
		case Circle:
			print("\nit's a circle, area = ", v.CalcArea())
		case Triangle:
			print("\nit's a triangle, area = ", v.CalcArea())
		case Rectangle:
			print("\nit's a rectangle, area = ", v.CalcArea(), "\n")
		default:
			print("\nuknown shape type\n")
		}
	}
}
