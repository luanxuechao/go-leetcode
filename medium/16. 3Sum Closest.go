package medium

import (
	"fmt"
	"sort"
)

// ThreeSumClose is
func ThreeSumClose(nums []int, target int) {
	sort.Ints(nums)
	n := len(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				fmt.Println(result)
			}
			if sum > target {
				right--
			}
			if sum < target {
				left++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}

	fmt.Println(result)
	// return result + target
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
