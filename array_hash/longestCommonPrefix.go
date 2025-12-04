package main

// LeetCode: Longest Common Prefix
// Problem: Write a function to find the longest common prefix string
//          amongst an array of strings. If there is no common prefix, return "".
//
// ========================= NOTES=========================
// Python vs Go Implementation
// -----------------------------------------------------------------------------


//
// GO:
// - Same character-by-character scan logic.
// - Uses rune conversion for safe comparison.
// - Loops explicitly instead of Python's concise syntax.
//
// Similarities:
// - Same algorithm → vertical scanning.
// - Same complexity: O(n * m) (n = number of words, m = prefix length).
// - Both return substring up to the mismatch index.
//
// Differences:
// - Python slicing is simpler (`str[:i]`).
// - Go requires explicit loops and rune comparison.
// - Python gracefully handles empty lists; Go requires caller to avoid that.
//
// This Go solution is a direct translation of the standard Python vertical scan.
// =============================================================================

func longestCommonPrefix(strs []string) string {
    // Iterate through characters of the first string
    for i := 0; i < len(strs[0]); i++ {
        // Compare this character against all strings
        for _, word := range strs {
            // Check out-of-bounds OR mismatch
            if i == len(word) || rune(word[i]) != rune(strs[0][i]) {
                return word[:i]
            }
        }
    }
    // Entire first string is a prefix
    return strs[0]
}

// Optional test
func main() {
    println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
    println(longestCommonPrefix([]string{"dog", "racecar", "car"}))     // ""
}
