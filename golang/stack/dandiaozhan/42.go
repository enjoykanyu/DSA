package dandiaozhan

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：

输入：height = [4,2,0,3,2,5]
输出：9
*/
// 思路 左右遍历单调递增入栈1 有单调递减的入栈2 有比栈2顶小的 循环对比栈1 2
func Trap(height []int) int {
	//stack_right := []int{} //递增栈存储空
	// 往右遍历当高度大于递减栈顶元素高度且递增栈为空 入递增栈   递增栈入栈时机
	//当小于递增栈顶元素高度 计算面积
	//递减栈清空 递增栈清空 递减栈顶存储递增栈顶的元素加上当前的元素位置
	stack_left := []int{} //递减栈存储原始位 当前元素高度小于递减栈顶位置 且递增栈为空 入递减栈
	result := 0
	//参考答案思想 栈内存储单调递减 往右有比当前栈顶大于的直接算差值面积 出栈逐次出栈对比当前的高度距离面积
	for i := 0; i < len(height); i++ {
		for len(stack_left) > 0 && height[i] >= height[stack_left[len(stack_left)-1]] { //出栈操作进行计算面积
			//当前单位面积的底部高度
			dibu_height := height[stack_left[len(stack_left)-1]]
			//出栈
			stack_left = stack_left[:len(stack_left)-1]
			//当前栈内=0 即左边无边框装不了雨水
			if len(stack_left) == 0 {
				break
			}
			//左边框位置
			left := stack_left[len(stack_left)-1]
			//面积高度
			h := min(height[left], height[i]) - dibu_height
			mianji := h * (i - left - 1)
			result += mianji

		}
		stack_left = append(stack_left, i) //入递减栈
	}
	return result
	//for i := 0; i < len(height); i++ {
	//	//left := i //左边的最高值
	//	mianijicha := 0
	//	mianijicha += i
	//	if len(stack_left) == 0 {
	//		stack_left = append(stack_left, i)
	//	} else {
	//		if height[i] <= height[stack_left[len(stack_left)-1]] {
	//			stack_left = append(stack_left, i)
	//
	//		} else { // 右边高于
	//
	//		}
	//	}
	//}
	//for i := 1; i < len(height); i++ {
	//	cur_index := -1
	//	if len(stack_right) == 0 {
	//		cur_index = stack_left[len(stack_left)-1]
	//		if height[i] < height[cur_index] {
	//			stack_left = append(stack_left, i) //循环入递减栈
	//		} else {
	//			stack_right = append(stack_right, i) //有变化入递增栈 后续会进入下边的判断
	//		}
	//	} else {
	//		cur_index = stack_right[len(stack_right)-1]
	//		if height[i] >= height[cur_index] {
	//			stack_right = append(stack_right, i) //循环入递增栈
	//		} else { //此时得进行面积计算和出栈操作
	//			//计算高度
	//			height_ := math.Abs(float64(height[stack_right[len(stack_right)-1]] - height[stack_left[len(stack_left)-1]]))
	//			//计算宽度
	//			width := stack_right[len(stack_right)-1] - stack_left[len(stack_left)-1] + 1
	//			//减去栈内所有的高度
	//
	//			//减去两边栈顶相差值*宽度
	//			//当前遍历面积
	//			//递减栈此时为空
	//			//递增栈此时存储了
	//			//特殊情况两边有相等
	//		}
	//	}
	//	//if len(stack_left)==0 && len(stack_right)>0{
	//	//	continue
	//	//}
	//	//if len(stack_right) > 0 && height[i] < height[stack_left[len(stack_left)-1]] {
	//	//	if height[stack_left[len(stack_left)-1]] > height[stack_right[len(stack_right)-1]] {
	//	//		//计算整体面积
	//	//		width := stack_right[len(stack_right)-1] - stack_left[len(stack_left)-1] + 1
	//	//		heights := height[stack_left[len(stack_left)-1]]
	//	//		cha := height[stack_left[len(stack_left)-1]] - height[stack_right[len(stack_right)-1]]
	//	//		mianji := width * heights
	//	//		mianjicha := cha * (width + 1)
	//	//		result += mianji
	//	//		result -= mianjicha
	//	//		for len(stack_left) > 0 {
	//	//			result -= height[stack_left[len(stack_left)-1]]
	//	//			stack_left = stack_right[:len(stack_right)-1] //出栈
	//	//		}
	//	//		for len(stack_right) > 1 {
	//	//			result -= height[stack_right[len(stack_right)-1]]
	//	//			stack_right = stack_right[:len(stack_right)-1] //出栈
	//	//		}
	//	//		result -= height[stack_right[len(stack_right)-1]]
	//	//	} else {
	//	//		//计算整体面积
	//	//		width := stack_left[len(stack_left)-1] - stack_right[len(stack_right)-1] + 1
	//	//		heights := height[stack_right[len(stack_right)-1]]
	//	//		cha := height[stack_right[len(stack_right)-1]] - height[stack_left[len(stack_left)-1]]
	//	//		mianji := width * heights
	//	//		mianjicha := cha * (width + 1)
	//	//		result += mianji
	//	//		result -= mianjicha
	//	//		for len(stack_left) > 0 {
	//	//			result -= height[stack_left[len(stack_left)-1]]
	//	//			stack_left = stack_right[:len(stack_right)-1] //出栈
	//	//		}
	//	//		for len(stack_right) > 1 {
	//	//			result -= height[stack_right[len(stack_right)-1]]
	//	//			stack_right = stack_right[:len(stack_right)-1] //出栈
	//	//		}
	//	//		result -= height[stack_right[len(stack_right)-1]]
	//	//	}
	//	//} else {
	//	//	if height[i] < height[len(stack_left)-1] {
	//	//		stack_left = append(stack_left, i)
	//	//	} else if height[i] >= height[stack_left[len(stack_left)-1]] && len(stack_right) < 1 {
	//	//		stack_right = append(stack_right, i)
	//	//	} else if height[i] >= height[stack_right[len(stack_right)-1]] {
	//	//		stack_right = append(stack_right, i)
	//	//	}
	//	//}
	//
	//}
}
