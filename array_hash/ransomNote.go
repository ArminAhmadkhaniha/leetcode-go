package main

// canConstruct determines whether the string ransomNote
// can be built from the characters available in magazine.
//
// We count occurrences of each character in magazine
// and then check if ransomNote can consume those counts.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) average (map over limited character set)
func canConstruct(ransomNote string, magazine string) bool {
    freq := make(map[rune]int)

    // Count character frequencies from magazine
    for _, c := range magazine {
        freq[c]++
    }

    // Try to use characters for ransomNote
    for _, c := range ransomNote {
        if freq[c] == 0 { // Character not available or used up
            return false
        }
        freq[c]--
    }

    return true
}
