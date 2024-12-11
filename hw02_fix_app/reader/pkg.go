package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/Ojiajia/hw-test/tree/hw02_fix_app/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	vsbyten, err := io.ReadAll(f)
	if err != nil {
		fmt.Printf("Error: %v", err)
		// return nil, nil
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal( // from JSON 2 &data
		vsbyten,
		&data,
	)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	// res := data

	return data, err
}
