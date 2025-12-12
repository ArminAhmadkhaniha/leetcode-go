package main

// majorityElement returns the element that appears more than n/2 times.
// We use the Boyer–Moore Voting Algorithm, which finds the majority
// in one linear pass and O(1) space.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func majorityElement(nums []int) int {
    res := 0      // candidate
    counter := 0  // vote balance

    for _, val := range nums {
        if val == res {
            // Same as current candidate → increase vote
            counter++
        } else if counter > 0 {
            // Different value → decrease vote
            counter--
        } else {
            // Vote balance dropped to zero → choose new candidate
            res = val
            counter = 1
        }
    }

    return res
}


// Hash
// Time Complexity: O(n)
// Space Complexity: O(n)

// func majorityElement(nums []int) int {
//     n := len(nums)
//     hashmap := make(map[int]int)
//     for _, val := range nums {
//         hashmap[val]++
//     }
//     for key, val := range hashmap {
//         if val > n/2 {
//             return key
//         }
//     }
//     return -1
    
    
// }
