package mygin

import "fmt"

type Trees map[string]*node

type node struct {
	path     string
	indices  string
	children []*node
	handlers HandlerList
}

func (n *node) addRoute(path string, handlers HandlerList) {
	if len(n.path) > 0 || len(n.children) > 0 {
		walk:
		for {
			//找到相等的index
			i := 0
			max := min(len(path), len(n.path))
			for max > i && path[i] == n.path[i] {
				i++
			}
			//需要把原来的作为子node放到新node中
			if i < len(n.path) {
				//新建node
				child := node{
					path:     n.path[i:],
					indices:  n.indices,
					handlers: n.handlers,
					children: n.children,
				}

				n.children = []*node{&child}
				n.indices = string([]byte{n.path[i]})
				n.path = path[:i]
				n.handlers = nil
			}
			// 判断子节点如果有相同开头的字符 则从新跳入循环
			if i < len(path) {
				c := path[i]
				for index := 0; index < len(n.indices); index++ {
					if c == n.indices[index] {
						n = n.children[index]
						path = path[i:]
						continue walk
					}
				}

				//把新请求的path加入到router中
				n.insertChild(path[i:], path, handlers, i)
				return
			}
			return
		}
	} else {
		//如果为空
		n.path = path
		n.handlers = handlers
	}
}

func (n *node) insertChild(path, fullPath string, handlers HandlerList, index int) {
	child := node{}
	child.handlers = handlers
	child.indices = ""
	child.path = path
	n.indices += string([]byte{fullPath[index]})
	n.children = append(n.children, &child)
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func (n *node) getValue(path string) (handlers HandlerList) {
	index := 1
	walk:
	for {
		fmt.Println("loop num: ", index)
		if len(path) > len(n.path) {
			path = path[len(n.path):]
			c := path[0]
			for i := 0; i < len(n.indices); i++ {
				if c == n.indices[i] {
					n = n.children[i]
					index++
					goto walk
				}
			}
		} else if len(path) == len(n.path) {
			handlers = n.handlers
			return
		}
	}
}