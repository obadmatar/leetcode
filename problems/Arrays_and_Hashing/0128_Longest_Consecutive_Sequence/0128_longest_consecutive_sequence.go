package problems

func longestConsecutive(nums []int) int {
	max := 0
	set := make(map[int]bool, len(nums))

	for _, n := range nums {
		set[n] = true
	}

	for _, n := range nums {
		if set[n-1] {
			continue
		}

		length := 0
		for set[n + length] {
			length++
			if length > max {
				max = length
			}
		}
	}

	return max
}
