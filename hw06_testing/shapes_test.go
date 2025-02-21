package shapes_test

import (
	// "fmt"
	"fmt"
	"testing"
	// "errors"
)

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func init() {
	fmt.Println("Hello from init hw05_shapes!")
}

// "/home/okorostelina/go_otus/hw-test/hw04_struct_comparator"
// "/home/okorostelina/go_otus/hw-test/hw03_chessboard"
// "/home/okorostelina/go_otus/hw-test/hw02_fix_app"
// )
func TestShapes(t *testing.T) {

}

func main() {

}
