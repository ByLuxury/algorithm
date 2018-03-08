package main

import (
	"fmt"
)

type List struct {
	data     int
	nextNode *List
}

var firstNode *List

func insertNode(data int) {
	dataList := &List{
		data: data,
	}
	dataList.nextNode = firstNode
	firstNode = dataList
}

func findNode(data int) int {

	currNode := firstNode
	for currNode != nil {
		if currNode.data == data {
			return currNode.data
		}
		currNode = currNode.nextNode
	}
	return -1

}
func isEmpty() bool {
	return firstNode == nil
}
func display() {
	if firstNode == nil {
		fmt.Println("Node is nil")
	}
	currNode := firstNode
	for currNode != nil {
		fmt.Print(currNode.data, "-")
		currNode = currNode.nextNode
	}
	fmt.Println()
}
func remove(data int) {
	if firstNode == nil {
		return
	}
	if firstNode.data == data {
		firstNode = firstNode.nextNode
	} else {
		preNode := firstNode
		currNode := firstNode.nextNode
		for currNode != nil {
			if currNode.data == data {
				preNode.nextNode = currNode.nextNode
			}
			preNode = currNode
			currNode = currNode.nextNode
		}
	}
}
func main() {

	//insertNode(1)
	//insertNode(2)
	//insertNode(3)
	//insertNode(4)
	//insertNode(5)
	for i := 10; i > 0; i-- {
		insertNode(i)
	}
	display()
	fmt.Println(findNode(3))
	remove(3)
	fmt.Println(findNode(3))
	display()

}
