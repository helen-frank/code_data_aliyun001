package main

import "fmt"

type AvlTreeNode struct {
	Height int
	Value  int
	Left   *AvlTreeNode
	Right  *AvlTreeNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (avl *AvlTreeNode) GetHeight() int {
	if avl == nil {
		return 0
	}
	return avl.Height
}

// 查找最大节点
func (avl *AvlTreeNode) FindMax() *AvlTreeNode {
	if avl == nil {
		return nil
	} else if avl.Right != nil {
		return avl.Right.FindMax()
	} else {
		return avl
	}
}

// 查找最小节点
func (avl *AvlTreeNode) FindMin() *AvlTreeNode {
	if avl == nil {
		return nil
	} else if avl.Left != nil {
		return avl.Left.FindMin()
	} else {
		return nil
	}
}

// 获取节点左子树高度差绝对值
func (avl *AvlTreeNode) GetBF() int {
	var lh, rh int
	if avl.Left != nil {
		lh = avl.Left.GetHeight()
	}
	if avl.Right != nil {
		rh = avl.Right.GetHeight()
	}

	bf := lh - rh
	if bf < 0 {
		bf = bf * -1
	}
	return bf
}

// 左旋
// avl为最小失衡子树的根节点
// 问题:在右子树插入右孩子导致AVL树失衡
// return 新的平衡树的根节点

func (avl *AvlTreeNode) LeftRotation() *AvlTreeNode {
	tmpNode := avl.Right
	avl.Right = tmpNode.Left
	tmpNode.Left = avl
	avl.Height = Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1
	tmpNode.Height = Max(tmpNode.Left.GetHeight(), tmpNode.Right.GetHeight()) + 1
	return tmpNode
}

// 右旋
// avl为最小失衡树的根节点
// 问题:在左子树插入左孩子导致AVL树失衡
// return 新的平衡树根节点

func (avl *AvlTreeNode) RightRotation() *AvlTreeNode {
	tmpNode := avl.Left
	avl.Left = tmpNode.Right
	tmpNode.Right = avl
	avl.Height = Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1
	tmpNode.Height = Max(tmpNode.Left.GetHeight(), tmpNode.Right.GetHeight()) + 1
	return tmpNode
}

// 右左双旋转
// 问题：通常因为在右子树上插入左孩子导致AVL失衡
// 解法：先右旋后左旋调整
// return 新的平衡树根节点

func (avl *AvlTreeNode) RightLeftRotation() *AvlTreeNode {
	avl.Right = avl.Right.RightRotation()
	return avl.Left.LeftRotation()
}

// 左右双旋转
// 问题：通常因为在左子树上插入右孩子导致AVL失衡
// 解法：先左旋后右旋调整
// return 新的平衡树根节点

func (avl *AvlTreeNode) LeftRightRotation() *AvlTreeNode {
	avl.Left = avl.Left.LeftRotation()
	return avl.Right.RightRotation()
}

// avl 插入

func (avl *AvlTreeNode) InsertAvl(v int) *AvlTreeNode {
	if avl == nil { // 空树
		avl := new(AvlTreeNode)
		avl.Value = v
	} else if v < avl.Value { // 待插入的值比单前节点值小，往左子树插入
		avl.Left = avl.InsertAvl(v)
		if avl.Left.GetHeight()-avl.Right.GetHeight() == 2 { //插入新节点后avl树失衡
			if v < acl.Left.Value {
				avl = avl.RightRotation(tree)
			} else if v > avl.Left.Value {
				avl = avl.LeftRightRotation()
			}
		}
	} else if v > avl.Value {
		avl.Right = avl.Right.InsertAvl(v)
		if avl.Right.GetHeight()-avl.Left.GetHeight() == 2 { //在右子树插入新节点后avl树失衡
			if v > avl.Right.Value {
				avl = avl.LeftRotation()
			} else if v < avl.Right.Value {
				avl = avl.RightLeftRotation()
			}
		}
	}
	avl.Height = Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1
	return avl
}

func main() {
	fmt.Println("vim-go")
}
