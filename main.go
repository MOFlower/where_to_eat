package main

import (
	"math/rand"
	"time"
	. "where_to_eat/w2e"
)

func NewNode(name string) *SimpleNode {
	children := make([]Node, 0)
	return &SimpleNode{
		name,
		children,
	}
}

func newRoot() Node {
	root := NewNode("where to eat:").Weights(1)
	newNList := func() []Node {
		return []Node{
			NewNode("一楼").Weights(30),
			NewNode("二楼"),
			NewNode("三楼"),
		}
	}
	return root.Add(
		NewNode("东苑").Weights(80).Add(
			NewNode("东一").Add(newNList()...),
			NewNode("东二").Add(newNList()...),
			NewNode("芳缘"),
		),
		NewNode("西苑").Add(newNList()[:2]...),
		NewNode("软件园").Weights(40).Add(newNList()[:2]...),
	)
}

func main() {
	rand.Seed(time.Now().Unix())
	r := newRoot().Exec(GenProbabilityFunc)
	cnt := rand.Intn(9000) + 1000
	for i := 1; i < cnt; i++ {
		r.Exec(RandSelect)
	}
	println(cnt)
	r.Exec(PrintDecisionTree)
}
