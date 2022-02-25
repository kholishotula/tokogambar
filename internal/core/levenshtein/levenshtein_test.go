package levenshtein

import "testing"

var testCases = []struct {
	source   string
	target   string
	distance int
}{
	{
		source:   "",
		target:   "a",
		distance: 1,
	},
	{
		source:   "a",
		target:   "aa",
		distance: 1,
	},
	{
		source:   "a",
		target:   "aaa",
		distance: 2,
	},
	{
		source:   "",
		target:   "",
		distance: 0,
	},
	{
		source:   "a",
		target:   "b",
		distance: 1,
	},
	{
		source:   "aaa",
		target:   "aba",
		distance: 1,
	},
	{
		source:   "aaa",
		target:   "ab",
		distance: 2,
	},
	{
		source:   "ab",
		target:   "ab",
		distance: 0,
	},
	{
		source:   "a",
		target:   "",
		distance: 1,
	},
	{
		source:   "kitten",
		target:   "sitting",
		distance: 3,
	},
	{
		source:   "Orange",
		target:   "Apple",
		distance: 5,
	},
	{
		source:   "ab",
		target:   "bc",
		distance: 2,
	},
	{
		source:   "abd",
		target:   "bec",
		distance: 3,
	},
}

func TestDistanceTwoStrings(t *testing.T) {
	for _, testCase := range testCases {
		distance := DistanceTwoStrings(
			testCase.source,
			testCase.target,
		)
		if distance != testCase.distance {
			t.Log(
				"Distance between",
				testCase.source,
				"and",
				testCase.target,
				"computed as",
				distance,
				". It should be",
				testCase.distance,
			)
			t.Fail()
		}
	}
}
