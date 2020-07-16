package search
/**
基本思路：先把数组构造出一颗二叉树的样纸，然后查找的时候在从root往下对比
*/
func BSTsearch(arr []int, key int) int {
	// 先在内存中构造 二叉树
	tree := new(Tree)
	for i, v := range arr {
		Insert(tree, v, i)
	}
	// 开始二叉树查找目标key
	return searchKey(tree.Root, key)
}

// 节点结构
type Node struct {
	Value, Index int // 元素的值和在数组中的位置
	Left, Right  *Node
}

// 树结构
type Tree struct {
	Root *Node
}

// 把数组的的元素插入树中
func Insert(tree *Tree, value, index int) {
	if nil == tree.Root {
		tree.Root = newNode(value, index)
	} else {
		InsertNode(tree.Root, newNode(value, index))
	}
}

// 把新增的节点插入树的对应位置
func InsertNode(root, childNode *Node) {
	// 否则，先和根的值对比
	if childNode.Value <= root.Value {
		// 如果小于等于跟的值，则插入到左子树
		if nil == root.Left {
			root.Left = childNode
		} else {
			InsertNode(root.Left, childNode)
		}
	} else {
		// 否则，插入到右子树
		if nil == root.Right {
			root.Right = childNode
		} else {
			InsertNode(root.Right, childNode)
		}
	}
}

func newNode(value, index int) *Node {
	return &Node{
		Value: value,
		Index: index,
	}
}

// 在构建好的二叉树中，从root开始往下查找对应的key 返回其在数组中的位置
func searchKey(root *Node, key int) int {
	if nil == root {
		return -1
	}
	if key == root.Value {
		return root.Index
	} else if key < root.Value {
		// 往左子树查找
		return searchKey(root.Left, key)
	} else {
		// 往右子树查找
		return searchKey(root.Right, key)
	}
	return -1
}
