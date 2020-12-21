/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "math/rand"

type Solution struct {
    head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution{head:head}
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    curr := this.head
    arr := []int{}
    
    for curr != nil{
        arr = append(arr, curr.Val)    
        curr = curr.Next
    }
    
    return arr[rand.Intn(len(arr))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
