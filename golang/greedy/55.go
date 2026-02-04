package greedy

func canJump(nums []int) bool {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if i > cur {
			return false
		}
		cur = max(cur, i+nums[i]) //当前能够到达的距离
	}
	return true
}
