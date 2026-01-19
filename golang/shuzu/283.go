package shuzu

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			mid := nums[left]
			nums[left] = nums[right]
			nums[right] = mid
			left++
		}
	}
}
