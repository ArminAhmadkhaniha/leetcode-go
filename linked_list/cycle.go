package main

// LeetCode: Linked List Cycle
// Problem: Determine if a linked list has a cycle.
//
// ========================= NOTES =========================
// - This solution uses Floyd’s Tortoise and Hare algorithm.
// - Two pointers traverse the list at different speeds:
//     slow moves 1 step at a time
//     fast moves 2 steps at a time
// - If a cycle exists, fast will eventually meet slow.
// - If fast reaches the end (nil), the list has no cycle.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// =========================================================

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
    fast := head
    slow := head

    // Traverse the list with two pointers
    for fast != nil && fast.Next != nil {
        slow = slow.Next          // move by 1
        fast = fast.Next.Next     // move by 2

        // If they meet, a cycle exists
        if slow == fast {
            return true
        }
    }

    // Fast pointer reached the end → no cycle
    return false
}

// Optional test
func main() {
    // Create a cycle: 1 -> 2 -> 3 -> 2 ...
    node1 := &ListNode{1, nil}
    node2 := &ListNode{2, nil}
    node3 := &ListNode{3, nil}

    node1.Next = node2
    node2.Next = node3
    node3.Next = node2 // cycle here

    println(hasCycle(node1)) // true
}
