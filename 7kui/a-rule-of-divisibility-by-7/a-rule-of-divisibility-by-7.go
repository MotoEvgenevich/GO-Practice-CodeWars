package kata

func Seven(n int64) []int {
	steps := 0
	for n >= 100 {
		lastDigit := n % 10
		n = n / 10
		n = n - 2*lastDigit
		steps++
	}

	return []int{int(n), steps}
}
