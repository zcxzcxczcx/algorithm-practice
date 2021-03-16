package main

func main() {

}

type TreeNode struct {
	Word     string
	Children [26]*TreeNode
}

func findWords(board [][]byte, words []string) []string {
	root := &TreeNode{}
	for _, word := range words {
		node := root
		for _, v := range word {
			if node.Children[v-'a'] == nil {
				node.Children[v-'a'] = &TreeNode{}
			}
			node = node.Children[v-'a']
		}
		node.Word = word
	}
	var result []string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root, &result)
		}
	}
	return result
}
func dfs(i, j int, board [][]byte, root *TreeNode, result *[]string) {
	if i < 0 || j < 0 || i >= len(board) || (len(board) > 0 && j >= len(board[0])) {
		return
	}
	v := board[i][j]
	if v == '#' || root.Children[v-'a'] == nil {
		return
	}
	root = root.Children[v-'a']
	if root.Word != "" {
		*result = append(*result, root.Word)
		root.Word = ""
	}

	board[i][j] = '#'
	dfs(i-1, j, board, root, result)
	dfs(i+1, j, board, root, result)
	dfs(i, j+1, board, root, result)
	dfs(i, j-1, board, root, result)

	board[i][j] = v
}
