package zhan

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
// 示例 1:
//
// 输入：heights = [2,1,5,6,2,3]
// 输出：10
// 解释：最大的矩形为图中红色区域，面积为 10
// 示例 2：
//
// 输入： heights = [2,4]
// 输出： 4
// 思路：大于一直入栈 同时计算栈内最里边的高度*总栈长度 哪个最大 当前栈高度大于乘积 出栈 两个栈 栈一存储从左到右的顺序 栈二记录栈内最小的高度 小于出栈 对比站内最小的乘以栈长度
// 参考答案思路 矩形面积的高度必为数组中的元素 面积可以换算成 该数组元素左右边找比最近它小的元素
// 遍历当前元素数组 左边最小 右边最小
// [2,1,5,6,2,3] 2 最左边无比它小的
// 2为栈顶 同时记录2的索引0 最左为-1   栈内[2]   栈存储当前单调递增的数值 同时 数组存储 当前遍历数组的最左值
// 1 左边无比它小的 比它大 1入栈 栈顶出栈  最小的-1 栈内[1]
// 5 左边对比栈顶1 比它小 存储 1  栈内[1,5]
// 6 左边对比栈顶 5 比它小 存储2 栈内[1,5,6]
// 2 左边对比栈顶 6 比它大 出栈  对比 5 出栈 对比 1 小 存储1 栈内[1,2]
// 3 左边对比栈顶 2 比它小 存储4 栈内[1,2,3]
func LargestRectangleArea(heights []int) int {
	left_stack := []int{-1}            //栈内存储
	index := make([]int, len(heights)) //0位置的数组对应的最左为-1
	for i := 0; i < len(heights); i++ {
		//栈顶元素比当前元素大 循环出栈
		for len(left_stack) > 1 && heights[i] <= heights[left_stack[len(left_stack)-1]] {
			left_stack = left_stack[:len(left_stack)-1] //栈顶出栈
		}
		index[i] = left_stack[len(left_stack)-1] // 更新元素左侧最小位置
		left_stack = append(left_stack, i)       // 当前元素比栈顶元素大 则入栈 栈顶存储索引
		//栈顶元素比当前元素小
	}
	right_stak := []int{len(heights)}   //栈内存储
	index1 := make([]int, len(heights)) //0位置的数组对应的最左为-1
	for i := len(heights) - 1; i >= 0; i-- {
		//栈顶元素比当前元素大 循环出栈
		for len(right_stak) > 1 && heights[i] <= heights[right_stak[len(right_stak)-1]] {
			right_stak = right_stak[:len(right_stak)-1] //栈顶出栈
		}
		index1[i] = right_stak[len(right_stak)-1] // 更新元素左侧最小位置
		right_stak = append(right_stak, i)        // 当前元素比栈顶元素大 则入栈 栈顶存储索引
		//栈顶元素比当前元素小
	}
	result := 0
	for i := 0; i < len(heights); i++ {
		cur := (index1[i] - index[i] - 1) * heights[i]
		if cur > result {
			result = cur
		}
	}
	return result
}
