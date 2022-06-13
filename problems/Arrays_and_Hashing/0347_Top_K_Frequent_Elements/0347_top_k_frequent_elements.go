package problems

func topKFrequent(nums []int, k int) []int {
	res := []int{}
	freqs := make([][]int, len(nums))
	count := make(map[int]int, len(nums))

	for _, n := range nums {
		count[n]++
	}

	for n, c := range count {
		freqs[c-1] = append(freqs[c-1], n)
	}

	for i := len(freqs)-1; i >= 0; i-- {
		for _, n := range freqs[i] {
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

