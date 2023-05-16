package TwoSumII

/*
因为传入的数组已经排序，所以，除了使用在TwoSum中采用map来记录已经遍历过的值来实现O(n)的算法，
本例还可以使用i,j两个索引分别指向首尾，当两个数的和小于target时，向后移动i索引；当和大于target时，向前移动j索引，直到和与target相等
*/
func twoSum(numbers []int, target int) []int {
	var (
		left, right = 0, len(numbers) - 1
	)

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right --
		}
	}

	return []int{}
}
