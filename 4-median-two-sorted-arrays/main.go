package main

import "fmt"

func main() {
	// nums1 := []int{3}
	// nums2 := []int{-2, -1}

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	med := findMedianSortedArrays(nums1, nums2)

	// test := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(med)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0)

	n1Index := 0
	n2Index := 0

	if len(nums1) == 0 {
		merged = nums2
		goto Median
	} else if len(nums2) == 0 {
		merged = nums1
		goto Median
	}

	for {
		if n1Index == len(nums1) {
			merged = append(merged, nums2[n2Index:]...)
			goto Median
		} else if n2Index == len(nums2) {
			merged = append(merged, nums1[n1Index:]...)
			goto Median
		}
		if nums1[n1Index] > nums2[n2Index] {
			merged = append(merged, nums2[n2Index])
			n2Index++
		} else {
			merged = append(merged, nums1[n1Index])
			n1Index++
		}
	}

Median:
	fmt.Println(merged)
	fmt.Println(1 % 2)

	if len(merged) == 1 {
		return float64(merged[0])
	}

	if len(merged)%2 == 1 {
		return float64(merged[len(merged)/2])
	}
	med := float64(merged[len(merged)/2-1] + merged[len(merged)/2])
	return med / 2
}
