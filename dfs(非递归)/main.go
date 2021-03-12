package main

import (
	"fmt"
	"time"
)

func main() {
	root := Node{
		Val: 1,
	}
	root.LeftNode = &Node{
		Val: 3,
	}

	root.LeftNode.LeftNode = &Node{
		Val: 5,
	}
	root.LeftNode.RightNode = &Node{
		Val: 4,
	}
	root.RightNode = &Node{
		Val: 2,
	}
	dfs(&root)
}

type Node struct {
	flag      int
	Val       int
	LeftNode  *Node
	RightNode *Node
}

// 非递归的写法
func dfs(root *Node) {
	if root == nil {
		return
	}
	var deqeue []*Node
	deqeue = append(deqeue, root)
	for len(deqeue) > 0 {
		fmt.Printf("deqeuedeqeuedeqeue1=%v\n", deqeue)
		l := len(deqeue)
		root = deqeue[l-1]
		if root.LeftNode != nil && root.LeftNode.flag != 1 {
			deqeue = append(deqeue, root.LeftNode)
			continue
		}

		deqeue = deqeue[0 : l-1]
		root.flag = 1
		fmt.Printf("deqeuedeqeuedeqeue2=%v\n", deqeue)
		fmt.Printf("dddddddddd=%v\n", root.Val)
		time.Sleep(1 * time.Second)
		if root.RightNode != nil {
			deqeue = append(deqeue, root.RightNode)
		}
	}

}
