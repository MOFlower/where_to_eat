package w2e

type Node interface {
	Add(nodes ...Node) Node
	Exec(f func(node Node, args ...interface{}), args ...interface{}) Node
	Child() []Node
}

// -- simpleNode -- //

type SimpleNode struct {
	Name     string
	Children []Node
}

func (n *SimpleNode) Add(nodes ...Node) Node {
	for _, node := range nodes {
		n.Children = append(n.Children, node)
	}
	return n
}

func (n *SimpleNode) Exec(f func(node Node, args ...interface{}), args ...interface{}) Node {
	f(n, args)
	return n
}

func (n *SimpleNode) Child() []Node {
	return n.Children
}

func getSimpleNode(n Node) *SimpleNode {
	var node *SimpleNode
	switch n.(type) {
	case *SimpleNode:
		node = n.(*SimpleNode)
	case *DecisionNode:
		node = &(n.(*DecisionNode).SimpleNode)
	}
	return node
}
