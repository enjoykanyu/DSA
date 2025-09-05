package dandiaozhan
/*


nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。

给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。

对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。

返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。



示例 1：

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
示例 2：

输入：nums1 = [2,4], nums2 = [1,2,3,4].
输出：[3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 2 ，用加粗斜体标识，nums2 = [1,2,3,4]。下一个更大元素是 3 。
- 4 ，用加粗斜体标识，nums2 = [1,2,3,4]。不存在下一个更大元素，所以答案是 -1 。 3 2 1 16 6 15 39 33 101

*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	stack:= []int{nums2[0]}
	result := make([]int,10000) //存储数组2的是否有更大的数值
	result_1 := make([]int,len(nums1))
	//遍历数组2 看数组2比右边更大的元素值
	for i := 0; i <len(nums2) ; i++ {
		result[nums2[i]] = -1
		for len(stack)>0 &&nums2[i] >stack[len(stack)-1] {
			result[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1] //出栈
		}
		stack = append(stack,nums2[i]) //入栈

	}
	for i := 0; i < len(nums1); i++ {
		result_1[i] = result[nums1[i]]
	}
	return result_1
}