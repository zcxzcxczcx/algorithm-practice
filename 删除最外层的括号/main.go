package main

func main() {

}
func removeOuterParentheses(S string) string {
	if S == "" {
		return ""
	}
	var count int
	var start int
	var result string
	for k, v := range S {
		if v == '(' {
			count++
		}
		if v == ')' {
			count--
		}
		if count == 0 {
			result = result + S[start+1:k]
			start = k + 1
		}
	}
	return result

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var result []int
	if head == nil {
		return result
	}
	for head != nil {
		result = reversePrint(head.Next)
	}
	result = append(result, head.Val)

	return result

}
