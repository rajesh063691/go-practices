package linkedlist

// import "fmt"

// type Node struct {
// 	data int
// 	next *Node
// }

// var head *Node = nil

// func (list *Node) InsertAtLast(item int) *Node {
// 	if list == nil {
// 		list.data = item
// 		list.next = nil
// 		//list = list
// 		return list
// 	} else {
// 		for list.next != nil {
// 			//this code is to prevent duplicate insertion
// 			if list.data == item {
// 				return list
// 			}
// 			list = list.next
// 		}
// 		newNode := new(Node)
// 		newNode.data = item
// 		newNode.next = nil
// 		list.next = newNode
// 		return list
// 	}
// }

// func (list *Node) DisplayList() {
// 	if list == nil {
// 		fmt.Println("List is Empty")
// 	}
// 	for list.next != nil {

// 		fmt.Print(list.data)
// 		fmt.Print("=>")
// 		list = list.next
// 	}
// 	fmt.Print(list.data)
// 	fmt.Println()
// }
