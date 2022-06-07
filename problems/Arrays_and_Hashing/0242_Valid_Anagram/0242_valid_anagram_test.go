package problems

import (
	"fmt"
	"testing"
)

type question struct {
	source   string
	target   string
	expected bool
}

var questions = []question{
	{"anagram", "nagaram", true},
	{"rat", "car", false},
}

func Test_IsAnagram(t *testing.T) {
	fmt.Println("--- 0242. Valid Anagram ")

	for _, q := range questions {
		expected := q.expected
		actual := isAnagram(q.source, q.target)
		if expected != actual {
			t.Errorf("[input]: {%v, %v}, [expected]: %v, [actual]: %v", q.source, q.target, expected, actual)
		}
	}
}
