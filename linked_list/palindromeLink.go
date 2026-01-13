package main

type ListNode struct {
    Val  int
    Next *ListNode
}

/*
Palindrome Linked List (Array-based approach)

Idea:
1. Traverse the linked list and store all node values in an array.
2. Use two pointers (left & right) to check if the array is a palindrome.

This works because a linked list palindrome reduces to
checking whether the sequence of values reads the same forward and backward.

Time Complexity: O(n)
Space Complexity: O(n)
*/
func isPalindrome(head *ListNode) bool {
    // Store linked list values
    arr := make([]int, 0)
    cur := head

    for cur != nil {
        arr = append(arr, cur.Val)
        cur = cur.Next
    }

    // Two-pointer palindrome check
    left, right := 0, len(arr)-1
    for left < right {
        if arr[left] != arr[right] {
            return false
        }
        left++
        right--
    }

    return true
}