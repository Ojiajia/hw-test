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

func (booBook *Book) Get(Book) {

	print("Id: ", booBook.ID, " Title: ", booBook.Title, " Author: ", booBook.Author, " Year: ", booBook.Year, " Size: ", booBook.Size, " Rate: ", booBook.Rate)

}

func main() {

	print("Input your Book data: ID, Title, Author, Year, Size, Rate\n")

	var A, D int
	var B, C string
	var E int
	var F float32

	fmt.Scan(&A, &B, &C, &D, &E, &F)

	var BiBook Book

	BiBook.Set(A, B, C, D, E, F)

	BiBook.Get(BiBook)
}
