package problems

import (
	"fmt"
	"testing"
)

type question struct {
	input  []int
	output int
}

var questions = []question{
	{
		input:  []int{0},
		output: 1,
	},
	{
		input:  []int{100, 4, 200, 1, 3, 2},
		output: 4,
	},
	{
		input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		output: 9,
	},
}

func Test_LongestConsecutive(t *testing.T) {
	fmt.Println("--- 0128. Longest Consecutive Sequence ")

	for _, q := range questions {

		actual := longestConsecutive(q.input)

		if actual != q.output {
			t.Errorf("[input]: %v, [expected]: %v, [actual]: %v", q.input, q.output, actual)
		}
	}
}
