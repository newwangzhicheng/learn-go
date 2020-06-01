package gee

// 这是一个节点
type node struct {
	pattern  string  // 待匹配的路由，如/p/:lang
	part     string  // 本体，当前节点路由的部分，如p
	children []*node // 子节点
	isWild   bool    //是否是精确匹配，如果节点包含:或*，true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	// range数组的时候，第一个参数是key，第二个参数是value
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点，用于查询
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
