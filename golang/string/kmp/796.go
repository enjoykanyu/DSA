package kmp

/*
796. 旋转字符串
简单
相关标签
premium lock icon
相关企业
给定两个字符串, s 和 goal。如果在若干次旋转操作之后，s 能变成 goal ，那么返回 true 。

s 的 旋转操作 就是将 s 最左边的字符移动到最右边。

例如, 若 s = 'abcde'，在旋转一次之后结果就是'bcdea' 。

示例 1:

输入: s = "abcde", goal = "cdeab"
输出: true
示例 2:

输入: s = "abcde", goal = "abced"
输出: false

提示:

1 <= s.length, goal.length <= 100
s 和 goal 由小写英文字母组成
*/
//思路和28题差不多 旋转多次可以 看成最多旋转len(s)次 则可以拼接成原始长的字符串 这样的题目和28题就差不多了
func RotateString(s string, goal string) bool {
	new_str := s + s //abcdeabcd
	i, j := 0, 0     //j从头开始遍历 j需匹配的目标字符串
	len1 := len(goal)
	arr := make([]int, len1)
	next_arr := nextStr(goal, arr)
	for i < len(new_str) && j < len(goal) {
		if j == -1 || new_str[i] == goal[j] {
			i++
			j++
		} else {
			j = next_arr[j]
		}
		//遍历全部完成 j==到匹配目标字符串的长度
	}
	if j == len(s) {
		return true //当前遍历位置减去目标长度=起始点
	}

	return false
}

func nextStr(p string, next []int) []int {
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
