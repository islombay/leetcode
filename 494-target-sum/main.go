package main

func findTargetSumWays(nums []int, target int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if (target+totalSum)%2 != 0 || target > totalSum || target < -totalSum {
		return 0
	}

	subsetSum := (target + totalSum) / 2
	dp := make([]int, subsetSum+1)
	dp[0] = 1

	for _, n := range nums {
		for j := subsetSum; j >= n; j-- {
			dp[j] += dp[j-n]
		}
	}
	return dp[subsetSum]
}
