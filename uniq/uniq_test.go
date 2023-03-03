package uniq

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	var TestsForUniqFunc = []struct {
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
	for _, test := range TestsForUniqFunc {
		res, err := Uniq(test.input, test.options)
		require.Equal(t, nil, err)
		require.Equal(t, test.output, res)
	}

	var TestSkipFields = []struct {
		input     string
		numFields int
		output    string
	}{
		{"I love music of Kartik.", 0, "I love music of Kartik."},
		{"I love music of Kartik.", 1, "love music of Kartik."},
		{"I love music of Kartik.", 2, "music of Kartik."},
	}
	for _, test := range TestSkipFields {
		res := skipFields(test.input, test.numFields)
		require.Equal(t, test.output, res)
	}

	var TestSkipChars = []struct {
		input    string
		numChars int
		output   string
	}{
		{"I love music of Kartik.", 0, "I love music of Kartik."},
		{"I love music of Kartik.", 1, " love music of Kartik."},
		{"I love music of Kartik.", 2, "love music of Kartik."},
		{"I love music of Kartik.", 4, "ve music of Kartik."},
	}
	for _, test := range TestSkipChars {
		res := skipChars(test.input, test.numChars)
		require.Equal(t, test.output, res)
	}

	var TestWordHandler = []struct {
		input   string
		options Options
		output  string
	}{
		{"I love MUSIC of Kartik.", Options{false, false, false, true, 0, 0}, "i love music of kartik."},
		{"I love music of Kartik.", Options{false, false, false, false, 2, 0}, "music of Kartik."},
		{"I love music of Kartik.", Options{false, false, false, false, 0, 3}, "ove music of Kartik."},
		{"I love MUSIC of Kartik.", Options{false, false, false, true, 1, 2}, "ve music of kartik."},
	}
	for _, test := range TestWordHandler {
		res := wordHandler(test.input, test.options)
		require.Equal(t, test.output, res)
	}

	var TestWithRepeat = []struct {
		input  string
		repeat int
		output string
	}{
		{"I love music of Kartik.", 1, "1 I love music of Kartik."},
		{"I love music of Kartik.", 2, "2 I love music of Kartik."},
	}
	for _, test := range TestWithRepeat {
		res := withRepeat(test.input, test.repeat)
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

		{[]string{},
			Options{false, false, false, false, 0, 0},
			[]string{}},
	}
	for _, test := range TestForError {
		_, err := Uniq(test.input, test.options)
		require.Error(t, err)
	}
}
