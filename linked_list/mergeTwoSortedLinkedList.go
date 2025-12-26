package main

// LeetCode: Merge Two Sorted Lists
// Problem: Merge two sorted linked lists and return it as a new sorted list.
// The new list is made by splicing together the nodes of the first two lists.
//
// ========================= NOTES =========================
// - We use a dummy node (`res`) to simplify edge cases.
// - `pointer` always points to the last node of the merged list.
// - At each step, we attach the smaller node from list1 or list2.
// - After one list is exhausted, we append the remaining nodes from the other list.
//
// Time Complexity: O(n + m)
// Space Complexity: O(1) (in-place, no new nodes created)
// =========================================================

// Definition for singly-linked list.
type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    // Dummy node to simplify list construction
    res := &ListNode{}
    pointer := res

    // Merge while both lists have nodes
    for list1 != nil && list2 != nil {
        if list1.Val < list2.Val {
            pointer.Next = list1
            list1 = list1.Next
        } else {
            pointer.Next = list2
            list2 = list2.Next
        }
        // Move pointer forward
        pointer = pointer.Next
    }

    // Attach remaining nodes (only one of these will execute)
    if list1 != nil {
        pointer.Next = list1
    }
    if list2 != nil {
        pointer.Next = list2
    }

    // The merged list starts from res.Next
    return res.Next
}

// Optional test
func main() {
    // list1: 1 -> 2 -> 4
    list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
    // list2: 1 -> 3 -> 4
    list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

    merged := mergeTwoLists(list1, list2)
    for merged != nil {
        print(merged.Val, " ")
        merged = merged.Next
    }
    // Output: 1 1 2 3 4 4
}
