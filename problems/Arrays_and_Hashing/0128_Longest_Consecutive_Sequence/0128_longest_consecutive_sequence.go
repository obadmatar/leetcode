package problems

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool, len(nums))

	// inti num set
	for _, n := range nums {
		numSet[n] = true
	}

	longest := 0
	for n := range numSet {
		delete(numSet, n)

		count := 1

		for left := n - 1; numSet[left]; {
			delete(numSet, left)
			count++
			left--
		}

		for right := n + 1; numSet[right]; {
			delete(numSet, right)
			count++
			right++
		}
		longest = max(count, longest)
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
