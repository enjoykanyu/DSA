package shuzu

func twoSum(nums []int, target int) []int {
	map_cur := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cur := target - nums[i]
		if value, ok := map_cur[cur]; ok {
			return []int{i, value}
		}
		map_cur[nums[i]] = i
	}
	return nil
}
