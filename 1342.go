package solution

func numberOfSteps(num int) int {
	count := 0
	for {
		if num%2 != 0 {
			count++
			num -= 1
			if num == 0 {
				return count
			}
		} else {
			count++
			num /= 2
			if num == 0 {
				return count
			}
		}
	}
}
