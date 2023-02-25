package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	var TestsForSuccess = []struct {
		input   []string
		options Options
		output  []string
	}{
		{[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
			Options{false, false, false, false, 0, 0},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik."}},

		{[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
			Options{true, false, false, false, 0, 0},
			[]string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik."}},

		{[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
			Options{false, true, false, false, 0, 0},
			[]string{
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik."}},

		{[]string{
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
			Options{false, false, true, false, 0, 0},
			[]string{
				"",
				"Thanks."}},

		{[]string{
			"I LOVE music.",
			"I love music.",
			"I love music.",
			"",
			"I Love music Of kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I LoVe music of Kartik."},
			Options{false, false, false, true, 0, 0},
			[]string{
				"I LOVE music.",
				"",
				"I Love music Of kartik.",
				"Thanks.",
				"I love music of Kartik."}},

		{[]string{
			"I loving music.",
			"We love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"I love musician of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I listen music of Kartik."},
			Options{false, false, false, false, 2, 0},
			[]string{
				"I loving music.",
				"",
				"I love music of Kartik.",
				"I love musician of Kartik.",
				"Thanks.",
				"I love music of Kartik."}},

		{[]string{
			"I love music.",
			"A love music.",
			"S love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik."},
			Options{false, false, false, false, 0, 1},
			[]string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik."}},

		{[]string{
			"I loving music.",
			"We love Tusic.",
			"They love Rusic.",
			"",
			"I love music of Kartik.",
			"I love musician of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I listen music of Kartik."},
			Options{false, false, false, false, 2, 1},
			[]string{
				"I loving music.",
				"",
				"I love music of Kartik.",
				"I love musician of Kartik.",
				"Thanks.",
				"I love music of Kartik."}},
	}

	for _, test := range TestsForSuccess {
		res, err := Uniq(test.input, test.options)
		require.Equal(t, nil, err)
		require.Equal(t, test.output, res)
	}

	var TestForError = []struct {
		input   []string
		options Options
		output  []string
	}{
		{[]string{
			"I love music.",
			"I love music.",
			"I love music of Kartik."},
			Options{true, true, false, false, 0, 0},
			[]string{}},
		
		{[]string{
			"I love music.",
			"I love music.",
			"I love music of Kartik."},
			Options{true, false, true, false, 0, 0},
			[]string{}},

		{[]string{
			"I love music.",
			"I love music.",
			"I love music of Kartik."},
			Options{false, true, true, false, 0, 0},
			[]string{}},
		
		{[]string{
			"I love music.",
			"I love music.",
			"I love music of Kartik."},
			Options{true, true, true, false, 0, 0},
			[]string{}},
	}

	for _, test := range TestForError {
		_, err := Uniq(test.input, test.options)
		require.Error(t, err)
	}
}
