package main

import "fmt"

type C struct {
	Name         string
	Nums1, Nums2 []int
	M, N         int
	Result       []int
}

func New(nums1, num2 []int, m, n int, name string, result []int) C {
	n1 := make([]int, n+m)
	for i := range nums1 {
		n1[i] = nums1[i]
	}
	return C{
		Name:   name,
		Nums1:  n1,
		Nums2:  num2,
		M:      m,
		N:      n,
		Result: result,
	}
}

func main() {
	cases := []C{
		New([]int{1, 2, 3}, []int{2, 5, 6}, 3, 3, "1", []int{1, 2, 2, 3, 5, 6}),
		New([]int{1}, []int{}, 1, 0, "2", []int{1}),
		New([]int{}, []int{1}, 0, 1, "3", []int{1}),
	}

	var (
		isBad = false
	)

	for _, c := range cases {
		isBad = false

		merge(c.Nums1, c.M, c.Nums2, c.N)

		for i, e := range c.Nums1 {
			if e != c.Result[i] {
				fmt.Printf("Test failed %s expected %v, got %v\n", c.Name, c.Result, c.Nums1)
				isBad = true
				break
			}
		}
		if !isBad {
			fmt.Printf("Test successful %s\n", c.Name)
		}
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for i := range nums2 {
			nums1[i] = nums2[i]
		}
		return
	}
	if n == 0 {
		return
	}

	for _, e := range nums2 {
		nums1[m] = e
		m++
	}
	fmt.Println(nums1)

	bubbleSort(nums1)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap elements
			}
		}
	}
}
