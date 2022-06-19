package problems

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
	"testing"
)

type question struct {
	nums   []int
	k      int
	output []int
}

var questions = []question{
	{
		nums:   []int{1, 1, 1, 2, 2, 3},
		k:      2,
		output: []int{1, 2},
	},
	{
		nums:   []int{1},
		k:      1,
		output: []int{1},
	},
}

func Test_IsAnagram(t *testing.T) {
	fmt.Println("--- 0347. Top K Frequent Elements ")

	for _, q := range questions {
		expected := q.output
		actual := topKFrequent(q.nums, q.k)

		sort.Ints(actual)
		sort.Ints(expected)

		if !slices.Equal(expected, actual) {
			t.Errorf("[input]: {nums: %v, k:%v}, [expected]: %v, [actual]: %v", q.nums, q.k, expected, actual)
		}
	}
}
