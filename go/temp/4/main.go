package main

import (
	"fmt"
)

type DataType int
type Node AVLTreeNode
type AVLTree *AVLTreeNode

type AVLTreeNode struct {
	key   DataType
	high  int
	left  *AVLTreeNode
	right *AVLTreeNode
}

func main() {
	data := []DataType{3, 2, 1, 4, 5, 6, 7, 16, 15, 14, 13, 12, 11, 10, 8, 9}
	fmt.Println(data)
	var root AVLTree = nil
	for _, value := range data {
		root = avl_insert(root, value)
		fmt.Println(" root -> key: ", root.key, "high :", root.high, "left", root.left, "right: ", root.right)
		//  preorder(root)

	}
	preorder(root)
	fmt.Println()
	midorder(root)
	fmt.Println()

}

func highTree(p AVLTree) int {
	if p == nil {
		return -1
	}
	return p.high
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*Please look LL*/
func left_left_rotation(k AVLTree) AVLTree {
	var kl AVLTree
	kl = k.left
	k.left = kl.right
	kl.right = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kl.high = max(highTree(kl.left), k.high) + 1
	return kl
}

/*Please look RR*/
func right_right_rotation(k AVLTree) AVLTree {
	var kr AVLTree
	kr = k.right
	k.right = kr.left
	kr.left = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kr.high = max(k.high, highTree(kr.right)) + 1
	return kr
}

/*so easy*/
func left_righ_rotation(k AVLTree) AVLTree {
	k.left = right_right_rotation(k.left)
	return left_left_rotation(k)
}

func right_left_rotation(k AVLTree) AVLTree {
	k.right = left_left_rotation(k.right)
	return right_right_rotation(k)
}

func avl_insert(avl AVLTree, key DataType) AVLTree {
	if avl == nil {
		avl = new(AVLTreeNode)
		if avl == nil {
			fmt.Println("avl tree create error!")
			return nil
		} else {
			avl.key = key
			avl.high = 0
			avl.left = nil
			avl.right = nil
		}

	} else if key < avl.key {
		avl.left = avl_insert(avl.left, key)
		if highTree(avl.left)-highTree(avl.right) == 2 {
			if key < avl.left.key { //LL
				avl = left_left_rotation(avl)
			} else { // LR
				avl = left_righ_rotation(avl)
			}
		}
	} else if key > avl.key {
		avl.right = avl_insert(avl.right, key)
		if (highTree(avl.right) - highTree(avl.left)) == 2 {
			if key < avl.right.key { // RL
				avl = right_left_rotation(avl)
			} else {
				fmt.Println("right right", key)
				avl = right_right_rotation(avl)
			}
		}
	} else if key == avl.key {
		fmt.Println("the key", key, "has existed!")

	}
	//notice: update high(may be this insert no rotation, so you should update high)
	avl.high = max(highTree(avl.left), highTree(avl.right)) + 1
	return avl
}

func preorder(avl AVLTree) {
	if avl != nil {
		fmt.Print(avl.key, "\t")
		preorder(avl.left)
		preorder(avl.right)
	}
}

func midorder(avl AVLTree) {
	if avl != nil {
		midorder(avl.left)
		fmt.Print(avl.key, "\t")
		midorder(avl.right)
	}
}

/*display avl tree*/
func display(avl AVLTree) {

}
