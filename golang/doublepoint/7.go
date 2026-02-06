package doublepoint

import "math"

func reverse(x int) int {
	value := 0
	for x != 0 {
		if value > math.MaxInt32/10 || value < math.MinInt32/10 { //范围限制大小比较<2
			return 0
		}
		jinwei := x % 10
		x = x / 10
		value = jinwei + value*10

	}
	return value
}
