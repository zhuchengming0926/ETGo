/**
 * @File: frontTailPrint.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/26 11:42
 */

/*
从尾到头打印链表：递归、栈
*/

package link

import "fmt"

type ListNode struct {
	Val    int
	Next   *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var result []int
	reverseLink(head, &result)
	return result
}

func reverseLink(root *ListNode, result *[]int)  {
	if root.Next != nil {
		reverseLink(root.Next, result)
	}
	*result = append(*result, root.Val)
}

func main()  {
	var head = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Next = &ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	result := reversePrint(head)
	fmt.Println(result)
}