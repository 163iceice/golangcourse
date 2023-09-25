func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    cur1 := l1
    cur2 := l2
    
    var root, cur *ListNode
    for cur1 != nil || cur2 != nil {
        if cur1 == nil || (cur2 != nil && cur1.Val >= cur2.Val) {
            if cur == nil {
                cur = cur2
                root = cur
            } else {
                cur.Next = cur2
                cur = cur.Next
            }
            cur2 = cur2.Next
        } else if cur2 == nil || (cur1 != nil && cur1.Val < cur2.Val) {
            if cur == nil {
                cur = cur1
                root = cur
            } else {
                cur.Next = cur1
                cur = cur.Next
            }
            cur1 = cur1.Next
        }
    }
    if cur != nil {
        cur.Next = nil   
    }
    return root
}