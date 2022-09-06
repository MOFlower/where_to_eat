package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Node interface {
	add(nodes ...Node) Node
	exec(f func(node Node)) Node
	child() []Node
}

// -- SimpleNode -- //

type SimpleNode struct {
	name     string
	children []Node
}

func (n *SimpleNode) add(nodes ...Node) Node {
	for _, node := range nodes {
		n.children = append(n.children, node)
	}
	return n
}

func (n *SimpleNode) exec(f func(node Node)) Node {
	f(n)
	return n
}

func (n *SimpleNode) child() []Node {
	return n.children
}

// -- DecisionNode -- //

const avgLevel = 60

type DecisionNode struct {
	SimpleNode
	level                 int
	cumulativeProbability int

	// statistics
	count int
}

func (n *DecisionNode) add(nodes ...Node) Node {
	n.SimpleNode.add(nodes...)
	return n
}

func (n *DecisionNode) exec(f func(node Node)) Node {
	f(n)
	return n
}

func (n *DecisionNode) child() []Node {
	return n.children
}

// -- other -- //

func (n SimpleNode) weights(level int) *DecisionNode {
	if level < 1 || level > 100 {
		print("level need in 1 to 100\n")
		os.Exit(-1)
	}

	return &DecisionNode{
		n,
		level,
		0,
		0,
	}
}

func newNode(name string) *SimpleNode {
	children := make([]Node, 0)
	return &SimpleNode{
		name,
		children,
	}
}
func newRoot() Node {
	root := newNode("where to eat:").weights(1)
	newNList := func() []Node {
		return []Node{
			newNode("一楼").weights(30),
			newNode("二楼"),
			newNode("三楼"),
		}
	}
	return root.add(
		newNode("东苑").weights(80).add(
			newNode("东一").add(newNList()...),
			newNode("东二").add(newNList()...),
			newNode("芳缘"),
		),
		newNode("西苑").add(newNList()[:2]...),
		newNode("软件园").weights(40).add(newNList()[:2]...),
	)
}

func getSimpleNode(n Node) *SimpleNode {
	var node *SimpleNode
	switch n.(type) {
	case *SimpleNode:
		node = n.(*SimpleNode)
		break
	case *DecisionNode:
		node = &(n.(*DecisionNode).SimpleNode)
		break
	}
	return node
}

// convert simpleNode to decisionNode
func genProbabilityFunc(n Node) {
	sum := 0
	node := getSimpleNode(n)
	for index, item := range node.children {
		switch item.(type) {
		case *SimpleNode:
			sum += avgLevel
			node.children[index] = &DecisionNode{
				*(item.(*SimpleNode)),
				avgLevel,
				sum,
				0,
			}
			break
		case *DecisionNode:
			sum += item.(*DecisionNode).level
			item.(*DecisionNode).cumulativeProbability += sum
			break
		}

		item.exec(genProbabilityFunc)
	}
}

func randSelect(n Node) {
	node := getSimpleNode(n)
	print(node.name)
	if len(node.children) <= 0 {
		print("\n")
		return
	}
	print(" -> ")

	sum := node.children[len(node.children)-1].(*DecisionNode).cumulativeProbability
	r := int(rand.Int31n(int32(sum)))
	for _, item := range node.children {
		if item.(*DecisionNode).cumulativeProbability > r {
			item.(*DecisionNode).count++
			item.exec(randSelect)
			break
		}
	}
}

func printDecisionTree(n Node) {
	sum := 0
	sum2 := 0
	for _, item := range getSimpleNode(n).children {
		sum += item.(*DecisionNode).count
		sum2 += item.(*DecisionNode).level
		// fmt.Printf("(%s, %d, %d, %d) ",
		// 	item.(*DecisionNode).name,
		// 	item.(*DecisionNode).level,
		// 	item.(*DecisionNode).cumulativeProbability,
		// 	item.(*DecisionNode).count)
	}
	for _, item := range getSimpleNode(n).children {
		i := item.(*DecisionNode)
		fmt.Printf("name := %s, target := %.2f%%, now := %.2f%%, diff := %.2f%%\n",
			i.name,
			float64(i.level)/float64(sum2)*100,
			float64(i.count)/float64(sum)*100,
			(float64(i.level)/float64(sum2)-float64(i.count)/float64(sum))*100,
		)
	}
	if len(getSimpleNode(n).children) > 0 {
		println("-----------------------------------" +
			"---------------------------------------")
	}
	for _, item := range getSimpleNode(n).children {
		item.exec(printDecisionTree)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	r := newRoot().exec(genProbabilityFunc)
	cnt := rand.Intn(10000) + 1000
	for i := 1; i < cnt; i++ {
		r.exec(randSelect)
	}
	println(cnt)
	r.exec(printDecisionTree)
}
