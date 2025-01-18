package main

import "fmt"

func main() {

var m,n int

fmt.Print("Type two numbers:")

fmt.Scan(&m, &n)

for i := 0; i < m; i++ {

	for j := 0; j < n; j++ {

	if ((j % 2) == 0) && ((i % 2) == 0) {

		fmt.Printf("\x23")

		} else if ((j % 2) != 0) && ((i % 2) != 0) {

			fmt.Printf("\x23")

		} else {

				fmt.Printf(" ")

		}

	}

fmt.Printf("\n")

}

}
