package rangesumqueryimmutable

type NumArray struct {
	pre []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{pre: make([]int, len(nums)+1)}

	for i, num := range nums {
		arr.pre[i+1] = arr.pre[i] + num
	}

	return arr
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.pre[right+1]-this.pre[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
