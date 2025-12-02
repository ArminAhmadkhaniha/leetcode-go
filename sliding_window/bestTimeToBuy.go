package main

// LeetCode: Best Time to Buy and Sell Stock
// Problem: Given an array where each element represents the stock price on a given day,
//          find the maximum profit that can be achieved from exactly one buy and one sell.
//          (Buy must come before sell.)
// GO:

// - Explicit increment of pointers (`sell++`).
// - Uses a helper `max` function because Go lacks built-in `max` for integers.
// - Very similar logic to Python but with explicit types and manual control flow.
//
// Similarities:
// - Same two-pointer strategy.
// - Same time and space complexity.
// - Same greedy update rule.
//
// Differences:
// - Python has built-in `max()`, while Go needs a custom one.
// - Python loops over indices with range; Go explicitly controls pointers.
//
// This Go solution is the direct analogue of the Python two-pointer method.
// =============================================================================

// helper max function since Go doesn't provide a built-in integer max
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// maxProfit returns the best possible profit from one buy and one sell.
func maxProfit(prices []int) int {
    buy := 0
    sell := 1
    maxprof := 0

    for sell < len(prices) {
		profit := prices[sell] - prices[buy]
        if prices[sell] > prices[buy] {
            maxprof = max(maxprof, profit)
        } else {
            // Found a new lower price; update buy pointer
            buy = sell
        }
        sell++
    }

    return maxprof
}

// Optional example run
func main() {
    println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
    println(maxProfit([]int{7, 6, 4, 3, 1}))     // 0
}
