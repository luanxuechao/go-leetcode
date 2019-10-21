package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	// var result [][]int = medium.FourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)
	var result []int = easy.TwoSumII([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}
