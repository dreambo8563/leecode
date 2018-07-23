package main

import (
	"fmt"
)

// 看来不得不循环了
// 思路:
//1. slice 转成  map[int][]int  value 对应 []index
// 循环 map  value <= target的时候 && diff 在map里 []index的lenght> 1 找到
func main() {
	target := -8
	origin := []int{-1, -2, -3, -4, -5}

	// show out for debug

	fmt.Println(twoSum(origin, target))
}

func twoSum(nums []int, target int) []int {
	mapOri := transform(nums)
	result := []int{}
	for key, value := range mapOri {
		diff := target - key
		if (diff == key) && len(value) > 1 {
			result = append(result, value...)
			break
		}
		if len(mapOri[diff]) > 0 && (mapOri[diff][0] != value[0]) {
			fmt.Println(value)
			fmt.Println(mapOri[diff])
			result = append(result, value...)
			result = append(result, mapOri[diff][0])
			break
		}
	}
	return result
}

func transform(input []int) map[int][]int {
	out := make(map[int][]int)
	for k, v := range input {
		out[v] = append(out[v], k)
	}

	return out
}
