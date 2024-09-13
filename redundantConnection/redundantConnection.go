package redundantConnection

import "fmt"

// 初步想法：找到 cycle 的地方，就是可以刪掉的
// dfs
func FindRedundantConnection(edges [][]int) []int {
	alreadyConnected := make(map[int][]int)
	visited := make(map[int]bool)

	var isAlreadyConnected func(int, int) bool
	isAlreadyConnected = func(x, y int) bool {
		if x == y {
			return true
		}
		for _, xAdjacent := range alreadyConnected[x] {
			if !visited[xAdjacent] {
				visited[xAdjacent] = true
				if isAlreadyConnected(xAdjacent, y) {
					return true
				}
			}
		}
		return false
	}

	for _, edge := range edges {
		x, y := edge[0], edge[1]
		visited = make(map[int]bool)
		if isAlreadyConnected(x, y) {
			return edge
		}
		alreadyConnected[x] = append(alreadyConnected[x], y)
		alreadyConnected[y] = append(alreadyConnected[y], x)
	}
	return []int{}
}

// union find
func FindRedundantConnectionUnion(edges [][]int) []int {
	// 初始化一个 parent 数组，这个数组用来表示每个节点的父节点
	// 初始化时，每个节点的父节点是它自己，表示每个节点最初都是独立的集合
	var parents = make([]int, len(edges)+1)
	for i := 0; i < len(edges)+1; i++ {
		parents[i] = i
	}

	// 对于给定的每条边，调用 doUnion 函数尝试合并两个端点
	// 如果两个端点已经在同一个集合中，表示添加这条边会形成环，因此这条边是冗余的。
	var result []int
	for i := 0; i < len(edges); i++ {
		doUnion(edges[i][0], edges[i][1], parents, &result)
	}

	return result
}

// 查找并合并：通过调用 doFind 函数查找两个端点 i 和 j 的根节点（即代表元）。
// 如果两者的根节点不同，说明它们之前不在同一个集合中，将它们合并（此处采用将一个根节点的父指针指向另一个的简单方法）。
// 如果两者的根节点相同，说明它们已经在同一个集合中，添加这条边会形成环，因此将这条边添加到 result 中。
func doUnion(i, j int, parents []int, result *[]int) {
    fmt.Printf("i: %d, j: %d\n", i, j)
	u1, u2 := doFind(i, parents), doFind(j, parents)
	fmt.Printf("u1:%d, u2:%d\n\n", u1, u2)
	if u1 != u2 {
		parents[u2] = u1
	} else {
		*result = append(*result, i)
		*result = append(*result, j)
	}
}

// 路径压缩：这是一个递归函数，用于查找节点 i 的根节点。如果节点 i 的父节点是其自身，说明找到了根节点
func doFind(i int, parents []int) int {
	if parents[i] == i {
		return i
	}
	return doFind(parents[i], parents)
}

// optimize
func FindRedundantConnectionOp(edges [][]int) []int {
    parent := make([]int, len(edges) + 1)
    for i := 0; i < len(edges) + 1; i++ {
        parent[i] = -1
    }

    rank := make([]int, len(edges) + 1)
    result := []int{}
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        pu, pv := findParent(u, parent), findParent(v, parent)

        if pu == pv {
            result = append(result, edge...)
            break
        }

        if rank[pu] > rank[pv]{
            parent[pv] =pu
        }else if rank[pu] < rank[pv]{
            parent[pu] = pv
        }else{
            rank[pu]++
            parent[pv] = pu
        }
    }

    return result
}

func findParent(i int, parent []int) int {
    if parent[i] == -1 {
        return i
    }

    return findParent(parent[i], parent)
}