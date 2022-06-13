package problems

import (
	"fmt"
	"testing"
	"golang.org/x/exp/slices"
)

type question struct {
	input  []string
	output [][]string
}

var questions = []question{
	{
		[]string{"eat","tea","tan","ate","nat","bat"},
		[][]string{{"eat","tea","ate"}, {"tan","nat"}, {"bat"}},
	},
	{
		[]string{""},
		[][]string {{""}},
	},
	{
		[]string{"bdddddddddd","bbbbbbbbbbc"},
		[][]string{{"bdddddddddd"}, {"bbbbbbbbbbc"}},
	},
	{
		[]string{"a"},	
		[][]string {{"a"}},
	},
}

func Test_ContainsDuplicate(t *testing.T) {
	fmt.Println("--- 0049. Group Anagrams ")

	for _, q := range questions {
		expected := q.output
		actual := groupAnagrams(q.input)
		if !equal(expected, actual) {
			t.Errorf("[input]: %v\n, [expected]: %v\n, [actual]: %v\n", q.input, q.output, actual)
		}
	}
}

func equal(expected, actual [][]string) bool {
	if actual == nil || len(expected) != len(actual) {
		return false
	}
	for i, _ := range expected {

		if 	!slices.Equal(expected[i], actual[i]) {
			return false
		}
	}
	return true
}
