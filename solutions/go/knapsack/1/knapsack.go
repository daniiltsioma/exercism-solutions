package knapsack

// import "fmt"

type Item struct {
	Weight, Value int
}

// max returns largest of two ints.
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) int {
	// allocate a dynamic table and init rows.
	dp := make([][]int, len(items)+1)
	for i := range dp {
		dp[i] = make([]int, maximumWeight+1)
	}

	// first row remains all zeros.
	for i, item := range items {
		for j := range maximumWeight + 1 {
			if item.Weight > j {
				// can't fit the weight.
				dp[i+1][j] = dp[i][j]
			} else {
				// choose the better option:
				// skip or
				// take and decrease the remaining weight.
				dp[i+1][j] = max(
					dp[i][j],
					dp[i][j - item.Weight] + item.Value,
				)
			}
		}
	}

	return dp[len(items)][maximumWeight]
}
