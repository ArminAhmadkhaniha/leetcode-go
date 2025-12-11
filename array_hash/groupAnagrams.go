package main

// groupAnagrams groups words that are anagrams.
// Each word is converted to a fixed-size 26-length frequency key.
// Words with identical frequency signatures belong to the same group.
//
// Time Complexity: O(n * k)   where k = avg word length
// Space Complexity: O(n * k)  storing grouped words
func groupAnagrams(strs []string) [][]string {
    hashmap := make(map[[26]int][]string)

    // Build frequency signature for each string
    for _, val := range strs {
        count := [26]int{}
        for _, c := range val {
            count[c-'a']++
        }
        hashmap[count] = append(hashmap[count], val)
    }

    // Collect grouped anagrams from the map
    var result [][]string
    for _, group := range hashmap {
        result = append(result, group)
    }

    return result
}
