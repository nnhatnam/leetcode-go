package _97_All_Paths_From_Source_to_Target

func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	dfs(graph, []int{0}, len(graph) - 1, &result)
	return result
}

func dfs(graph [][]int, path []int, target int, ans *[][]int) {

	currentNode := path[len(path) - 1]
	if currentNode == target {
		*ans = append(*ans, path)
		return
	}

	for _, next := range graph[currentNode] {
		newPath := make([]int, len(path) + 1)
		copy(newPath, path)
		newPath[len(newPath) - 1] = next

	}

	for _, next := range graph[currentNode] {
		newPath := make([]int, len(path) + 1)
		copy(newPath, path)
		newPath[len(newPath) - 1] = next

		dfs(graph, newPath, target, ans)
	}
}
