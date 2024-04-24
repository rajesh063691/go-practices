package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	head   *node
	length int
}

func findDecimal(head *node) int {
	cur := head
	decNum := make([]int, 1)
	for cur != nil {
		decNum = append(decNum, cur.data)
		cur = cur.next
	}
	return 0
}

func main() {
	var node1 *node
	if node1 == nil {
		fmt.Printf("Node is Empty, Node=%v \n", node1)
	} else {
		{
			fmt.Printf("Node is not Empty, Node=%v \n", node1)
		}
	}
}
