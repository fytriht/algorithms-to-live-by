package solution

func rotateString(A string, B string) bool {
	for i := 0; i < len(A)-1; i++ {
		A = A[1:] + string(A[0])
		if A == B {
			return true
		}
	}
	return A == B
}
