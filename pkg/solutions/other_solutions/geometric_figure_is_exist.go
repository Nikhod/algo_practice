package other_solutions

// проверяет существует ли такой многоугольник, return true, otherwise - false
func IsExist(n int, sides []int) bool {
	var total int
	for i := 0; i < n; i++ {
		total += sides[i]
	}

	for i := 0; i < n; i++ {
		if sides[i] >= (total - sides[i]) {
			return false
		}
	}

	return true
}
