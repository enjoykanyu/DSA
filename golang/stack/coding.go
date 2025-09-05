/* 给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。



示例 1：

输入：s = "3[a]2[bc]"
输出："aaabcbc"
示例 2：

输入：s = "3[a2[c]]"
输出："accaccacc"
示例 3：

输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"
示例 4：

输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
*/

// 思路 未遇到]全部入栈 遇到 ] 出栈 记录单位字符组 当遇到[ 记录数字 当栈内为空 则字符组*数字为当前的字符串 当栈内不为空 当前的记录 继续出栈
package stack

import (
	"strconv"
	"strings"
)

// DecodeString 解码字符串
func DecodeString(s string) string {
	numStack := []int{} //"33[a2[c]]"
	strStack := []string{}
	num := 0
	result := ""
	for _, c := range s {
		if c >= '0' && c <= '9' { //数字叉乘
			n, _ := strconv.Atoi(string(c))

			//fmt.Println(n)
			num = num*10 + n
		} else if c == '[' { //遇到左括号 入栈
			numStack = append(numStack, num)    //数字入栈
			strStack = append(strStack, result) //字符串入栈
			num = 0                             //数字清零
			result = ""                         //栈内的字符串组合清空

		} else if c == ']' { //遇到右括号
			// 字符串组合
			count := numStack[len(numStack)-1]                   //获取数字
			numStack = numStack[:len(numStack)-1]                //数字出栈
			str := strStack[len(strStack)-1]                     //获取字符串
			strStack = strStack[:len(strStack)-1]                //出栈
			result = string(str) + strings.Repeat(result, count) //组合
		} else {
			result = result + string(c) //字符直接拼接
		}
	}
	return result
}
