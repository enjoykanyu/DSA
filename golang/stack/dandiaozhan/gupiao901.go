package dandiaozhan

import "math"

/*
设计一个算法收集某些股票的每日报价，并返回该股票当日价格的 跨度 。

当日股票价格的 跨度 被定义为股票价格小于或等于今天价格的最大连续日数（从今天开始往回数，包括今天）。

例如，如果未来 7 天股票的价格是 [100,80,60,70,60,75,85]，那么股票跨度将是 [1,1,1,2,1,4,6] 。

实现 StockSpanner 类：

StockSpanner() 初始化类对象。
int next(int price) 给出今天的股价 price ，返回该股票当日价格的 跨度 。


示例：

输入：
["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]
输出：
[null, 1, 1, 1, 2, 1, 4, 6]

解释：
StockSpanner stockSpanner = new StockSpanner();
stockSpanner.next(100); // 返回 1
stockSpanner.next(80);  // 返回 1
stockSpanner.next(60);  // 返回 1
stockSpanner.next(70);  // 返回 2
stockSpanner.next(60);  // 返回 1
stockSpanner.next(75);  // 返回 4 ，因为截至今天的最后 4 个股价 (包括今天的股价 75) 都小于或等于今天的股价。
stockSpanner.next(85);  // 返回 6 单调栈 这里的算法思想为 当后边的比当前小入栈 后边的比栈顶大 则可以出栈计算出当前股价的天数（距离次栈顶的距离）  次栈顶的距离通过curDay循环递增给出
*/
type StockSpanner struct {
	stack [][2]int //结构体栈属性 [0]存储price [1]存储当前第几天
	curDay int //记录第几次调用next
}


func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{math.MaxInt32,-1}},-1}
}


func (this *StockSpanner) Next(price int) int {
	this.curDay ++ //第几次调用
	for price >= this.stack[len(this.stack)-1][0] { //当前价格大于栈顶 出栈
		//result := this.curDay - this.stack[len(this.stack)-1][1] //记录当前天数与栈顶的距离
		this.stack = this.stack[:len(this.stack)-1] //出栈
	}
	this.stack = append(this.stack,[2]int{price,this.curDay}) //当前价格小于栈顶 入栈 遍历到头了
	return this.curDay - this.stack[len(this.stack)-2][1]   //当前调用次数-次栈顶的天数=差值
}
