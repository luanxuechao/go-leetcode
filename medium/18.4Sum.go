package medium

import (
	"fmt"
	"sort"
)

// FourSum  is  a
func FourSum(nums []int, target int) [][]int {
	results := [][]int{}
	sort.Ints(nums)
	fmt.Println(nums, target)
	var n int = len(nums)
	if len(nums) < 4 {
		return results
	}
	for i := 0; i < n-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j <= n-3; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			tartget1 := target - nums[i]
			left := j + 1
			right := n - 1
			for left < right {
				sum := nums[j] + nums[left] + nums[right]
				if sum == tartget1 {
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				}
				if sum > tartget1 {
					right--
				}
				if sum < tartget1 {
					left++
				}
			}
		}
	}
	return results
}

//[[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],
// [-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
