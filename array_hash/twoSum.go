package main

// LeetCode: Two Sum
// Problem: Given an array of integers, return indices of the two numbers
//          such that they add up to a specific target.
//
// ========================= NOTES =========================
// Python vs Go Implementation
// -----------------------------------------------------------------------------
//
// Python has a built-in dictionary and uses the `in` operator for O(1) lookups.
// Everything is compact due to Python's high-level abstractions.
//
// GO:
// - Go does NOT have built-in dictionaries named dict, but has `map`, which
//   serves the same purpose (hash map with O(1) average operations).
// - Unlike Python, Go requires explicit declaration of map types.
// - Go does not support `diff in hashMap`, so we use:
//       idx, ok := hashMap[diff]
// - Go maps also return two values: the value and a boolean indicating existence.
//
// Similarities:
// - Both use a hash map of `value -> index`.
// - Both achieve O(n) time complexity.
// - Both return early when a valid pair is found.
//
// Differences:
// - Python map syntax is shorter and more dynamic.
// - Go requires explicit type declarations.
// - Go uses `ok` idiom for membership check.
//
// This Go solution is the exact equivalent of the Python logic.
// =============================================================================

// twoSum returns the indices of the two numbers that add up to target.
func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int) // value -> index

    for index, val := range nums {
        diff := target - val

        // Check if complement exists
        if idx, ok := hashMap[diff]; ok {
            return []int{idx, index}
        }

        // Store current number and its index
        hashMap[val] = index
    }

    // No solution found
    return nil
}

func main() {
    result := twoSum([]int{2, 7, 11, 15}, 9)
    if result != nil {
        println(result[0], result[1]) // Output: 0 1
    }
}
