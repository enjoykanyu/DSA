package greedy

// 贪心 [0,0,1,2,3,4,5,6,7,8]
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 0
	hashset := make(map[int]struct{})
	count := 1
	for _, value := range nums {
		hashset[value] = struct{}{}
	}
	//从小遍历到大 从大遍历到小
	for value := range hashset {
		if _, ok := hashset[value-1]; ok {
			continue
		}
		cur := value + 1
		for {
			if _, ok := hashset[cur]; !ok {
				break
			}
			cur++
			count++
		}
		if count > result {
			result = count
		}
		count = 1
	}
	if count > result {
		result = count
	}

	return result
}
