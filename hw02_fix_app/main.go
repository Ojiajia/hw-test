package main

import (
	"fmt"

	"github.com/Ojiajia/hw-test/tree/hw02_fix_app/hw02_fix_app/printer"
	"github.com/Ojiajia/hw-test/tree/hw02_fix_app/hw02_fix_app/reader"
	"github.com/Ojiajia/hw-test/tree/hw02_fix_app/hw02_fix_app/types"
)

func main() {
	// либо это:
	path := "data.json"

	// либо это:

	// fmt.Printf("Enter data file path: ")
	// fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
