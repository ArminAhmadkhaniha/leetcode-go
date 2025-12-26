package main

// LeetCode: Reverse Linked List
// Problem: Reverse a singly linked list.
//
// ========================= NOTES =========================
// - We reverse the list iteratively using three pointers:
//     prev : previously reversed part of the list
//     cur  : current node being processed
//     temp : temporary storage for the next node
// - At each step, we reverse the direction of the `Next` pointer.
// - This is done in-place with no extra memory allocation.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// =========================================================

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head

    for cur != nil {
        // Save next node
        temp := cur.Next

        // Reverse pointer
        cur.Next = prev

        // Move prev and cur forward
        prev = cur
        cur = temp
    }

    // prev becomes the new head
    return prev
}

// Optional test
func main() {
    // Build list: 1 -> 2 -> 3 -> nil
    head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

    reversed := reverseList(head)

    for reversed != nil {
        print(reversed.Val, " ")
        reversed = reversed.Next
    }
    // Output: 3 2 1
}
