package main

// import (
// 	"fmt"
// )

// type node struct {
// 	data int
// 	next *node
// }

// type linkedList struct {
// 	head   *node
// 	length int
// }

// func (l *linkedList) prepand(n *node) {
// 	//1st approach
// 	// second := l.head
// 	// l.head = n
// 	// l.head.next = second
// 	// l.length++

// 	head := l.head
// 	n.next = head
// 	l.head = n
// 	l.length++
// }

// func (l *linkedList) addAtLast(n *node) {
// 	head := l.head
// 	if head == nil {
// 		l.head = n
// 		n.next = nil
// 		l.length++
// 		return
// 	}
// 	for head.next != nil {
// 		head = head.next
// 	}
// 	head.next = n
// 	n.next = nil
// 	l.length++
// }

// func (l linkedList) printListData() {
// 	head := l.head
// 	// 1st approach
// 	if head == nil {
// 		fmt.Printf("list is empty=%v\n", l)
// 		return
// 	}

// 	for head.next != nil {
// 		fmt.Printf("%d ", head.data)
// 		head = head.next
// 	}

// 	fmt.Printf("%d ", head.data)
// 	fmt.Println()

// 	//2nd approach
// 	// for l.length > 0 {
// 	// 	fmt.Printf("%d ", head.data)
// 	// 	head = head.next
// 	// 	l.length--
// 	// }
// 	// fmt.Println()

// }

// func (l *linkedList) deleteNode(val int) {
// 	if l.length == 0 {
// 		return
// 	}

// 	tmphead := l.head
// 	//if val matches at head
// 	if tmphead.data == val {
// 		l.head = tmphead.next
// 		l.length--
// 		return
// 	}

// 	for tmphead.next.data != val {
// 		//it reaches out to the end of the list
// 		if tmphead.next.next == nil {
// 			return
// 		}
// 		tmphead = tmphead.next
// 	}
// 	tmphead.next = tmphead.next.next
// 	l.length--

// }

// func (l *linkedList) addNode(pos, val int) {
// 	newNode := &node{data: val, next: nil}
// 	prevHead := l.head
// 	//if pos found at head
// 	if prevHead.data == pos {
// 		newNode.next = prevHead.next
// 		prevHead.next = newNode
// 		l.length++
// 		return
// 	}

// 	for prevHead.next.data != pos {
// 		//data not found
// 		if prevHead.next.next == nil {
// 			return
// 		}
// 		prevHead = prevHead.next
// 	}
// 	newNode.next = prevHead.next.next
// 	prevHead.next.next = newNode
// 	l.length++
// }

// func main_1() {
// 	list := linkedList{}
// 	node1 := &node{data: 40}
// 	node2 := &node{data: 16}
// 	node3 := &node{data: 21}
// 	node4 := &node{data: 1}
// 	node5 := &node{data: 11}
// 	node6 := &node{data: 31}
// 	list.prepand(node1)
// 	list.prepand(node2)
// 	list.prepand(node3)
// 	list.prepand(node4)
// 	list.prepand(node5)
// 	list.prepand(node6)
// 	fmt.Println("Prepand methods List")
// 	list.printListData()
// 	// list.addAtLast(node1)
// 	// list.addAtLast(node2)
// 	// list.addAtLast(node3)
// 	// list.addAtLast(node4)
// 	// list.addAtLast(node5)
// 	// list.addAtLast(node6)
// 	// fmt.Println("addAtLast methods List")
// 	// list.printListData()
// 	// list.deleteNode(40)
// 	// fmt.Println("List after deleting 40")
// 	// list.printListData()
// 	list.addNode(1, 2)
// 	fmt.Println("After Addition List")
// 	list.printListData()

// 	// emptylist := linkedList{}
// 	// fmt.Println("Empty List")
// 	// emptylist.printListData()

// }
