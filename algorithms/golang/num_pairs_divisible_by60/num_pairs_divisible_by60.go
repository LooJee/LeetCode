package main

import "fmt"

/*
In a list of songs, the i-th song has a duration of time[i] seconds.
Return the number of pairs of songs for which their total duration in seconds is divisible by 60.
Formally, we want the number of indices i < j with (time[i] + time[j]) % 60 == 0.

Example 1:
Input: [30,20,150,100,40]
Output: 3
Explanation: Three pairs have a total duration divisible by 60:
(time[0] = 30, time[2] = 150): total duration 180
(time[1] = 20, time[3] = 100): total duration 120
(time[1] = 20, time[4] = 40): total duration 60

Example 2:
Input: [60,60,60]
Output: 3
Explanation: All three pairs have a total duration of 120, which is divisible by 60.

Note:
1 <= time.length <= 60000
1 <= time[i] <= 500
*/

func main() {
	fmt.Println(numPairsDivisibleBy60([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60([]int{60, 60, 60}))

	fmt.Println(numPairsDivisibleBy60V2([]int{30, 20, 150, 100, 40}))
	fmt.Println(numPairsDivisibleBy60V2([]int{60, 60, 60}))
}

func numPairsDivisibleBy60(time []int) int {
	dict := make(map[int][]int, len(time))
	choices := make([]int, 0)
	cnt := 0

	for i, v := range time {
		if arr, ok := dict[v]; !ok {
			dict[v] = []int{i}
		} else {
			arr = append(arr, i)
			dict[v] = arr
		}
	}

	i := 1
	choice := i * 60
	for choice < 1000 {
		choices = append(choices, choice)
		i++
		choice = i * 60
	}

	//遍历所有 time
	for idx, v := range time {
		//从所有可选的60的倍数中遍历，求与v的差
		for _, ch := range choices {
			part := ch - v
			if part < 1 {
				continue
			}

			if vList, ok := dict[part]; ok {
				for _, timeIdx := range vList {
					if timeIdx > idx {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

/*
两个数的余数相加为60，则两个数的和必为60的倍数。
则可以创建一个长度为60的数组，用于记录余数为0到59的数的个数。
因为time数组中数据的先后顺序，后面的数据查到的必定是在它前面的数据，满足题目中要求i<j的要求
*/
func numPairsDivisibleBy60V2(time []int) int {
	dict := make([]int, 60)
	cnt := 0

	for _, v := range time {
		mod := v % 60
		subMod := (60 - mod) % 60
		cnt += dict[subMod]
		dict[mod]++
	}

	return cnt
}
