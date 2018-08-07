// 88. Merge Sorted Array
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

// Note:

// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
// Example:

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3

// Output: [1,2,2,3,5,6]
package main

import "fmt"

func main() {
	num1 := []int{1}
	num2 := []int{}
	m := 1
	n := 0
	merge(num1[:], m, num2[:], n)
	for _, v := range num1[:m+n] {

		fmt.Println(v)
	}
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		fmt.Println("m is 0")
		copy(nums1[m:], nums2[0:n])
		return
	}
	if n == 0 {
		return
	}
	// if n != 0 && m != 0 {

	// 	return
	// }
	fmt.Printf("1 - %d, 2 - %d \n", nums1[0], nums2[0])
	if nums1[0] < nums2[0] {
		fmt.Printf("nums1[0] %d, < nums2[0] %d \n", nums1[0], nums2[0])
		fmt.Printf("nums1[0] < nums2[0] m - %d, n - %d \n", m, n)
		merge(nums1[1:], m-1, nums2, n)
	} else {
		fmt.Println("else")
		fmt.Printf("else m - %d, n - %d \n", m, n)

		copy(nums1[1:m+1], nums1[0:m])

		nums1[0] = nums2[0]
		fmt.Println("before merge")
		for i, v := range nums1[:] {
			fmt.Println(v)
			fmt.Println(i)
		}
		if m == 1 && n == 1 {
			// copy(nums1[1:], nums1[0:1])
			fmt.Println(nums1[:])
			fmt.Println(nums2[:])
			// merge(nums1[0:1], 1, nums2[1:], n-1)
			return
		}
		merge(nums1[1:], m, nums2[1:], n-1)
	}

	fmt.Printf("other m - %d, n - %d \n", m, n)

}
