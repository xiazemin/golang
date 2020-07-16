package search
const (
	RED   = true
	BLACK = false
)

// 节点
type RBNode struct {
	Color               bool // true == 红  false == 黑
	Parent, Left, Right *RBNode
	Value, Index        int
}

type RBTree struct {
	Root *RBNode
}

/*
* 对红黑树的节点(x)进行左旋转
*
* 左旋示意图(对节点 x 进行左旋)：
*      px                              px
*     /                               /
*    x                               y
*   /  \      --(左旋)-.           / \                #
*  lx   y                          x  ry
*     /   \                       /  \
*    ly   ry                     lx  ly
*
*
*/

func (t *RBTree) leftSpin(node *RBNode) {
	// 先提出自己的 右子
	y := node.Right

	// 自己的新右子 是前右子的左子
	node.Right = y.Left

	if nil != y.Left {
		y.Left.Parent = node
	}

	// 你以前的爹，就是我现在的爹
	y.Parent = node.Parent

	// 如果被旋转的节点是 之前树的根
	// 那么，新的跟就是 y 节点
	if nil == node.Parent {
		t.Root = y
	} else {
		// 被旋转的是普通节点时, 需要结合自身的父亲来确认自己之前是属于 左子还是右子
		if node.Parent.Left == node { // 被旋转节点之前是 左子时
			// 用 y 来作为之前 该节点父亲的 新左子
			node.Parent.Left = y
		} else { // 否则，是 右子
			node.Parent.Right = y
		}
	}

	// 将 node 设为 y 的左子
	y.Left = node
	// 将 y 设为 node 的新父亲
	node.Parent = y
}

/*
 * 对红黑树的节点(y)进行右旋转
 *
 * 右旋示意图(对节点 y 进行左旋)：
 *            py                               py
 *           /                                /
 *          y                                x
 *         /  \      --(右旋)-.            /  \                     #
 *        x   ry                           lx   y
 *       / \                                   / \                   #
 *      lx  rx                                rx  ry
 *
 */
func (t *RBTree) rightSpin(node *RBNode) {
	// 先提出自己的 左子
	x := node.Left
	node.Left = x.Right

	if nil != x.Left {
		x.Right.Parent = node
	}

	x.Parent = node.Parent

	// 如果被旋转的节点是 之前树的根
	// 那么，新的跟就是 x 节点
	if nil == node.Parent {
		t.Root = x
	} else {

		if node.Parent.Right == node {
			node.Parent.Right = x
		} else {
			node.Parent.Left = x
		}
	}

	x.Right = node

	node.Parent = x
}

func insertValue(tree *RBTree, val, index int) {
	node := &RBNode{Value: val, Index: index, Color: BLACK}
	if nil == tree.Root {
		tree.Root = node
	} else {
		tree.insert(node)
	}
}


func (t *RBTree) insert(node *RBNode) {
	//int cmp;
	var tmpNode *RBNode
	root := t.Root

	// 1. 将红黑树当作一颗二叉查找树，将节点添加到二叉查找树中。
	for nil != root {
		if node.Value < root.Value {
			root = root.Left
		} else {
			root = root.Right
		}
		tmpNode = root
	}

	node.Parent = tmpNode
	if nil != tmpNode {

		if node.Value < tmpNode.Value {
			tmpNode.Left = node
		} else {
			tmpNode.Right = node
		}
	} else {
		t.Root = node
	}

	// 2. 设置节点的颜色为红色
	node.Color = RED

	// 3. 将它重新修正为一颗二叉查找树
	t.adjustRBTree(node)
}

// 修正树
func (t *RBTree) adjustRBTree(node *RBNode) {
	var parent, gparent *RBNode // 父亲 和 祖父

	// 若“父节点存在，并且父节点的颜色是红色”
	for nil != node.Parent && RED == node.Parent.Color {
		parent = node.Parent
		gparent = parent.Parent //祖父不可能是空(否则父节点是黑色),也不可能是红色(否则不满足4),故一定是黑色

		//若“父节点”是“祖父节点的左孩子”
		if parent == gparent.Left {
			// Case 1条件：叔叔节点是红色
			uncle := gparent.Right
			if nil != uncle && RED == uncle.Color {
				uncle.Color = BLACK
				parent.Color = BLACK
				gparent.Color = RED
				node = gparent
				continue
			}

			// Case 2条件：叔叔是黑色，且当前节点是右孩子
			//左旋转的目的:变成     B
			//                  /  \
			//        node:    R    B
			//                /
			//      parent:  R
			//形式
			if node == parent.Right {
				var tmp *RBNode
				t.leftSpin(parent)
				tmp = parent
				parent = node
				node = tmp
			}
			//变成                B
			//                  /  \
			//      parent:    R    B
			//                /
			//        node:  R
			//形式

			// Case 3条件：叔叔是黑色，且当前节点是左孩子。
			parent.Color = BLACK
			gparent.Color = RED
			//变成                R
			//                  / \
			//                 B   B
			//                /
			//               R
			//形式
			t.rightSpin(gparent)
			//变成
			//                  B
			//                /  \
			//               R    R
			//                     \
			//                      B
			//形式
		} else { //若“z的父节点”是“z的祖父节点的右孩子”
			// Case 1条件：叔叔节点是红色
			uncle := gparent.Left
			if nil != uncle && RED == uncle.Color {
				uncle.Color = BLACK
				parent.Color = BLACK
				gparent.Color = RED
				node = gparent
				continue
			}

			// Case 2条件：叔叔是黑色，且当前节点是左孩子
			//右旋转的目的:变成    B
			//                  / \
			//                 B   R
			//                      \
			//                       R
			//形式
			if node == parent.Left {
				var tmp *RBNode
				t.rightSpin(parent)
				tmp = parent
				parent = node
				node = tmp
			}

			// Case 3条件：叔叔是黑色，且当前节点是右孩子。
			parent.Color = BLACK
			gparent.Color = RED
			t.leftSpin(gparent)
		}
	}
	// 将根节点设为黑色
	t.Root.Color = BLACK
}

/**
红黑树查找
 */
func RedBlackTreeSearch(arr []int, key int) int{
	// 先构造树
	tree := new(RBTree)
	for i, v := range arr {
		insertValue(tree, v, i)
	}
	// 开始二叉树查找目标key
	return tree.serch(key)
}

func (t *RBTree) serch(key int) int {
	return serch(t.Root, key)
}
func serch(node *RBNode, key int) int {
	if nil == node {
		return -1
	}
	if key < node.Value {
		return serch(node.Left, key)
	} else if key > node.Value {
		return serch(node.Right, key)
	} else {
		return node.Index
	}
}