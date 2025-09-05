package dandiaozhan
/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。



示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]

*/
func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{}

	stack = append(stack,0) //入栈未找到更高温度的下标 单调栈
	for i := 1; i <len(temperatures) ; i++ {
		for len(stack)>0 && temperatures[i]>temperatures[stack[len(stack)-1]] { //循环对比栈顶与当前温度 当前温度>栈顶的位置的温度 则栈顶位置对应的距离存入 同时栈顶出栈 再次循环对比栈顶和当前温度
			result[stack[len(stack)-1]] =  i-stack[len(stack)-1] //存入距离
			stack = stack[:len(stack)-1] //出栈
		}
		//当前温度未找到更高温度下标 入栈 即栈为空 对比栈顶小
		stack = append(stack,i)

	}
	return result
}