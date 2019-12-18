package main



/**
 * 1.二叉树的遍历，查找，curd
 */
import (
	"fmt"
)

type Node struct {
	Data        int
	Right, left *Node
}

func (node *Node) setVaule(val int) {
	if node == nil {
		fmt.Println("setting value to nil.node ignored.")
		return
	}

	node.Data = val
}

func (node *Node) Print() {
	fmt.Print(node.Data, " ")
}

func createNode(val int) *Node {
	return &Node{Data: val}
}

/**
 * 前序遍历递归
 */
func (node *Node) PreOrder() {
	if node == nil {
		return
	}

	node.Print()
	node.left.PreOrder()
	node.Right.PreOrder()
}

/**
 * 中序遍历递归
 */
func (node *Node) MidOrder() {
	if node == nil {
		return
	}

	node.left.MidOrder()
	node.Print()
	node.Right.MidOrder()
}

/**
 * 后续遍历递归
 */
func (node *Node) PostOrder() {
	if node == nil {
		return
	}

	node.left.PostOrder()
	node.Right.PostOrder()
	node.Print()
}

/**
 * 前序遍历非递归
 */
func (node *Node) SearchPreOrder() {
	stack := []*Node{}
	for node !=nil || len(stack) > 0 {
		for node !=nil  {
			fmt.Print(node.Data, " ")
			stack = append(stack, node)
			node = node.left
		}

		if len(stack) > 0 {
			move := stack[len(stack) -1]
			stack = stack[:len(stack)-1]
			node = move.Right
		}
	}
}

func (node *Node)SearchMidOrder()  {
	result := []*Node{}
	p :=node
	for p !=nil || len(result) >0 {
		for p !=nil {
			result = append(result,p)
			p = p.left
		}

		if len(result) >0 {
			move := result[len(result)-1] //出栈的元素
			fmt.Print(move.Data, " ")
			result = result[:len(result)-1]
			p = move.Right
		}
	}
}

/**
 * 广度优先
 * 队列实现 根节点入队，出队，判断左右结点入队
 */

func (node *Node) BFS () {
	queue := []*Node{node}

	data :=[]int{}
	for len(queue) >0 {
		result := queue[0]
		data =append(data, result.Data)
		queue = queue[1:]
		if result.left !=nil {
			queue = append(queue, result.left)
		}

		if result.Right !=nil {
			queue =append(queue, result.Right)
		}
	}

	for _,v :=range data {
		fmt.Print(v, " ")
	}
}

/**
 * 递归树的高度
 */
func (node *Node) TreeHeight() int {
	if node ==nil {
		return 0
	}

	leftHeight := node.left.TreeHeight()
	rightHeight := node.Right.TreeHeight()

	if leftHeight > rightHeight {
		 return  leftHeight +1
	} else {
		return  rightHeight +1
	}
}

/**
 * 非递归 求二叉树的高度
 */
func (node *Node) TreeHeight2() int  {
	queue := []*Node{node}
	lays :=0
	for len(queue) >0 {
		lays ++
		size := len(queue)
		for count:=0; count < size; count ++  {
			result := queue[0]
			queue = queue[1:]
			if result.left !=nil {
				queue = append(queue, result.left)
			}

			if result.Right !=nil {
				queue =append(queue, result.Right)
			}
		}
	}

	return lays
}

func (node *Node) BST(Val int) int {
	for node !=nil {
		if Val < node.Data {
			node = node.left
		} else if Val > node.Data {
			node = node.Right
		}
		return node.Data
	}
	return  -1
}

/**
 * 二叉树的最大值
 */
func (node *Node) MaxBST() int {
	if node ==nil {
		return  -1
	}

	for node.Right !=nil  {
		node = node.Right
	}
	return node.Data
}

/**
 * 二叉树的最小值
 */
func (node *Node) MinBST() int  {
	if node ==nil {
		return  -1
	}

	for node.left !=nil {
		node = node.left
	}
	return  node.Data
}

/**
 * 二叉搜索树的插入
 */
func (node *Node) InsertBST(X int) *Node  {
	if node ==nil {
		node = createNode(X)
	} else {
		if X < node.Data {
			node.left = node.left.InsertBST(X)
		} else if X > node.Data {
			node.Right =node.Right.InsertBST(X)
		}
	}

	return node
}


/**
 * 二叉树搜素树的删除
 */
func (node *Node) deleteBST()  {

}




func main() {
	root := createNode(10)
	root.left = createNode(7)
	root.left.left = createNode(4)
	root.left.Right = createNode(8)
	root.left.Right.left = createNode(6)

	root.Right = createNode(11)
	root.Right.left = createNode(9)
	root.Right.Right = createNode(12)


	fmt.Print("\n前序遍历: ")
	root.PreOrder()

	fmt.Print("\n中序遍历: ")
	root.MidOrder()

	fmt.Print("\n后序遍历: ")
	root.PostOrder()

	fmt.Print("\n广度优先: ")
	root.BFS()

	fmt.Print("\n前序非递归遍历: ")
	root.SearchPreOrder()

	fmt.Print("\n中序非递归遍历: ")
	root.SearchMidOrder()

	fmt.Print("\n二叉树的高度: ")
	fmt.Print(root.TreeHeight())

	fmt.Print("\n二叉树的高度非递归: ")
	fmt.Print(root.TreeHeight2())

	fmt.Print("\n二叉树的最大值: ")
	fmt.Print(root.MaxBST())

	fmt.Print("\n二叉树的最小值: ")
	fmt.Print(root.MinBST())


	fmt.Print("\n二叉树搜索树的插入: ")
	root.InsertBST(1)
	root.InsertBST(13)
	root.PreOrder()









	//
	//
	//
	//a :=[]int{1,2,3,4}
	//
	//fmt.Println(a[:len(a)-1])



}
