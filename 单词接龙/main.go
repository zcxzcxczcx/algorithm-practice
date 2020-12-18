package main

import "math"

func main() {

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	var graph [][]int
	var addWord = func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	var addEdge = func(word string) int {
		id1 := addWord(word)
		for k, v := range word {
			word = word[:k] + "*" + word[k+1:]
			id2 := addWord(word)
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			word = word[:k] + string(v) + word[k+1:]
		}
		return id1
	}

	for _, v := range wordList {
		addEdge(v)
	}
	endId, has := wordId[endWord]
	if !has {
		return 0
	}
	beginId := addEdge(beginWord)
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		id := queue[0]
		queue = queue[1:]
		if id == endId {
			return dist[id]/2 + 1
		}
		for _, v := range graph[id] {
			if dist[v] == math.MaxInt64 {
				dist[v] = dist[id] + 1
				queue = append(queue, v)
			}

		}
	}
	return 0

}
