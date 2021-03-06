package solution

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		n = getNext(n)
		if m[n] {
			return false
		}
		if n == 1 {
			return true
		}
		m[n] = true
	}
}

func getNext(n int) int {
	s := 0
	for n != 0 {
		s += (n % 10) * (n % 10)
		n /= 10
	}
	return s
}
