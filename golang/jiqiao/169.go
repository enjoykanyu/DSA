package jiqiao

// map解决
func majorityElement(nums []int) int {

	count := 0           //最大数次数
	cur := map[int]int{} //当前最大数
	value := 0           //返回
	for i := 0; i < len(nums); i++ {
		cur[nums[i]]++
		if count < cur[nums[i]] {
			count = cur[nums[i]]
			value = nums[i]
		}
	}
	return value
}
