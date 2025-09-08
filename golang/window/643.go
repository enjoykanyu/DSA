package window

// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。

// 请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。

// 任何误差小于 10-5 的答案都将被视为正确答案。

// 示例 1：

// 输入：nums = [1,12,-5,-6,50,3], k = 4
// 输出：12.75
// 解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
// 示例 2：

// 输入：nums = [5], k = 1
// 输出：5.00000

func findMaxAverage(nums []int, k int) float64 {
	resutl := -1 << 31 //当前平均值最小
	cur_add := 0
	//第一次直接计算
	for i := 0; i < k; i++ {
		cur_add += nums[i]
	}
	if cur_add > resutl {
		resutl = cur_add

	}
	for i := 1; i < len(nums)-k+1; i++ {
		cur_add = cur_add - nums[i-1] + nums[i+k-1]
		if cur_add > resutl {
			resutl = cur_add
		}
	}
	c := float64(k)
	return float64(resutl) / c
}
