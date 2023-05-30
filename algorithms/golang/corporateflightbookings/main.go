package corporateflightbookings

func corpFlightBookings(bookings [][]int, n int) []int {
	var (
		nums = make([]int, n+1)
	)

	for _, booking := range bookings {
		nums[booking[0]] += booking[2]

		if booking[1]+1 < len(nums) {
			nums[booking[1]+1] -= booking[2]
		}
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
	}

	return nums[1:]
}
