package courseScheduleII

func FindOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	graph := make([][]int, numCourses)
	degre := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		graph[prerequisite[1]] = append(graph[prerequisite[1]], prerequisite[0])
		degre[prerequisite[0]]++
	}

	for course, d := range degre {
		if d == 0 {
			result = append(result, course)
		}
	}

	for i := 0; i < len(result); i++ {
		course := result[i]
		for _, j := range graph[course] {
			degre[j]--
			if degre[j] == 0 {
				result = append(result, j)
			}
		}
	}
    
	if len(result) == numCourses {
		return result
	}
	return []int{}
}