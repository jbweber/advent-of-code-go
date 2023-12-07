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

type Range struct {
	Start int
	End   int
}

func FindIntersection(r1, r2 Range) (Range, bool) {
	if r1.End < r2.Start || r2.End < r1.Start {
		return Range{}, false
	}

	iStart := max(r1.Start, r2.Start)
	iEnd := min(r1.End, r2.End)

	return Range{Start: iStart, End: iEnd}, true
}

func FindPartialIntersection(r1, r2 Range) (Range, []Range, bool) {
	// no intersection at all
	intersection, ok := FindIntersection(r1, r2)
	if !ok {
		return Range{}, nil, false
	}

	// no remainders
	if r1.Start >= r2.Start && r1.End <= r2.End {
		return intersection, nil, true
	}

	var remainder []Range

	// left
	if intersection.Start > r1.Start {
		start := r1.Start
		end := intersection.Start - 1
		remainder = append(remainder, Range{start, end})
	}

	// right
	if intersection.End < r1.End {
		start := intersection.End + 1
		end := r1.End
		remainder = append(remainder, Range{start, end})
	}

	return intersection, remainder, true
}
