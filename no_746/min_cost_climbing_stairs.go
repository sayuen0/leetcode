package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[n], dp[n-1] = 0, cost[len(cost)-1]

	for i := n - 1; i >= 0; i++ {
		dp[i] = cost[i] + min(dp[i+1], dp[i+2])
	}

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
