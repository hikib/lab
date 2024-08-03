package kata

import "sort"

func RemovNb(m uint64) [][2]uint64 {
	// find all a and b, where:
	// sum(1...m) - a - b == a * b

	var a, b uint64
	var result [][2]uint64

	// x = sum(1...m)
	x := sum(m)
	// --> x - a - b = a * b
	// --> x = a * b + a + b

	// b = (x*a+x-a)/a - x
	solveB := func() uint64 { return (x*a+x-a)/a - x }

	// x == a * b + a + b  ?
	productEqualsSum := func() bool { return a*b+a+b == x }

	for i := range make([]any, m) {
		a = uint64(i)
		if a == 0 {
			continue
		}
		b = solveB()

		if productEqualsSum() {
			result = append(result, [2]uint64{a, b})
			result = append(result, [2]uint64{b, a})
		}
	}

	// needs sorting to be accepted as result ...
	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	return result
}

func sum(m uint64) uint64 {
	var x uint64
	for i := range make([]any, m+1) {
		x += uint64(i)
	}
	return x
}
