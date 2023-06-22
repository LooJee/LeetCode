package pondsizeslcci

import "sort"

func pondSizes(land [][]int) []int {
	var (
		result = []int{}
	)

	if len(land) == 0 || len(land[0]) == 0 {
		return result
	}

	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 0 {
				land[i][j] = -1
				result = append(result, bfs(land, [][2]int{{i, j}}))
			}
		}
	}

	sort.Ints(result)

	return result
}

func bfs(land [][]int, pos [][2]int) int {
	var size int
	for len(pos) > 0 {
		size++

		var (
			p0 = pos[0][0]
			p1 = pos[0][1]
		)

		for a := -1; a <= 1; a++ {
			for b := -1; b <= 1; b++ {
				i := p0 + a
				j := p1 + b

				if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) || land[i][j] != 0 {
					continue
				}

				pos = append(pos, [2]int{i, j})

				land[i][j] = -1
			}
		}

		pos = pos[1:]
	}

	return size
}
