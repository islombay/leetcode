package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBinaryTreeToArray(head *TreeNode) []int {
	tmp := []int{}
	if head.Left != nil {
		tmp = append(tmp, convertBinaryTreeToArray(head.Left)...)
	}
	tmp = append(tmp, head.Val)

	if head.Right != nil {
		tmp = append(tmp, convertBinaryTreeToArray(head.Right)...)
	}
	return tmp
}

func main() {
	cases := []struct {
		id   int
		nums []int
	}{
		{
			id:   1,
			nums: []int{-10, -3, 0, 5, 9},
		},
		{
			id:   2,
			nums: []int{1, 3},
		},
		{
			id:   3,
			nums: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, c := range cases {
		res := sortedArrayToBST(c.nums)
		resArray := convertBinaryTreeToArray(res)

		if len(resArray) != len(c.nums) {
			fmt.Printf("Test %d failed, expected len %d, got len %d\n", c.id, len(c.nums), len(resArray))
			continue
		}
		for i := range c.nums {
			if c.nums[i] != resArray[i] {
				fmt.Printf("Test %d failed, expected %v, got %v\n", c.id, c.nums, resArray)
				break
			}
		}
	}
	fmt.Println("Finished")
}
func convert(nums []int, left, right int) *TreeNode {
	mid := left + (right-left)/2
	tmp := &TreeNode{Val: nums[mid]}
	if mid-1 >= left {
		tmp.Left = convert(nums, left, mid-1)
	}
	if mid+1 <= right && mid+1 < len(nums) {
		tmp.Right = convert(nums, mid+1, right)
	}
	return tmp
}
func sortedArrayToBST(nums []int) *TreeNode {
	return convert(nums, 0, len(nums))
}
