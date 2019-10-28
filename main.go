package main

import (
	"fmt"
	"leetcode/easy"
	"leetcode/hard"
	"leetcode/medium"
)

func main() {
	var result [][]int = medium.FourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	fmt.Println(result)
	// var result int = medium.FourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	fmt.Println(result)
	var result1 int = easy.GetSum(60, 13)
	fmt.Println(result1)
	// var result2 string = easy.AddStrings("9", "99")
	var result4 float64 = hard.FindMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(result4)
	var result914 bool = easy.HasGroupsSizeX([]int{0, 0, 0, 1, 1, 1, 2, 2, 2})
	fmt.Println(result914)

}
