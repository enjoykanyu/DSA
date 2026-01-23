package jiqiao

func singleNumber(nums []int) int {

	t := 0
	for i := 0; i < len(nums); i++ {
		t = t ^ nums[i]
	}
	return t
}
