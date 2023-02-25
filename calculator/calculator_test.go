package calculator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculator(t *testing.T) {
	var TestsForSuccess = []struct {
		input  string
		output float64
	}{
		{"1+1", 2},
		{"1*1", 1},
		{"1/2", 0.5},
		{"-1", -1},
		{"1", 1},
		{"1*2+(3+4*(5+6))", 49},
		{"((6+5)*4+3)+2*1", 49},
		{"2*(-3)", -6},
		{"((1+2)+3/4)*5", 18.75},
		{"1000*3/5/6", 100},
		{"1*2+3*4*(1+2)/3", 14},
		{"1/(-1)", -1},
		{"-10+2-3-4", -15},
		{"123/3-(5*6/(10-(2*2*2)))", 26},
		{"123/3-(-5*6/(10-(2*2*2)))", 56},
		{"1.2+1", 2.2},
		{"4.5/1.5", 3},
		{"4.5/3", 1.5},
		{"100/0.2", 500},
		{"3/0.1*0.01", 0.3},
	}

	for _, test := range TestsForSuccess {
		res, err := Calculate(Parse(test.input))
		require.Equal(t, nil, err)
		require.Equal(t, test.output, res)
	}

	var TestForError = []struct {
		input  string
		output float64
	}{
		{"", 0},
		{"1a3", 0},
		{"qqq", 0},
		{"1++1", 0},
		{"+", 0},
		{")))", 0},
		{"*1-5", 0},
		{"(((1+5))", 0},
		{"2+3+()", 0},
		{"(3+1)(5-4)", 0},
		{"5/0", 0},
		{".1+5", 0},
		{"1..4+3", 0},
		{"1.+2", 0},
		{"1+.2", 0},
	}

	for _, test := range TestForError {
		res, err := Calculate(Parse(test.input))
		require.Error(t, fmt.Errorf("wrong data"), err, fmt.Sprintf("Calculate(%s) = %f, expected error",
			test.input, res))
	}
}
