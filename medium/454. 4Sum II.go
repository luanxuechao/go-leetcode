package medium

func FourSumCount(A []int, B []int, C []int, D []int) int {
	ret := 0
	n := len(A)
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for i := 0; i <= n-1; i++ {
		for idx := range B {
			sum1 := A[i] + B[idx]
			sum2 := C[i] + D[idx]
			map1[sum1] = map1[sum1] + 1
			map2[sum2] = map2[sum2] + 1
		}
	}
	for k := range map1 {
		if map2[-k] != 0 {
			ret = ret + map1[k]*map2[-k]
		}
	}
	return ret
}
