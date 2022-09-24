package w2e

import "os"

// -- DecisionNode -- //

const avgLevel = 60

type DecisionNode struct {
	SimpleNode
	level                 int
	cumulativeProbability int

	// statistics
	count int
}

func (n *DecisionNode) Add(nodes ...Node) Node {
	n.SimpleNode.Add(nodes...)
	return n
}

func (n *DecisionNode) Exec(f func(node Node)) Node {
	f(n)
	return n
}

func (n *DecisionNode) Child() []Node {
	return n.Children
}

// Weights @param level from [1, 100]
func (n *SimpleNode) Weights(level int) *DecisionNode {
	if level < 1 || level > 100 {
		print("level need in 1 to 100\n")
		os.Exit(-1)
	}

	return &DecisionNode{
		*n,
		level,
		0,
		0,
	}
}
