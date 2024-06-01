package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	output := isSymmetricSolution2(input)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetricSolution1(root *TreeNode) bool {
	var leftFn func(*TreeNode, int, string) (int, []string)
	leftFn = func(node *TreeNode, depth int, flag string) (int, []string) {
		if node == nil {
			return depth - 1, nil
		}
		leftDepth, leftSlice := leftFn(node.Left, depth+1, "L")
		slice := append([]string{strconv.Itoa(node.Val) + flag}, leftSlice...)
		rightDepth, rightSlice := leftFn(node.Right, depth+1, "R")
		slice = append(slice, rightSlice...)
		return max(leftDepth, rightDepth), slice
	}
	var rightFn func(*TreeNode, int, string) (int, []string)
	rightFn = func(node *TreeNode, depth int, flag string) (int, []string) {
		if node == nil {
			return depth - 1, nil
		}
		rightDepth, rightSlice := rightFn(node.Right, depth+1, "L")
		slice := append([]string{strconv.Itoa(node.Val) + flag}, rightSlice...)
		leftDepth, leftSlice := rightFn(node.Left, depth+1, "R")
		slice = append(slice, leftSlice...)
		return max(rightDepth, leftDepth), slice
	}
	leftDepth, leftSlice := leftFn(root, 1, "P")
	rightDepth, rightSlice := rightFn(root, 1, "P")
	if leftDepth != rightDepth || !isSameSlice(leftSlice, rightSlice) {
		return false
	}
	return true
}

func isSameSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func isSymmetricSolution2(root *TreeNode) bool {
	var recursiveFn func(*TreeNode, *TreeNode) bool
	recursiveFn = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val &&
			recursiveFn(left.Right, right.Left) &&
			recursiveFn(left.Left, right.Right)
	}
	if root != nil {
		return recursiveFn(root.Left, root.Right)
	}
	return true
}
