package medium

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	nums := make([]rune, len(num1)+len(num2))
	for i := 0; i <= len(num2)-1; i++ {
		v1 := num2[i] - 0x30
		for j := 0; j <= len(num1)-1; j++ {

			v2 := num1[j] - 0x30
			r := v1 * v2
			nums[i+j+1] = nums[i+j+1] + rune(r%10)
			nums[i+j] = nums[i+j] + rune(r/10)
		}
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i-1] += nums[i] / 10
		nums[i] = nums[i] % 10
	}
	retInRune := make([]rune, 0)
	start := 0
	for nums[start] == 0 {
		start++
	}
	maxLen := len(num1) + len(num2) - 1
	for start <= maxLen {
		retInRune = append(retInRune, nums[start]+0x30)
		start++
	}
	return string(retInRune)
}
