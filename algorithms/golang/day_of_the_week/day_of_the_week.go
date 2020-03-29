package main

import "fmt"

/*
Given a date, return the corresponding day of the week for that date.
The input is given as three integers representing the day, month and year respectively.
Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.

Example 1:
Input: day = 31, month = 8, year = 2019
Output: "Saturday"

Example 2:
Input: day = 18, month = 7, year = 1999
Output: "Sunday"

Example 3:
Input: day = 15, month = 8, year = 1993
Output: "Sunday"

Constraints:
The given dates are valid dates between the years 1971 and 2100.

1971 1 1 is friday
*/

func main() {
	fmt.Println(dayOfTheWeek(31, 8, 2019))
	fmt.Println(dayOfTheWeek(18, 7, 1999))
	fmt.Println(dayOfTheWeek(15, 8, 1993))
	fmt.Println(dayOfTheWeek(1, 1, 1971))
	fmt.Println(dayOfTheWeek(29, 2, 2016))
	fmt.Println(dayOfTheWeek(29, 2, 1972))
	fmt.Println(dayOfTheWeek(29, 2, 1976))
	fmt.Println(dayOfTheWeek(21, 12, 1980))
}

func dayOfTheWeek(day int, month int, year int) string {
	monthDay := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	weakDay := []string{"Thursday", "Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Wednesday"}
	leapYears := 0
	monthDays := 0

	if year-1972 > 0 {
		leapYears = (year - 1972) / 4
		if (year-1972)%4 != 0 {
			leapYears++
		}
	}

	for i := 0; i < month; i++ {
		monthDays += monthDay[i]
		if (year-1972)%4 == 0 && i == 2 {
			monthDays++
		}
	}

	return weakDay[((year-1971)*365+leapYears+monthDays+day)%7]
}
