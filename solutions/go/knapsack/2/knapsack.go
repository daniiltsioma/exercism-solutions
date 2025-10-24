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
	// one-dimensional dp array.
	dp := make([]int, maximumWeight + 1)

	for _, item := range items {
		for w := maximumWeight; w >= item.Weight; w-- {
			dp[w] = max(dp[w], dp[w - item.Weight] + item.Value)
		}
	}

	return dp[maximumWeight]
}
