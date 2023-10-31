package internal

func ReturnError(err error) (string, string, error) {
	return "", "", err
}

// Permutations https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations[T any](data []T) [][]T {
	var result [][]T
	var helper func(int, []T)

	helper = func(n int, a []T) {
		if n == 1 {
			cp := make([]T, len(a))
			copy(cp, a)
			result = append(result, cp)
			return
		}

		for i := 0; i < n; i++ {
			helper(n-1, a)
			if n%2 == 0 {
				a[i], a[n-1] = a[n-1], a[i]
			} else {
				a[0], a[n-1] = a[n-1], a[0]
			}
		}
	}

	helper(len(data), data)
	return result
}
