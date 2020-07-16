package search

import "fmt"

func main() {
	/**
	图的形式:
		var r = make(map[string][]string)
		r["A"] = []string{"B", "C", "D"}
		r["B"] = []string{"A", "E"}
		r["C"] = []string{"A", "E"}
		r["D"] = []string{"A"}
		r["E"] = []string{"B", "C", "F"}
		r["F"] = []string{"E"}
	 */

	// 构造图
	fillGraph()

	// 广度遍历
	g.BFS(func(node *TNode) {
		fmt.Printf("B-F-S visiting... %v\n", node)
	})

	// 深度遍历
	g.DFS(func(node *TNode) {
		fmt.Printf("DFS visiting... %v\n", node)
	})
}


/**
图按照不同的特征可以分为以下类别:
- 有向图
- 无向图
- 有权图
- 无权图
- 连通图
- 非连通图
 */

/**
图的两种表示方法
邻接矩阵
var s = [][]int{
{0, 1, 1, 1, 0, 0},
{1, 0, 0, 0, 1, 0},
{1, 0, 0, 0, 1, 0},
{1, 0, 0, 0, 0, 0},
{0, 1, 1, 0, 0, 0},
{0, 0, 0, 0, 1, 0},
}
邻接表表示
var r = make(map[string][]string)
r["A"] = []string{"B", "C", "D"}
r["B"] = []string{"A", "E"}
r["C"] = []string{"A", "E"}
r["D"] = []string{"A"}
r["E"] = []string{"B", "C", "F"}
r["F"] = []string{"E"}
*/




// 组成图的顶点
type TNode struct {
	Value interface{}
}

// 定义一个图的结构, 图有顶点与边组成 V  E
type ItemGraph struct {
	Nodes []*TNode 				// 	顶点 集合
						      /**  采用 邻接表 */
	Edges map[TNode][]*TNode	//	边 集合
}

// 添加节点
func (g *ItemGraph) AddNode(n *TNode)  {
	g.Nodes = append(g.Nodes, n)
}

// 添加边
func (g *ItemGraph) AddEdge(n1, n2 *TNode)  {
	if g.Edges == nil{
		g.Edges = make(map[TNode][]*TNode)
	}
	// 无向图
	g.Edges[*n1] = append(g.Edges[*n1], n2)    // 设定从节点n1到n2的边
	g.Edges[*n2] = append(g.Edges[*n2], n1)    // 设定从节点n2到n1的边
}

// 打印图
func (g *ItemGraph) String()  {
	s := ""
	for i := 0; i< len(g.Nodes); i++{
		s += g.Nodes[i].String() + "->"
		near := g.Edges[*g.Nodes[i]]

		for j :=0; j<len(near); j++{
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (n *TNode) String() string  {
	return fmt.Sprintf("%v", n.Value)
}

// 深度用栈，广度用队列
// 深度优先遍历可以通过递归实现，而广度优先遍历要转换成类似于树的层序遍历来实现

/////////////////////////////////////////////////////////// 广度遍历 ///////////////////////////////////////////////////////////


/**
首先bfs 广度优先搜索
此处结合队列实现图的广度优先遍历
广度优先遍历，这个遍历类似于层序遍历，每遍历一层，需要记住当前层的节点，
然后与遍历当前层相连的节点，如此实现遍历。需要一个队列来记住当前层，先进先出。
还有一个问题，就是需要防止回路，也就是说，一个节点不能遍历两次。这里用了Golang内置的map实现
*/

// 定义缓冲队列
type NodeQueue struct {
	Items []TNode
}

func (s *NodeQueue) New() {
	s.Items = make([]TNode, 0)
}

// 从队尾入队
func (s *NodeQueue) Enqueue(t TNode)  {
	s.Items = append(s.Items, t)
}

// 从对头出队
func (s *NodeQueue) Dequeue() *TNode {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

// 是否空队列
func (s *NodeQueue) IsEmpty() bool  {
	return len(s.Items) == 0
}

/**
广度优先 BFS
 */
func (g *ItemGraph) BFS(f func(node *TNode))  {
	// 初始化队列
	q := new(NodeQueue)
	q.New()
	// 获取图的首个节点
	n := g.Nodes[0]
	// 首节点入队
	q.Enqueue(*n)
	// 缓存是否被某节点是否被遍历过的检查
	visited := make(map[*TNode]bool)
	// 初始时， 首节点标识位遍历过
	visited[n] = true

	for {
		// 如果全部遍历完时，缓冲队列为空
		if q.IsEmpty(){
			break
		}
		// 节点出队
		node := q.Dequeue()
		// 获取当前节点的边，即和该节点相邻的其他节点
		nearArr := g.Edges[*node]
		// 先遍历完该节点所有相邻的节点 (也就是这一层的)
		for i :=0; i < len(nearArr); i++{
			j := nearArr[i]
			// 只要是该节点没有被曾经遍历过，即可入队
			if !visited[j]{
				q.Enqueue(*j)
				visited[j] = true
			}
		}
		if f != nil{
			// 打印当前节点信息
			f(node)
		}
	}
}

/**
BFS 广度优先搜索
此处结合队列实现图的广度优先遍历
广度优先遍历，这个遍历类似于层序遍历，每遍历一层，需要记住当前层的节点，
然后与遍历当前层相连的节点，如此实现遍历。需要一个队列来记住当前层，先进先出。
还有一个问题，就是需要防止回路，也就是说，一个节点不能遍历两次。这里用了Golang内置的map实现
*/
/**
DFS 深度优先搜索
此处结合栈实现图的深度优先遍历
深度优先遍历， 是沿着一个方向先遍历到底，再沿另外的方向一路到底
 */

// 定义缓冲栈
type NodeStack struct {
	Items []TNode
}


func (n *NodeStack) New() {
	n.Items = make([]TNode, 0)
}

// 压栈
func (n *NodeStack) push(q TNode) {
	n.Items = append(n.Items, q)
}

// 弹栈
func (n *NodeStack) pop() *TNode {
	item := n.Items[len(n.Items) - 1] //取最后一个
	n.Items = n.Items[0: len(n.Items) - 1]
	return &item
}

// 判断是否空栈
func (n *NodeStack) IsEmpty() bool {
	return len(n.Items) == 0
}

// 栈宽
func (n *NodeStack) Size() int {
	return len(n.Items)
}


/**
深度优先 DFS
 */
func (g *ItemGraph) DFS(f func(node *TNode)) {
	// 初始化栈
	stack := new(NodeStack)
	stack.New()
	// 取出图的首节点
	n := g.Nodes[0]
	// 压栈
	stack.push(*n)
	// 缓存是否被某节点是否被遍历过的检查
	visited := make(map[*TNode] bool)
	// 初始时， 首节点标识位遍历过
	visited[n] = true
	for {
		// 如果全部遍历完时，缓冲栈为空
		if stack.IsEmpty(){
			break
		}
		// 弹栈
		node := stack.pop()
		/** 注意：这里和 广度的不一样，因为在深度中，当前这一层的某个节点可能在之前的某次深度中已经被遍历过了 */
		if !visited[node]{
			visited[node] = true
		}
		// 获取当前节点的边，即和该节点相邻的其他节点
		nearArr := g.Edges[*node]
		for i:= 0; i< len(nearArr); i++{
			j := nearArr[i]
			if !visited[j]{
				visited[j] = true
				stack.push(*j)
			}
		}
		if f != nil{
			// 打印当前节点信息
			f(node)
		}
	}
}



/**
测试
 */
var g ItemGraph

func fillGraph() {
	// 构造所有游离态的节点
	nA := TNode{"A"}
	nB := TNode{"B"}
	nC := TNode{"C"}
	nD := TNode{"D"}
	nE := TNode{"E"}
	nF := TNode{"F"}

	// 往图中聚集节点
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	// 添加边；图中的节点通过 链表关联起来
	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

}