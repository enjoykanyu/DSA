package greedy

func jump(nums []int) int {
	curRight := 0  //当前最右侧 有变更说明跳跃了一次
	nextRight := 0 //维护的可以走到的最右侧
	value := 0
	for i := 0; i < len(nums)-1; i++ { //只遍历到len-2 这时候会判断最右侧端点是否大于当前位置 <=则需多跳跃一次 而多跳跃一次题目保证肯定会到达
		nextRight = max(nextRight, i+nums[i])
		if i == curRight {
			curRight = nextRight
			value++
		}
	}
	return value
}
