package insertdeletegetrandomo1

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	nums      []int
	idxRecord map[int]int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixMilli())
	return RandomizedSet{
		nums:      []int{},
		idxRecord: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idxRecord[val]; ok {
		return false
	}

	this.nums = append(this.nums, val)
	this.idxRecord[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.idxRecord[val]
	if !ok {
		return false
	}

	// 将 val 的索引和最后一个元素交换
	this.nums[idx], this.nums[len(this.nums)-1] = this.nums[len(this.nums)-1], this.nums[idx]
	this.nums = this.nums[:len(this.nums)-1]

	// 如果 val 自身不是最后一个元素，则将交换后的元素的索引修改掉
	if idx != len(this.nums) {
		this.idxRecord[this.nums[idx]] = idx
	}
	delete(this.idxRecord, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.nums) == 1 {
		return this.nums[0]
	}

	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
