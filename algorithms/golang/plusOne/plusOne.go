package plusOne

func plusOne(digits []int) []int {
	lenth := len(digits)
	carry := 1

	for i := lenth - 1; i >= 0; i-- {
		j := digits[i] + carry
		if j > 9 {
			//有进位则继续和高位相加
			digits[i] = j % 10
		} else {
			//数据没有进位，则退出循环
			digits[i] = j
			carry = 0
			break
		}
	}

	//如果遍历完，进位还是1，则将1放在最高为
	if carry == 1 {
		digits = append([]int{carry}, digits...)
	}

	return digits
}
