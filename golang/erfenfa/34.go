package erfenfa

/*
34. 在排序数组中查找元素的第一个和最后一个位置
中等
相关标签
premium lock icon
相关企业
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6 找到<target的数值
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

提示：

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
func searchRange(nums []int, target int) []int {
	result_left := subRange(nums, target)
	if result_left == len(nums) || nums[result_left] != target {
		return []int{-1, -1}
	}
	result_right := subRange(nums, target+1) - 1

	return []int{result_left, result_right}
}

func subRange(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := left + (right-left)/2
	//找小于target的最大值
	for left < right {
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			//大于等于不符合预期
			right = mid
		}
	}

	return left
}
