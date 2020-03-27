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
}

func numPairsDivisibleBy60(time []int) int {
	cnt := 0
	for i := 0; i < len(time); i++ {
		for j := i + 1; j < len(time); j++ {
			if (time[i]+time[j])%60 == 0 {
				cnt++
			}
		}
	}

	return cnt
}
