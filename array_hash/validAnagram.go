package main

// LeetCode: Valid Anagram
// Problem: Determine if two strings are anagrams of each other.
//
// ========================= NOTES =========================
// Python vs Go Implementation
// -----------------------------------------------------------------------------
//
// GO (this file):
// - Go requires manual creation of hash maps using `make(map[rune]int)`.
// - Go strings are byte sequences, but `range` gives Unicode runes.
// - Counting must be done explicitly.
// - Like Python's dict, Go maps offer O(1) average time operations.
//
// Similarities:
// - Both use hash maps (dict in Python, map in Go) to count character frequency.
// - Both run in O(n) time and O(1) alphabet space for ASCII-like inputs.
//
// Differences:
// - Python has built-in `Counter` which simplifies the logic.
// - Go explicitly converts bytes to runes when needed.
// - Python strings are Unicode by default; Go needs careful handling of runes.
//
// This Go solution mirrors the Python manual dictionary-counting logic.
// =============================================================================

// isAnagram checks whether two strings contain the same characters
// with the same frequencies.
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    hashMapS := make(map[rune]int)
    hashMapT := make(map[rune]int)

    // Count characters in both strings
    for i := 0; i < len(s); i++ {
        hashMapS[rune(s[i])]++
        hashMapT[rune(t[i])]++
    }

    // Compare frequency tables
    for _, val := range s {
        if hashMapS[val] != hashMapT[val] {
            return false
        }
    }

    return true
}

func main() {
    println(isAnagram("listen", "silent")) // true
    println(isAnagram("rat", "car"))       // false
}
