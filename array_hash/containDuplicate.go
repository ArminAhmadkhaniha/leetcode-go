package main

// LeetCode: Contains Duplicate
// Problem: Check if an integer slice contains any duplicate values.
//
// ========================= NOTES =========================
// In Python, the idiomatic approach uses a built-in `set`, because Python has a
// dedicated hash-based set type. Example:
//     def hasDuplicate(nums):
//         return len(nums) != len(set(nums))
//
// Go does NOT have a built-in `set` type.
// Instead, we simulate a set using a map with the form:
//     map[int]struct{}
// Why `struct{}`? Because it occupies **zero bytes**. Using `bool` would work
// too, but it wastes more memory.
//
// Similarities:
// - Python set  <-->  Go map[T]struct{}
// - Both give average O(1) lookup and insertion.
//
// Differences:
// - Python has a dedicated set type.
// - Go requires a map to act as a set.
// - Python's solution is shorter because `set()` is built-in.
//
// This Go solution mirrors the Python logic by building a set of values and
// checking if duplicates existed via size comparison.
// =============================================================================

// hasDuplicate checks whether the given slice contains duplicates.
func hasDuplicate(nums []int) bool {
    hashMap := make(map[int]struct{}) // set-like map

    for _, val := range nums {
        hashMap[val] = struct{}{}     // insert element into the "set"
    }

    // If duplicates existed, the map will have fewer elements than nums.
    return len(hashMap) < len(nums)
}

// Example main() for manual testing (not needed for LeetCode).
func main() {
    sample := []int{1, 2, 3, 1}
    println(hasDuplicate(sample)) // true
}

// other sol
// func hasDuplicate(nums []int) bool {
//     set := make(map[int]int)
//     for _, val := range nums{
//         _, ok := set[val]
//         if ok{
//             return true
//         } else{
//             set[val] = 1

//         }
 
//     }
//     return false

    
// }

