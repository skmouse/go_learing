package main

import (
	"fmt"
	"os"
)

type MyLinkedList struct {
	val  int
	next *MyLinkedList
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/**
 *  1.尾插法
 */
func (List *MyLinkedList) addAtTail(val int) {
	for List.next != nil {
		List = List.next
	}

	var node MyLinkedList
	node.val = val
	List.next = &node
}

/**
 * 2.头插法类似在两个节点之间插入一个结点的意思
 */
func (List *MyLinkedList) addAtHead(val int) {
	var node MyLinkedList
	node.val = val
	node.next = List.next
	List.next = &node
}

func (List *MyLinkedList) getAll() {
	for List.next != nil {
		fmt.Println(List.next)
		List = List.next
	}
}

/**
 * 3.查找某个结点是否存在
 */
func (List *MyLinkedList) get(index int) int {
	j := 0
	for List.next != nil && j < index {
		List = List.next
		j++
	}

	if j == index {
		return List.val
	} else {
		return -1
	}
}

/**
 * 4.在链表中的第 index 个节点之前添加值为 val  的节点。
 *   如果 index 等于链表的长度，则该节点将附加到链表的末尾。
 *   如果 index 大于链表长度，则不会插入节点。
 *   如果index小于0，则在头部插入节点。
 */
func (List *MyLinkedList) addAtIndex(index, val int) {
	var node MyLinkedList
	if index < 0 {
		//插入头结点
		node.next = List.next
		node.val = val
		List.next = &node
	} else if index < List.count() {
		for j := 0; j < index; j++ {
			List = List.next
		}
		node.next = List.next
		node.val = val
		List.next = &node
	} else if index == List.count() {
		for List.next != nil {
			List = List.next
		}
		node.val = val
		List.next = &node
	} else if index > List.count() {
		fmt.Println("不插入")
	}

}

/**
 * 5.delete 如果索引 index 有效，则删除链表中的第 index 个节点。
 */
func (List *MyLinkedList) deleteIndex(index int) {
	if index < 1 || index > List.count() {
		fmt.Println("索引无效")
		os.Exit(-1)
	}

	for i := 0; i < index-1; i++ {
		List = List.next
	}
	List.next = List.next.next
}

/**
 * 6.统计链表的结点
 */
func (List *MyLinkedList) count() int {
	i := 0
	for List.next != nil {
		List = List.next
		i++
	}
	return i
}

/**
 * 7.链表排序-- 冒泡
 */
func (List *MyLinkedList) SortList() *MyLinkedList {
	count := List.count()
	node := List.next

	//for count > 1 {
	//	for node.next!= nil {
	//		if node.val > node.next.val {
	//			node.val, node.next.val  = node.next.val, node.val
	//		}
	//
	//		node = node.next
	//	}
	//	count--
	//	node = List.next
	//}

	for i := 0; i < count-1; i++ {
		for j := 0; j < count-1-i; j++ {
			if node.val > node.next.val {
				node.val, node.next.val = node.next.val, node.val
			}
			node = node.next
		}
		node = List.next
	}
	return node
}

func mergeTwoLists(l1 *MyLinkedList, l2 *MyLinkedList) *MyLinkedList {
	//if l1 == nil {
	//	return l2
	//} else if l2 ==nil {
	//	return l1
	//} else if l1.val < l2.val {
	//	l1.next = mergeTwoLists(l1.next ,l2)
	//	return l1
	//} else {
	//	l2.next = mergeTwoLists(l1 ,l2.next)
	//	return l2
	//}

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *MyLinkedList
	if l1.val <= l2.val {
		res = l1
		res.next = mergeTwoLists(l1.next, l2)

		println(res.next)
	} else {
		res = l2
		res.next = mergeTwoLists(l1, l2.next)
		println(res.next)
	}
	return res
}

/**
  反转链表
*/
func reverList(Head *MyLinkedList) *MyLinkedList {
	if Head.next == nil {
		return Head
	}

	var res *MyLinkedList
	res = reverList(Head.next)
	return res
}

/**
 * 迭代反转链表
 */
func reverDdList(List *MyLinkedList) *MyLinkedList {
	if List == nil || List.next == nil {
		return List
	}

	var preNode, nextNode *MyLinkedList
	for List != nil {
		nextNode = List.next
		List.next = preNode
		preNode = List
		List = nextNode
	}
	return preNode
}

func middleNode(List *MyLinkedList) *MyLinkedList {
	tmp := [100]*MyLinkedList{}
	i := 0
	for List != nil {
		tmp[i] = List
		List = List.next
		i++
	}

	fmt.Print(i)
	return tmp[i/2]
}

func (List *MyLinkedList) test() *MyLinkedList {
	var tmp = MyLinkedList{-1, nil}
	tmp.next = List
	p := &tmp

	for p.next != nil {
		if p.next.val == 0 {
			p.next = p.next.next
			continue
		}
		p = p.next
	}
	return tmp.next
}

/**
 *判断是否是回文链表
 */
func (List *MyLinkedList) isPalindrome() bool {

	if List == nil && List.next == nil {
		return true
	}
	var tmp []int
	i := 0
	for List != nil {
		tmp = append(tmp, List.val)
		List = List.next
		i++
	}

	j := len(tmp)
	for i := 0; i <= j/2; i++ {
		if tmp[i] != tmp[j-1] {
			return false
		}

		j--
	}
	return true
}

func removeDuplicates(nums []int) {
	//如果是空切片，那就返回0
	if len(nums) == 0 {
		fmt.Println(-1)

	}
	//用两个标记来比较相邻位置的值
	//当一样的话，那就不管继续
	//当不一样的时候，就把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(nums); right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left] = nums[right]
	}
	//fmt.Println(nums)
	//fmt.Println(nums[:left+1])
}
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	//for k,v :=range nums {
	//	if v== val {
	//		nums = append(nums[:k],nums[k+1:]...)
	//	}
	//}

	fmt.Println(nums)

	return len(nums)
}

func (head *MyLinkedList) deleteDuplicates() *MyLinkedList {
	slow := head
	fast := head.next

	for fast != nil {
		if slow.val != fast.val {
			slow = slow.next
			slow.val = fast.val
		}
		fast = fast.next
	}

	return head

}

func create(List *MyLinkedList, data int) {
	for List != nil {
		List = List.next
	}
	var node MyLinkedList
	List = &node
	node.val = data
}

func get1(List *MyLinkedList) {
	for List.next != nil {
		fmt.Println(List.next)
		List = List.next
	}
}

func get2(List *MyLinkedList) {
	for List != nil {
		fmt.Println(List)
		List = List.next
	}
}

func (head *MyLinkedList) oddEvenList() {
	p := head
	q := p.next
	t := q

	for p != nil && q != nil && q.next != nil {
		p.next, q.next = p.next.next, q.next.next
		p, q = p.next, q.next
	}

	p.next = t

	for head != nil {
		fmt.Println(head.next)
		head = head.next
	}

	//slow,fast :=head.next,head.next
	//f,s := slow,fast
	//
	//for fast !=nil && fast.next !=nil {
	//	slow = slow.next
	//	fast = fast.next.next
	//}
	//os.Exit(1)
	//
	// slow.next = s
	//
	//for f.next !=nil  {
	//	fmt.Println(f)
	//	f = f.next
	//}
	// os.Exit(1)
}

func main() {

	obj1 := Constructor()
	arr1 := []int{1, 2, 3, 4}
	for _, v := range arr1 {
		obj1.addAtTail(v)
	}
	obj1.getAll()

	obj1.oddEvenList()

	//nums :=[]int{3,2,2,3}
	//removeElement(nums,3)

	////obj2 := Constructor()
	//arr1 := []int{1,2,3,3,4,4,5}
	////arr2 := []int{3}
	//for _, v := range arr1 {
	//	obj1.addAtTail(v)
	//}
	//
	//obj1.getAll()
	//
	//obj1.deleteDuplicates()
	//
	//fast.next != null && fast.val !==fast.next.val

	//a := [5]int {1,2,3,4,5}
	//obj1 := Constructor()
	//
	//for _,v :=range a {
	//	create(&obj1,v)
	//}

	//var a MyLinkedList
	//var b MyLinkedList
	//var c MyLinkedList
	//
	//var head *MyLinkedList
	//
	//a.val = 10
	//b.val =  20
	//c.val = 30
	//a.next = &b
	//b.next = &c
	//c.next = nil
	//
	//head = &a
	//
	////if head==nil || head.next ==nil {
	////	fmt.Println(-1)
	////}
	////
	////slow :=head
	////fast :=head.next
	////
	////for slow != fast  {
	////	if fast ==nil || fast.next ==nil {
	////		fmt.Println(-1)
	////		os.Exit(2)
	////	}
	////
	////	slow = slow.next
	////	fast = fast.next.next
	////}
	//
	//fmt.Println(head)

	//a := [3]int{1,1,2}
	//
	//for i:=0; i<len(a); i++  {
	//	for j:=i+1; j<len(a); j++  {
	//		if a[j] ==a[i] {
	//
	//		}
	//	}
	//}

	//
	//
	//a := [5]int{}
	//
	//fmt.Println(a)
	//
	//var b [5] int
	//
	//fmt.Println(b)
	//
	//
	// //c :=[5] int {2:1,3:2,4:3}
	//
	//fmt.Println([...] int {2:1,4:3})
	//
	//
	//slice :=make([]int, 6,10)
	//
	//
	//fmt.Println(slice, len(slice),cap(slice))

	//obj1 := Constructor()
	////obj2 := Constructor()

	//
	//obj1.getAll()

	//for _, v := range arr2 {
	//	obj2.addAtTail(v)
	//}
	//obj1.getAll()
	//fmt.Println(obj1.isPalindrome())

	//fmt.Println("---链表1-------")
	//obj1.getAll()
	//fmt.Println("---链表2-------")
	//obj2.getAll()
	////
	////
	////fmt.Println("---链表反转-------")
	////
	////a :=reverDdList(&obj1)
	////a.getAll()
	//
	//
	//
	////middleNode(&obj1)
	//
	//for _, v := range arr {
	//	obj.addAtTail(v)
	//}
	////fmt.Println(obj.get(5))
	////fmt.Println(obj.get(10))
	//fmt.Println(obj.count()) //统计结点
	//obj.deleteIndex(5)
	//obj.getAll()
	//
	//obj.addAtIndex(-1,-1)
	//obj.addAtIndex(20,-1)
	//obj.getAll()
	//
	//
	//obj.addAtHead(1)
	//obj.addAtTail(3)
	//obj.addAtIndex(1,2)
	//
	//obj.getAll()
	//fmt.Println(obj.get(1))
	//obj.deleteIndex(1)
	//obj.getAll()
	//fmt.Println("---链表排序-------")
	//obj.SortList()

}
