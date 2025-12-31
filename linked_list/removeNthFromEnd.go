package main

type ListNode struct {
    Val  int
    Next *ListNode
}

/*
Remove Nth Node From End of List (Reverse-based approach)

Idea:
1. Reverse the linked list.
2. Remove the Nth node from the start of the reversed list.
3. Reverse the list again to restore original order.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // Reverse the linked list
    rev := reverse(head)

    // If we need to remove the first node in the reversed list
    if n == 1 {
        rev = rev.Next
    } else {
        // Move to the (n-1)th node in the reversed list
        cur := rev
        for i := 0; i < n-2; i++ {
            cur = cur.Next
        }
        // Remove the nth node
        cur.Next = cur.Next.Next
    }

    // Reverse again to restore original order
    return reverse(rev)
}

/*
Reverse a singly-linked list iteratively.

Returns the new head of the reversed list.
*/
func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head

    for cur != nil {
        temp := cur.Next
        cur.Next = prev
        prev = cur
        cur = temp
    }

    return prev
}
