package cloneGraph

type Node struct {
    Val int
	Neighbors []*Node
}

// bfs
func CloneGraphBFS(node *Node) *Node {
    if node == nil {
        return nil
    }
    record := make(map[*Node]*Node)
    queue := []*Node{node}
    record[node] = &Node{Val: node.Val}

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, child := range current.Neighbors {
            if _, exists := record[child]; !exists {
                record[child] = &Node{Val: child.Val}
                queue = append(queue, child)
            }
            record[current].Neighbors = append(record[current].Neighbors, record[child])
        }
    }

    return record[node]
}

// dfs
func CLoneGraphDFS(node *Node) *Node {
    if node == nil {
        return nil
    }
    record := make(map[*Node]*Node)

    var dfs func(*Node)
    dfs = func(node *Node) {
        if node == nil {
            return
        }

        newNode := new(Node)
        newNode.Val = node.Val
        record[node] = newNode

        for _, child := range node.Neighbors {
            if _, exists := record[child]; !exists {
                dfs(child)
            }
            newNode.Neighbors = append(newNode.Neighbors, record[child])
        }
    }

    dfs(node)
    return record[node]
}