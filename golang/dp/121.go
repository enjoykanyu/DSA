package dp

func maxProfit(prices []int) int {
	zuo := prices[0] //当前遍历的最小
	value := 0
	for i := 0; i < len(prices); i++ {
		zuo = min(prices[i], zuo)
		value = max(prices[i]-zuo, value)
	}

	return value
}
