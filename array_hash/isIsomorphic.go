package main

// LeetCode: Isomorphic Strings
// Problem: Determine if two strings s and t are isomorphic.
// Two strings are isomorphic if characters in s can be replaced to get t,
// with a one-to-one (bijective) character mapping.
//
// ========================= NOTES =========================
// - We maintain two hash maps to enforce a bijection:
//     s -> t mapping
//     t -> s mapping
// - This prevents cases where two different characters in s map to
//   the same character in t (and vice versa).
// - We iterate once through the strings and validate consistency.
//
// Time Complexity: O(n)
// Space Complexity: O(1) (bounded by character set)
// =========================================================

func isIsomorphic(s string, t string) bool {
    hashmapS := make(map[rune]rune) // mapping from s -> t
    hashmapT := make(map[rune]rune) // mapping from t -> s

    for i := 0; i < len(s); i++ {
        cS := rune(s[i])
        cT := rune(t[i])

        // Check existing mapping from s to t
        if val, ok := hashmapS[cS]; ok && val != cT {
            return false
        }

        // Check existing mapping from t to s
        if val, ok := hashmapT[cT]; ok && val != cS {
            return false
        }

        // Record mappings
        hashmapS[cS] = cT
        hashmapT[cT] = cS
    }

    return true
}
