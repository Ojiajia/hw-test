package main

import (
	"fmt"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

func (booBook *Book) Set(bId int, bTitle, bAuth string, bYear int, bSize int, bRate float32) *Book {
	booBook.ID = bId
	booBook.Title = bTitle
	booBook.Author = bAuth
	booBook.Year = bYear
	booBook.Size = bSize
	booBook.Rate = bRate
	return booBook
}

func (booBook Book) Get(Book) {

	print("Id: ", booBook.ID, " Title: ", booBook.Title, " Author: ", booBook.Author, " Year: ", booBook.Year, " Size: ", booBook.Size, " Rate: ", booBook.Rate, "\n")

}

type Mode int

const (
	YearMode Mode = iota + 1
	SizeMode
	RateMode
)

type ModeVariety struct {
	Choice Mode
}

func (m ModeVariety) ToCompare(first Book, second Book) bool {

	switch m.Choice {
	case 1:
		print("you choose Year\n")
		if first.Year > second.Year {
			print("true\n")
			return true
		} else {
			print("false\n")
			return false
		}
	case 2:
		print("you choose Size\n")
		if first.Size > second.Size {
			print("true\n")
			return true
		} else {
			print("false\n")
			return false
		}
	case 3:
		print("you choose Rate\n")
		if first.Rate > second.Rate {
			print("true\n")
			return true
		} else {
			print("false\n")
			return false
		}
	default:
		print("something's wrong\n")

	}

	return true // ?

}

func main() {

	print("Input your first Book data: ID, Title, Author, Year, Size, Rate\n")

	var A, D int
	var B, C string
	var E int
	var F float32

	fmt.Scan(&A, &B, &C, &D, &E, &F)

	var BookOne Book

	BookOne.Set(A, B, C, D, E, F)

	A = 0
	B = ""
	C = ""
	D = 0
	E = 0
	F = 0.

	print("Input your second Book data: ID, Title, Author, Year, Size, Rate\n")

	fmt.Scan(&A, &B, &C, &D, &E, &F)

	var BookTwo Book

	BookTwo.Set(A, B, C, D, E, F)

	var CompData ModeVariety
	print("Input mode for comparison (1 = Year, 2 = Size, 3 = Rate)\n")
	fmt.Scan(&CompData.Choice)

	CompData.ToCompare(BookOne, BookTwo)

}
