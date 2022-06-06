package problems

import (
	"fmt"
	"testing"
)

type question struct {
	input  []int
	output bool
}

var questions = []question{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func Test_ContainsDuplicate(t *testing.T) {
	fmt.Println("--- 0217. Contains Duplicate ")

	for _, q := range questions {
		expected := q.output
		actual := containsDuplicate(q.input)
		if expected != actual {
			t.Errorf("[input]: %v, [expected]: %v, [actual]: %v", q.input, q.output, actual)
		}
	}
}
