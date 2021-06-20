package valid_parentheses

func isValid(s string) bool {
	stack := make([]byte, 0)

	parentheses := map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if (s[i] == '}' || s[i] == ']' || s[i] == ')') && len(stack) == 0 {
			return false
		} else if stack[len(stack)-1] != parentheses[s[i]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
