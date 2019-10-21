package easy

func TwoSumII(numbers []int, target int) []int {
	// var map1 map[int]string = make(map[int]string)
	// for idx, v := range numbers {
	// 	if map1[v] != "" {
	// 		idx1, _ := strconv.Atoi(map1[v])
	// 		return []int{idx1 + 1, idx + 1}
	// 	}
	// 	ret := target - v
	// 	fmt.Println(ret)
	// 	map1[ret] = strconv.Itoa(idx)
	// 	fmt.Println(map1[ret])
	// }
	i, j := 0, len(numbers)-1
	sum := numbers[i] + numbers[j]
	for sum != target {
		if sum < target {
			i++
		}
		if sum > target {
			j--
		}
		sum = numbers[i] + numbers[j]
	}
	return []int{i + 1, j + 1}
}
