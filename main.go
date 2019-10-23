package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	// var result [][]int = medium.FourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	// fmt.Println(result)
	var result int = medium.FourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
	fmt.Println(result)
}
