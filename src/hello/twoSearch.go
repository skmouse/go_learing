package main

import (
	"fmt"
	"os"
)

/**
 * 二叉搜索树
 */
type TreeNode struct {
	data int
	Left, Right *TreeNode
}

func CreateNode(data int) *TreeNode  {
	return &TreeNode{data: data}
}

/**
 * 二叉搜索树
 */
func CreateTwoSearchTree(root *TreeNode, data int) *TreeNode  {
	if root==nil {
		root =  &TreeNode{data:data}
	}

	if data > root.data {
		root.Right =   CreateTwoSearchTree(root.Right, data)
	} else if data < root.data {
		root.Left =  CreateTwoSearchTree(root.Left, data)
	}

	return  root
}

func SearchTree(root *TreeNode, data int) int {
	if root ==nil {
		return  -1
	}

	if data > root.data {
		root = root.Right
	}else if data < root.data {
		root = root.Left
	}

	return  root.data
}

func Max(root *TreeNode) *TreeNode  {
	if root ==nil {
		return nil
	} else if root.Right ==nil {
		return  root
	} else {
		return Max(root.Right)
	}
}

func Min(root *TreeNode) *TreeNode  {
	if root ==nil {
		return nil
	} else if root.Left ==nil {
		return root
	} else {
		return  Min(root.Left)
	}
}


func ForMin(root *TreeNode) *TreeNode  {
	if root !=nil {
		for root.Left !=nil {
			root = root.Left
		}
	}
	return root
}

func DeleteTree(root *TreeNode, data int) *TreeNode  {
	if root ==nil {
		fmt.Println("未找到该元素")
		os.Exit(-1)
	}

	if data < root.data {
		root.Left = DeleteTree(root.Left, data)
	} else if data > root.data {
		root.Right = DeleteTree(root.Right, data)
	}else {
		if root.Left!= nil && root.Right!=nil  {
			// 右字数最小值 ||左子树最大值
		} else {
			tmp := Min(root.Right)
			root.data = tmp.data
			root.Right = DeleteTree(root.Right, root.data)
			if root.Right ==nil {
				root = root.Left
			} else if root.Left ==nil {
				root = root.Right
			} else {
				root = nil
			}
		}
	}
	return  root
}



func PreOrder(root *TreeNode) {
	if root ==nil {
		return
	}
	fmt.Print(root.data, " ")
	PreOrder(root.Left)
	PreOrder(root.Right)
}
func main()  {
	BST := CreateNode(50)
	CreateTwoSearchTree(BST,40)
	CreateTwoSearchTree(BST,60)
	CreateTwoSearchTree(BST,37)
	CreateTwoSearchTree(BST,41)
	CreateTwoSearchTree(BST,58)
	PreOrder(BST)


	fmt.Println("\n"+"二叉搜索树查找")
	fmt.Println(SearchTree(BST,50))
	//
	//fmt.Println("\n"+"二叉搜索树删除")
	//fmt.Println(DeleteTree(BST,60))
	//
	//PreOrder(BST)

	fmt.Println("\n"+"二叉搜索树最小值")



	fmt.Println(Min(BST))
	fmt.Println(ForMin(BST))

	fmt.Println(Max(BST))

}