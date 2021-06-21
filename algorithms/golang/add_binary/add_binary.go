package add_binary

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	if lb > la {
		la, lb = lb, la
		a, b = b, a
	}
	c := []byte(a)
	var temp byte = 0
	for i := la - 1; i >= 0; i-- {
		sum := a[i] - '0' + temp
		if j := lb + i - la; j >= 0 {
			sum += b[j] - '0'
		}
		temp = sum / 2
		c[i] = sum%2 + '0'
	}
	res := string(c)
	if temp == 1 {
		res = "1" + res
	}
	return res
}
