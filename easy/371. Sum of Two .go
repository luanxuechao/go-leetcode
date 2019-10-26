package easy

// GetSum  is  a function not use + -
func GetSum(a int, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return GetSum(sum, carry)
}
