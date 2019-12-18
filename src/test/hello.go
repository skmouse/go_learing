package main

/**
 * 单链表的curd
 */

type Node struct {
	Data int
	Next *Node
	Pre *Node
}


/**
 * 双向链表尾插法的创建
 */

func endCreate(head *Node,data int,i int)  {
	point :=head
	for point.Next !=nil {
		point = point.Next
	}

	var node Node
	node.Data = data
	point.Next = &node

}



func main() {
	var head Node
	head.Data = -1
	head.Next = nil
	head.Pre = nil
	//for i :=0; i<10; i++ {
	//	headCreate(&head,i)
	//}
	//get(&head)

	for i :=0; i<10; i++ {
		endCreate(&head,i,i)
	}
}
