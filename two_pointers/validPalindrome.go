package main

import "unicode"

// LeetCode: Valid Palindrome
// Problem: Given a string, determine whether it is a palindrome considering
//          only alphanumeric characters and ignoring case.
//
// ========================= NOTES =========================
// Python vs Go Implementation
// -----------------------------------------------------------------------------

//
// Explanation:
// - Two-pointer approach.
// - Skip non-alphanumeric characters.
// - Compare characters case‑insensitively.
// - O(n) time, O(1) space.
//
// GO :
// - Same pointer logic, but uses `unicode` package for case conversion
//   and character checks.
// - Go strings are bytes → converting to runes is required.
// - `unicode.IsLetter` and `unicode.IsDigit` replace Python's `isalnum()`.
// - Manual pointer increments inside loops.
//
// Similarities:
// - Two-pointer strategy identical in logic.
// - Same complexity: O(n) time, O(1) space.
// - Both skip non-alphanumeric characters and compare lowercase.
//
// Differences:
// - Python has built-in string methods (`isalnum`, `lower`).
// - Go requires helper function `isAlpha` and explicit rune conversion.
//
// This Go solution is the exact equivalent of the Python two-pointer approach.
// =============================================================================

// isPalindrome checks if a string is a valid palindrome ignoring
func isPalindrome(s string) bool {
    left := 0
    right := len(s) - 1

    for left < right {
        // Move left pointer until it points to alphanumeric
        for left < right && !isAlpha(rune(s[left])) {
            left++
        }
        // Move right pointer until it points to alphanumeric
        for right > left && !isAlpha(rune(s[right])) {
            right--
        }

        // Compare characters case-insensitively
        if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
            return false
        }

        left++
        right--
    }

    return true
}

// isAlpha checks whether a rune is alphanumeric
func isAlpha(ch rune) bool {
    return unicode.IsLetter(ch) || unicode.IsDigit(ch)
}

// Optional test
func main() {
    println(isPalindrome("A man, a plan, a canal: Panama")) // true
    println(isPalindrome("race a car"))                       // false
}

