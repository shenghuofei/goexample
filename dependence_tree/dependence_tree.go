package main

import (
	"fmt"
	"strings"
)


/*
如depModel所定义，src依赖Depend,可扩展依赖的具体内容，比如资源量的依赖关系，补充上依赖系数，即可根据源的资源量计算出所需依赖的资源量
依赖数据结构如下：
deps := []depModel{
		{"A", "B"},
		{"A", "F"},
		{"B", "C"},
		{"B", "F"},
		{"C", "D"},
		{"D", "E"},
		{"E", "B"}, // 这里是一个循环依赖
}
根据依赖关系生成的依赖数据结构如下：
nodes := []*DepData{
		{ID: 1, Src: "A", RootID: 0, ParentID: 0, Depend: "B"},
		{ID: 2, Src: "B", RootID: 1, ParentID: 1, Depend: "C"},
		{ID: 3, Src: "C", RootID: 1, ParentID: 2, Depend: "D"},
		{ID: 4, Src: "B", RootID: 1, ParentID: 1, Depend: "F"},
		{ID: 5, Src: "A", RootID: 0, ParentID: 0, Depend: "F"},
		{ID: 6, Src: "F", RootID: 5, ParentID: 5, Depend: "G"},
		//...
}
要生成依赖的数据，首先要从顶层找到每一个节点的依赖路径，并要检查是否有循环依赖
生成数据时需要记录根节点及父节点，然后支持按层级展示，按层级可选展示根和所有子节点两层，也可以从根节点逐层展示各层子节点
这里实现有：
1.根据依赖关系生成全部依赖路径的实现
2.检查依赖是否有环的实现
3.展示根和其所有子节点两层展示实现
4.展示根节点及逐层展示各层子节点实现
*/


// 定义结构体 dep
type depModel struct {
	Src    string `json:"src"`
	Depend string `json:"depend"`
}

// 构建依赖图
func buildGraph(deps []depModel) map[string][]string {
	graph := make(map[string][]string)
	for _, d := range deps {
		graph[d.Src] = append(graph[d.Src], d.Depend)
	}
	return graph
}

// 非递归深度优先搜索  生成所有依赖路径
func nonRecursiveDFS(graph map[string][]string, startNode string, results map[string][][]string) {
	type stackItem struct {
		node    string          // 当前节点
		path    []string        // 记录当前节点依赖路径
		visited map[string]bool // 记录访问节点
	}

	stack := []stackItem{{startNode, []string{startNode}, map[string]bool{startNode: true}}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := item.node
		path := item.path
		visited := item.visited

		if neighbors, exists := graph[node]; exists {
			for _, neighbor := range neighbors {
				if _, has := visited[neighbor]; has {
					// 检测到循环依赖
					//results[startNode] = append(results[startNode], append(path, neighbor))
					// 如果循环依赖的节点不需要在路径中，path不要添加当前节点即可,如下：
					results[startNode] = append(results[startNode], path)
					continue
				}

				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				newVisited := copyVisited(visited)
				newVisited[neighbor] = true
				stack = append(stack, stackItem{neighbor, newPath, newVisited})
			}
		} else {
			results[startNode] = append(results[startNode], path)
		}
	}
}

// 非递归深度优先搜索,查找全部依赖路径,并检查是否有循环依赖, 如果不同业务依赖模型不同，且要按业务优先找模型，则使用这里的方法，生成唯一key时加上业务名
//func nonRecursiveDFSGetProductDependencePath(graph map[string][]string, startNode string, results map[string][][]string) (bool, string) {
//	type stackItem struct {
//		node    string          // 当前组件
//		path    []string        // 记录当前组件依赖路径
//		visited map[string]bool // 记录已访问节点
//	}
//
//	stack := []stackItem{
//		{startNode, []string{startNode}, map[string]bool{startNode: true}},
//	}
//
//	for len(stack) > 0 {
//		item := stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		node := item.node
//		path := item.path
//		visited := item.visited
//
//		defaultNode := ""
//		if !strings.HasPrefix(node, "default") {
//			nodeList := strings.Split(node, ",")
//			defaultNode = strings.Join(append([]string{"default"}, nodeList[1:]...), ",")
//		}
//
//		if neighbors, exists := graph[node]; exists {
//			// 组件依赖其他组件
//			for _, neighbor := range neighbors {
//				//if productDependencePathContains(path, neighbor) {
//				if _, has := visited[neighbor]; has {
//					// 检测到循环依赖,直接返回
//					results[startNode] = append(results[startNode], append(path, neighbor))
//					// 如果循环节点不需要记录在path里，则这里path不用加循环节点neighbor，如下:
//					// results[startNode] = append(results[startNode], path)
//					circularDependency := strings.Join(path, " -> ") + " -> " + neighbor
//					log.Info("circularDependency path: %s", circularDependency)
//					// 如果要获取所有的依赖路径，这里continue就行了，不用return
//					// continue
//					return true, circularDependency
//				}
//
//				newPath := append([]string{}, path...)
//				newPath = append(newPath, neighbor)
//				newVisited := copyVisitedProductDependenceNode(visited)
//				newVisited[neighbor] = true
//				stack = append(stack, stackItem{neighbor, newPath, newVisited})
//			}
//		} else if neighbors1, defaultExists := graph[defaultNode]; defaultNode != "" && defaultExists {
//			// 组件依赖其他组件
//			for _, neighbor := range neighbors1 {
//				//if productDependencePathContains(path, neighbor) {
//				if _, has := visited[neighbor]; has {
//					// 检测到循环依赖,直接返回
//					results[startNode] = append(results[startNode], append(path, neighbor))
//					// 如果循环节点不需要记录在path里，则这里path不用加循环节点neighbor，如下:
//					// results[startNode] = append(results[startNode], path)
//					circularDependency := strings.Join(path, " -> ") + " -> " + neighbor
//					log.Info("circularDependency path: %s", circularDependency)
//					// 如果要获取所有的依赖路径，这里continue就行了，不用return
//					// continue
//					return true, circularDependency
//				}
//
//				newPath := append([]string{}, path...)
//				newPath = append(newPath, neighbor)
//				newVisited := copyVisitedProductDependenceNode(visited)
//				newVisited[neighbor] = true
//				stack = append(stack, stackItem{neighbor, newPath, newVisited})
//			}
//		} else {
//			// 组件不依赖其他组件，记录依赖路径
//			results[startNode] = append(results[startNode], path)
//		}
//	}
//	return false, ""
//}

// 非递归深度优先搜索  检查是否有循环依赖
func nonRecursiveDFSCheckCircle(graph map[string][]string, startNode string, results map[string][][]string) (bool, string) {
	type stackItem struct {
		node    string          // 当前节点
		path    []string        // 记录当前节点依赖路径
		visited map[string]bool // 记录访问节点
	}

	stack := []stackItem{{startNode, []string{startNode}, map[string]bool{startNode: true}}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node := item.node
		path := item.path
		visited := item.visited

		if neighbors, exists := graph[node]; exists {
			for _, neighbor := range neighbors {
				if _, has := visited[neighbor]; has {
					// 检测到循环依赖
					results[startNode] = append(results[startNode], append(path, neighbor))
					circularDependency := strings.Join(path, " -> ") + " -> " + neighbor
					// 如果循环依赖的节点不需要在路径中，path不要添加当前节点即可,如下：
					//results[startNode] = append(results[startNode], path)
					return true, circularDependency
				}

				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				newVisited := copyVisited(visited)
				newVisited[neighbor] = true
				stack = append(stack, stackItem{neighbor, newPath, newVisited})
			}
		} else {
			results[startNode] = append(results[startNode], path)
		}
	}
	return false, ""
}

// 检查路径中是否包含节点
func pathContains(path []string, node string) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

// 复制访问节点记录
func copyVisited(visited map[string]bool) map[string]bool {
	newVisited := make(map[string]bool)
	for k, v := range visited {
		newVisited[k] = v
	}
	return newVisited
}

func findDependencies(deps []depModel) map[string][][]string {
	graph := buildGraph(deps)
	Results := make(map[string][][]string)

	for node := range graph {
		nonRecursiveDFS(graph, node, Results)
	}

	return Results
}

func HasCircleDep(deps []depModel) (bool, string) {
	graph := buildGraph(deps)
	Results := make(map[string][][]string)

	for node := range graph {
		hasCircle, circlePath := nonRecursiveDFSCheckCircle(graph, node, Results)
		if hasCircle {
			return hasCircle, circlePath
		}
	}
	return false, ""
}

func CheckCircleDep() {
	// 示例输入
	deps := []depModel{
		{"A", "B"},
		{"A", "F"},
		{"B", "C"},
		{"B", "F"},
		{"C", "D"},
		{"D", "E"},
		{"E", "B"}, // 这里是一个循环依赖
	}

	// 检查是否有循环依赖
	hasCircle, circlePath := HasCircleDep(deps)
	if hasCircle {
		fmt.Println("存在循环依赖:", circlePath)
	} else {
		fmt.Println("没有循环依赖")
	}

	// 调用 findDependencies 函数,获取全路径依赖
	results := findDependencies(deps)

	// 输出结果
	fmt.Println("全路径依赖：")
	fmt.Println(results)
	for src, paths := range results {
		fmt.Printf("%s:\n", src)
		for _, path := range paths {
			fmt.Printf("  -> %v\n", path)
		}
	}

}

// 根据依赖模型生成的依赖数据，按依赖关系生成依赖树
type DepData struct {
	ID          int
	Src         string
	Depend      string
	RootID      int
	ParentID    int
	ResourceNum float64
	Children    []*DepData
}

// DepDataEveryLevelTree 按依赖关系生成依赖树,每一层之间的关系都要显示
func DepDataEveryLevelTree() {
	nodes := []*DepData{
		{ID: 1, Src: "A", RootID: 0, ParentID: 0, Depend: "B"},
		{ID: 2, Src: "B", RootID: 1, ParentID: 1, Depend: "C"},
		{ID: 3, Src: "C", RootID: 1, ParentID: 2, Depend: "D"},
		{ID: 4, Src: "B", RootID: 1, ParentID: 1, Depend: "F"},
		{ID: 5, Src: "A", RootID: 0, ParentID: 0, Depend: "F"},
		{ID: 6, Src: "F", RootID: 5, ParentID: 5, Depend: "G"},
		//...
	}

	nodeMap := mapNodesByParent(nodes)
	rootNodes := buildTree(nodeMap, 0) // 假设所有 root 节点的 ParentID = 0
	for i := range rootNodes {
		fmt.Printf("data %d tree: %v\n", i, rootNodes[i])
	}
	printNodes(rootNodes, 0)
}

func mapNodesByParent(nodes []*DepData) map[int][]*DepData {
	nodeMap := make(map[int][]*DepData)
	for _, node := range nodes {
		nodeMap[node.ParentID] = append(nodeMap[node.ParentID], node)
	}
	return nodeMap
}

func buildTree(nodeMap map[int][]*DepData, parentID int) []*DepData {
	nodes := nodeMap[parentID]
	for _, node := range nodes {
		node.Children = buildTree(nodeMap, node.ID)
	}
	return nodes
}

func printNodes(nodes []*DepData, level int) {
	for i, node := range nodes {
                if level == 0 {
			fmt.Printf("data %d tree: %v\n", i, node)
		}
		tabs := strings.Repeat("\t", level)
		fmt.Printf("%s%d, %s, %s\n", tabs, node.ID, node.Src, node.Depend)
		printNodes(node.Children, level+1)
	}
}

// DepDataTwoLevelTree 按依赖关系生成依赖树,只显示根节点和所有子节点两层之间的关系
func DepDataTwoLevelTree() {
	nodes := []*DepData{
		{ID: 1, Src: "A", RootID: 0, ParentID: 0, Depend: "B"},
		{ID: 2, Src: "B", RootID: 1, ParentID: 1, Depend: "C"},
		{ID: 3, Src: "C", RootID: 1, ParentID: 2, Depend: "D"},
		{ID: 4, Src: "B", RootID: 1, ParentID: 1, Depend: "F"},
		{ID: 5, Src: "A", RootID: 0, ParentID: 0, Depend: "F"},
		{ID: 6, Src: "F", RootID: 5, ParentID: 5, Depend: "G"},
		//...
	}

	nodeMap := mapNodesByRoot(nodes)
	rootNodes := buildTreeByRoot(nodeMap, nodes)
	for i := range rootNodes {
		fmt.Printf("data %d tree: %v\n", i, rootNodes[i])
		for _, children := range rootNodes[i].Children {
			fmt.Printf("children: %v\n", children)
		}
	}
}

func buildTreeByRoot(nodeMap map[int][]*DepData, nodes []*DepData) []*DepData {
	rootNodes := make([]*DepData, 0)
	for _, node := range nodes {
		if node.RootID == 0 {
			rootNodes = append(rootNodes, node)
		}
		node.Children = nodeMap[node.ID]
	}
	return rootNodes
}

func mapNodesByRoot(nodes []*DepData) map[int][]*DepData {
	nodeMap := make(map[int][]*DepData)
	for _, node := range nodes {
		// 根节点不记录在map中
		if node.RootID == 0 {
			continue
		}
		nodeMap[node.RootID] = append(nodeMap[node.RootID], node)
	}
	return nodeMap
}

func main() {
    	CheckCircleDep()
	DepDataEveryLevelTree()
	DepDataTwoLevelTree()
}
