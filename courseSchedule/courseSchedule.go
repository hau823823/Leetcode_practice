package courseSchedule

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	degre := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		degre[prerequisite[0]] += 1
	}

	bfs := make([]int, 0)
	for course, d := range degre {
		if d == 0 {
			bfs = append(bfs, course)
		}
	}

	for i := 0; i < len(bfs); i++ {
		course := bfs[i]
		for _, j := range graph[course] {
			degre[j] -= 1
			if degre[j] == 0 {
				bfs = append(bfs, j)
			}
		}
	}

	return len(bfs) == numCourses
}

func CanFinishDFS(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
	}

	todo := make([]bool, numCourses) // todo 标记正在访问的节点
	done := make([]bool, numCourses) // done 标记已访问完成的节点
	var dfs func(int) bool           // dfs  检查是否存在环
	dfs = func(node int) bool {
		if todo[node] {
			return false
		}
		if done[node] {
			return true
		}
		todo[node], done[node] = true, true
		for _, v := range graph[node] {
			if !dfs(v) {
				return false
			}
		}
		todo[node] = false
		return true
	}

	for i := 0; i < numCourses; i++ {
		if !done[i] && !dfs(i) {
			return false
		}
	}
	return true
}
