package jiqiao

func sortColors(nums []int) {

	l := 0
	for i := 0; i <= 2; i++ {
		for j := l; j < len(nums); j++ {
			if nums[j] == i {
				nums[l], nums[j] = nums[j], nums[l]
				l++
			}
		}
	}
}
