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

//**********

type VarShape struct { // для выбора структуры
	Circle
	Triangle
	Rectangle
}

type Mode int // выбор варианта фигуры для подсчета площади

const (
	CircleMode    Mode = iota + 1 // считаем площадь круга
	TriangleMode                  // считаем площадь треугольника
	RectangleMode                 // считаем площадь прямоугольника
)

type ModeVariety struct {
	Choice Mode
}
type Area interface {
	CalcArea(VarShape) float32
}

//**********

func (m ModeVariety) CalcArea(someShape VarShape) float32 { // - ?

	switch m.Choice { // выберем фигуру для посдчета площади
	case 1:
		print("you choose Circle\n")
		return math.Pi * (someShape.Circle.R) * (someShape.Circle.R)
	case 2:
		print("you choose Triangle\n")
		return someShape.Triangle.H * someShape.Triangle.L / 2
	case 3:
		print("you choose Rectangle\n")
		return someShape.Rectangle.A * someShape.Rectangle.B
	default:
		print("you were wrong")
		return 0
	}
}

//**********

func main() {

	var SomeVar VarShape
	var ChoosedVar ModeVariety
	var a Area

	SomeVar.Circle = Circle{3.}
	ChoosedVar.Choice = 1 // площадь круга

	ChoosedVar.CalcArea(SomeVar)

	a.CalcArea(SomeVar)

	print("Circle with R = 3 area = ", a.CalcArea(SomeVar))
	// print("Circle with R = 3 area = ", a)

	SomeVar.Triangle = Triangle{5., 4.}
	ChoosedVar.Choice = 2 // площадь треугольника
	a.CalcArea(SomeVar)
	print("Triangle with H = 5, L = 4, area = ", a.CalcArea(SomeVar))

	SomeVar.Rectangle = Rectangle{2, 3}
	ChoosedVar.Choice = 3 // площадь прямоугольника
	print("Rectangle with A = 2, B = 3, area = ", a.CalcArea(SomeVar))

}
