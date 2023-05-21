package openthelock

func openLock(deadends []string, target string) int {
	var (
		queue       = []string{"0000"}
		depth       = 0
		used        = map[string]struct{}{"0000": {}}
		deadendsMap = make(map[string]struct{})
	)

	for _, deadend := range deadends {
		deadendsMap[deadend] = struct{}{}
	}

	for len(queue) > 0 {
		queLen := len(queue)

		for i := 0; i < queLen; i++ {
			answer := queue[0]
			queue = queue[1:]

			//判断是否是正确答案
			if answer == target {
				return depth
			}

			//判断是否是死亡数字
			if _, ok := deadendsMap[answer]; ok {
				continue
			}

			for j := 0; j < 4; j++ {
				up := plusOne([]byte(answer), j)
				if _, ok := used[up]; !ok {
					used[up] = struct{}{}
					queue = append(queue, up)
				}

				below := minusOne([]byte(answer), j)
				if _, ok := used[below]; !ok {
					used[below] = struct{}{}
					queue = append(queue, below)
				}
			}
		}

		depth++
	}

	return -1
}

func plusOne(answer []byte, idx int) string {
	if answer[idx] == '9' {
		answer[idx] = '0'
	} else {
		answer[idx]++
	}

	return string(answer)
}

func minusOne(answer []byte, idx int) string {
	if answer[idx] == '0' {
		answer[idx] = '9'
	} else {
		answer[idx]--
	}

	return string(answer)
}
