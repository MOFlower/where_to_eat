package w2e

import (
	"fmt"
	"math/rand"
)

// convert simpleNode to decisionNode
func GenProbabilityFunc(n Node) {
	sum := 0
	node := getSimpleNode(n)
	for index, item := range node.Children {
		switch item.(type) {
		case *SimpleNode:
			sum += avgLevel
			node.Children[index] = &DecisionNode{
				*(item.(*SimpleNode)),
				avgLevel,
				sum,
				0,
			}
		case *DecisionNode:
			sum += item.(*DecisionNode).level
			item.(*DecisionNode).cumulativeProbability += sum
		}

		item.Exec(GenProbabilityFunc)
	}
}

func RandSelect(n Node) {
	node := getSimpleNode(n)
	print(node.Name)
	if len(node.Children) <= 0 {
		print("\n")
		return
	}
	print(" -> ")

	sum := node.Children[len(node.Children)-1].(*DecisionNode).cumulativeProbability
	r := int(rand.Int31n(int32(sum)))
	for _, item := range node.Children {
		if item.(*DecisionNode).cumulativeProbability > r {
			item.(*DecisionNode).count++
			item.Exec(RandSelect)
			break
		}
	}
}

func PrintDecisionTree(n Node) {
	sum := 0
	sum2 := 0
	for _, item := range getSimpleNode(n).Children {
		sum += item.(*DecisionNode).count
		sum2 += item.(*DecisionNode).level
		// fmt.Printf("(%s, %d, %d, %d) ",
		// 	item.(*DecisionNode).name,
		// 	item.(*DecisionNode).level,
		// 	item.(*DecisionNode).cumulativeProbability,
		// 	item.(*DecisionNode).count)
	}
	for _, item := range getSimpleNode(n).Children {
		i := item.(*DecisionNode)
		fmt.Printf("name := %s, target := %.2f%%, now := %.2f%%, diff := %.2f%%\n",
			i.Name,
			float64(i.level)/float64(sum2)*100,
			float64(i.count)/float64(sum)*100,
			(float64(i.level)/float64(sum2)-float64(i.count)/float64(sum))*100,
		)
	}
	if len(getSimpleNode(n).Children) > 0 {
		println("-----------------------------------" +
			"---------------------------------------")
	}
	for _, item := range getSimpleNode(n).Children {
		item.Exec(PrintDecisionTree)
	}
}
