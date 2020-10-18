package calculator_test

import (
	"testing"

	"github.com/leowmjw/go-course/FTLOG/calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

// func TestSubtract(t *testing.T) {
// 	t.Parallel()
// 	var want float64 = 2
// 	got := calculator.Subtract(4, 2)
// 	if want != got {
// 		t.Errorf("want %f, got %f", want, got)
// 	}
// }
