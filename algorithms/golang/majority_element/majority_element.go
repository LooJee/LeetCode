package majority_element

/*
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/majority-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func majorityElement(nums []int) int {
	var count = 0
	var curNum = 0
	for _, num := range nums {
		if curNum == num {
			count++
		} else if curNum != num && count > 0 {
			count--
		} else {
			curNum = num
		}
	}

	return curNum
}
