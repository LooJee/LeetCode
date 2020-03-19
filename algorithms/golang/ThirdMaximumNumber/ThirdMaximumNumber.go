package ThirdMaximumNumber

/*这解法真是太傻了。。。*/
func thirdMax(nums []int) int {
	a := make([]int, 0, 3)
	for _, v := range nums {
		if len(a) == 0 {
			a = append(a, v)
		} else if len(a) > 0 && v > a[0] {
			v, a[0] = a[0], v
		}

		if v == a[0] {
			continue
		}

		if len(a) == 1 {
			a = append(a, v)
		} else if len(a) > 1 && v > a[1] {
			v, a[1] = a[1], v
		}

		if v == a[1] {
			continue
		}

		if len(a) == 2 {
			a = append(a, v)
		} else if len(a) > 2 && v > a[2] {
			v, a[2] = a[2], v
		}
	}

	if len(a) < 3 {
		return a[0]
	} else {
		return a[2]
	}
}
