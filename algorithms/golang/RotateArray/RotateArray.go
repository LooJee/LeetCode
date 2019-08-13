package RotateArray

func rotate(nums []int, k int)  {
	l := len(nums)-1
	for i := 0; i < k; i++ {
		t := nums[l]
		copy(nums[1:], nums[:l])
		nums[0] = t
	}
}
