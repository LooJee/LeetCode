package PalindromeNumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := 0
	xCopy := x

	for x != 0 {
		num = num*10 + x%10
		x /= 10
	}

	return num == xCopy
}
