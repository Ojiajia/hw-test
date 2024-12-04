package printer

import (
	"fmt"

	"github.com/Ojiajia/hw-test/tree/hw02_fix_app/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var string1, string2 string
	for i := 0; i < len(staff); i++ {
		string1 = fmt.Sprintf("User ID: %d; Age: %d ", staff[i].UserID, staff[i].Age)

		string2 = fmt.Sprintf("Name: %s; Department ID: %d; ", staff[i].Name, staff[i].DepartmentID)

		fmt.Println(string1, string2)
	}

	fmt.Println(string1, string2)
}
