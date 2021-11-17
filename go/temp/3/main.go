package main

import (
	"fmt"
)

type Type = int

type AvlTreeNode struct {
	Key    Type
	Height int
	Left   *AvlTreeNode
	Right  *AvlTreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *AvlTreeNode) GetHeight() int {
	if a == nil {
		return 0
	}
	return a.Height
}

func (a *AvlTreeNode) LeftRotation() *AvlTreeNode {
	tmpNode := a.Right
	a.Right = tmpNode.Left
	tmpNode.Left = a
	tmpNode.Height = Max(tmpNode.Left.GetHeight(), tmpNode.Right.GetHeight()) + 1
	a.Height = Max(a.Left.GetHeight(), a.Right.GetHeight()) + 1
	return tmpNode
}

func (a *AvlTreeNode) RightRotation() *AvlTreeNode {
	tmpNode := a.Left
	a.Left = tmpNode.Right
	tmpNode.Right = a
	tmpNode.Height = Max(tmpNode.Left.GetHeight(), tmpNode.Right.GetHeight()) + 1
	a.Height = Max(a.Left.GetHeight(), a.Right.GetHeight()) + 1
	return tmpNode
}

func (a *AvlTreeNode) LeftRightRotation() *AvlTreeNode {
	a.Left = a.Left.LeftRotation()
	return a.RightRotation()
}

func (a *AvlTreeNode) RightLeftRotation() *AvlTreeNode {
	a.Right = a.Right.RightRotation()
	return a.LeftRotation()
}

func (a *AvlTreeNode) InsertAvl(v Type) *AvlTreeNode {
	if a == nil {
		a = new(AvlTreeNode)
		if a == nil {
			fmt.Println("error")
			return nil
		}
		a.Key = v
		return a
	} else if v > a.Key {
		a.Right = a.Right.InsertAvl(v)
		if a.Right.GetHeight()-a.Left.GetHeight() == 2 {
			if v > a.Right.Key {
				a = a.LeftRotation()
			} else if v < a.Right.Key {
				a = a.RightLeftRotation()
			}
		}
	} else if v < a.Key {
		a.Left = a.Left.InsertAvl(v)
		if a.Left.GetHeight()-a.Right.GetHeight() == 2 {
			if v < a.Right.Key {
				a = a.RightRotation()
			} else if v > a.Left.Key {
				a = a.LeftRightRotation()
			}
		}
	}
	a.Height = Max(a.Left.GetHeight(), a.Right.GetHeight()) + 1
	return a
}

func (a *AvlTreeNode) PreOrder() {
	if a != nil {
		fmt.Print(a.Key, " ")
		a.Left.PreOrder()
		a.Right.PreOrder()
	}
}

func main() {
	var a *AvlTreeNode
	b := []Type{2, 1, 4, 3, 5, 6, 7, 8}
	for i := range b {
		fmt.Printf("%d ", b[i])
		a = a.InsertAvl(b[i])
	}
	fmt.Println()
	a.PreOrder()
	fmt.Println()
}
