package easy

func AddStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	i := len(num1) - 1
	j := len(num2) - 1
	carry := byte(0)
	res := ""
	for i >= 0 || j >= 0 {
		sum := byte(0)
		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}
		sum += carry
		carry = sum / 10
		n := sum % 10
		res = string(n+'0') + res
		i--
		j--
	}
	if carry > 0 {
		res = string(carry+'0') + res
	}
	return res
}
