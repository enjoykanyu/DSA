package erfenfa

/*
给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。



示例 1：

输入：x = 4
输出：2
示例 2：

输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
*/
//思路 从1枚举到给定数值的一半 0 1 2 3  4    0  1 4 9 16 二分
func main() {

	//fmt.Print(mySqrt(8))
}

func MySqrt(x int) int {
	right := x / 2            //右边界 4
	left := 0                 //左边界 0
	mid := (right - left) / 2 //2
	for left <= right {
		if mid*mid <= x {
			left = left + 1
		} else {
			right = right - 1
		}
		mid = left + (right-left)/2 //2
	}
	return mid - 1 //至于这里为啥需减一呢
}
