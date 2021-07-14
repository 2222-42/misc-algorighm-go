package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{
		Val:  0,
		Next: nil,
	}
	nextA := l1
	nextB := l2
	advance := 0
	current := &result

	for {
		// 終了条件
		if nextA == nil && nextB == nil {
			break
		}

		resultOfDigit := current.Val

		if nextA != nil {
			resultOfDigit = resultOfDigit + nextA.Val
		}

		if nextB != nil {
			resultOfDigit = resultOfDigit + nextB.Val
		}

		current.Val = resultOfDigit % 10
		advance = resultOfDigit / 10

		nextNode := ListNode{
			Val:  advance,
			Next: nil,
		}

		if nextA != nil {
			nextA = nextA.Next
		}

		if nextB != nil {
			nextB = nextB.Next
		}

		// 終了条件(余計なものが追加されてしまうので)
		if nextA == nil && nextB == nil {
			if advance > 0 {
				current.Next = &nextNode
			}
			break
		}
		current.Next = &nextNode
		// 次のステップのための準備
		current = current.Next
	}

	return &result
}
