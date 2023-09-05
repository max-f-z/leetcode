package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type NodeR struct {
	Val    int
	Next   *NodeR
	Random *NodeR
}

//lint:ignore U1000 unused
func copyRandomList(head *NodeR) *NodeR {
	if head == nil {
		return nil
	}

	ans := &NodeR{
		Val: head.Val,
	}

	cur := head
	curAns := ans

	ref := map[*NodeR]*NodeR{}
	ref[head] = ans

	for cur.Next != nil {
		tmp := &NodeR{
			Val: cur.Next.Val,
		}

		ref[cur.Next] = tmp

		cur = cur.Next
		curAns.Next = tmp
		curAns = curAns.Next
	}

	cur = head
	curAns = ans

	for cur.Next != nil {
		curAns.Random = ref[cur.Random]

		cur = cur.Next
		curAns = curAns.Next
	}

	curAns.Random = ref[cur.Random]

	return ans
}
