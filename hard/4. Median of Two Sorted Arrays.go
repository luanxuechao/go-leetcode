package hard

import (
	"fmt"
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	var median float64
	if len(nums1)%2 == 1 {
		median = float64(nums1[len(nums1)/2])
	} else {
		median = float64(nums1[len(nums1)/2] + nums1[len(nums1)/2-1])
		median = median / 2
	}
	return median
}
