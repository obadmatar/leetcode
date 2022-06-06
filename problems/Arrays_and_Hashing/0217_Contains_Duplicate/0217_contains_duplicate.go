package problems

func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, n := range nums {
		if visited[n] {
			return true
		}
		visited[n] = true
	}
	return false
}
