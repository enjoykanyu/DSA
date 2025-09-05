package kmp

/*
28. 找出字符串中第一个匹配项的下标
简单
相关标签
premium lock icon
相关企业
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果 needle 不是 haystack 的一部分，则返回  -1 。

示例 1：

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
示例 2：

输入：haystack = "leetcode", needle = "leeto"
输出：-1
解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1 。

提示：

1 <= haystack.length, needle.length <= 104
haystack 和 needle 仅由小写英文字符组成
*/
//解法一：个人思路 ： 首先遍历目标字符串 依次看原先有无匹配 暴力双指针
func StrStr(haystack string, needle string) int {
	index := 0 // needle
	count := 0
	for i := 0; i < len(haystack); i++ { //haystack
		cur := i
		for index < len(needle) && i < len(haystack) && needle[index] == haystack[i] {
			i++
			index++
			count++
		}
		if count == len(needle) {
			return cur
		}
		i = cur
		index = 0
		count = 0
	}

	return -1
}

// 解法二：参考答案思路基础上将解法一进行剪纸
func StrStr2(haystack string, needle string) int {
	index := 0 // needle
	count := 0
	for i := 0; i <= len(haystack)-len(needle); i++ { //haystack 剪枝节点在这里 haystack遍历次数做多=len(haystack)=-len(needle) 注意这里为<=
		cur := i
		for index < len(needle) && i < len(haystack) && needle[index] == haystack[i] {
			i++
			index++
			count++
		}
		if count == len(needle) {
			return cur
		}
		i = cur
		index = 0
		count = 0
	}

	return -1
	return 1
}

// 解法二：参考思路 ： KMP
// 不匹配发生时，前面匹配的那一小段 abbaab 于我俩是相同的」，甲想，「这样的话，用 abbaab 的头部去匹配 abbaab 的尾部，最长的那段就是答案
// 重新匹配的时候对比已匹配的头部和尾部重合点 找到重合点 则无需重新回到起始点 只需回答哦重合点的下个位置进行匹配
// abbaab 的头部有 a, ab, abb, abba, abbaa（不包含最后一个字符。下文称之为「真前缀」）
// abbaab 的尾部有 b, ab, aab, baab, bbaab（不包含第一个字符。下文称之为「真后缀」）
// 这样最长匹配是 ab。也就是说甲不回退时，乙需要回退到第三个字符去和甲继续匹配。
// 先解释一下字符串的前缀和后缀。如果字符串A和B，存在A=BS，其中S是任意的非空字符串，那就称B为A的前缀。例如，”Harry”的前缀包括{”H”, ”Ha”, ”Har”, ”Harr”}，我们把所有前缀组成的集合，称为字符串的前缀集合。同样可以定义后缀A=SB， 其中S是任意的非空字符串，那就称B为A的后缀，例如，”Potter”的后缀包括{”otter”, ”tter”, ”ter”, ”er”, ”r”}，然后把所有后缀组成的集合，称为字符串的后缀集合。要注意的是，字符串本身并不是自己的后缀
// kmp算法实际只需解决找到匹配字符串的前缀与后缀的公共部分的最大长度
func StrStr3(haystack string, needle string) int {
	i, j := 0, 0 //j从头开始遍历 j需匹配的目标字符串
	len1 := len(needle)
	arr := make([]int, len1)
	next_arr := next(needle, arr)
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next_arr[j]
		}
		//遍历全部完成 j==到匹配目标字符串的长度
	}
	if j == len(needle) {
		return i - j //当前遍历位置减去目标长度=起始点
	}

	return -1
}

// pmt[] 数组为包括当前位置最大的长度 next往后移动一位 表示计算前个位置最大的长度
func next(p string, next []int) []int {
	next[0] = -1
	i, j := 0, -1 //j从头开始遍历
	for i < len(p) {
		if j == -1 || p[i] == p[j] { //此时匹配了 而且从0到j的位置为 i索引位置的最大公共长度 j==-1匹配
			i++
			j++
			if i >= len(p) {
				break
			}
			next[i] = j
		} else { //不匹配 j索引回退到最大长度
			j = next[j]

		}
	}
	return next
}
