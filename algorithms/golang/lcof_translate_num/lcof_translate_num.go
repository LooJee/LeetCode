package lcof_translate_num

//这是一道变形的跳台阶问题，加了限制条件
func translateNum(num int) int {
	if num == 0 {
		return 1
	}

	a, b := 1, 1
	d := num / 10
	c := num % 10

	for d != 0 {
		tmp := a
		a = b
		t := (d%10)*10 + c
		if 10 <= t && t <= 25 {
			b = b + tmp
		}

		c = d % 10
		d = d / 10
	}

	return b
}
