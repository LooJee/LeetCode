package MoveZeroes

func moveZeroes(nums []int)  {
	length := len(nums)
	z := 0
	for i := 0; i < length-z; i++ {
		if nums[i] == 0 {
			for j := i; j < length - z - 1; j++ {
				nums[j] = nums[j+1]
			}
			nums[length-z-1] = 0
			z++
			i--
		}
	}
}

/*
上面的解法是我自己想的，虽然空间复杂度为O(1)，但是时间复杂度和操作数都不是很完美。所以，我看了下别人的思路，感觉大家都好强啊。
下面是按照推荐思路的解法
*/
func moveZeroesBetter(nums []int)  {
	lastZeroFound := 0

	for i, v := range nums {
		if v != 0 {
			tmp := v
			nums[i] = nums[lastZeroFound]
			nums[lastZeroFound] = tmp
			lastZeroFound++
		}
	}
}

