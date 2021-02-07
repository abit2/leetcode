package main

/*
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import (
	"math"
	"strconv"
	"strings"
)

type Codec struct {
	val []string
}

func Constructor() Codec {
	return Codec{val: []string{}}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	this.treeToString(root)
	return strings.Join(this.val, ",")
}

func (this *Codec) treeToString(root *TreeNode) {
	if root == nil {
		return
	}

	this.val = append(this.val, strconv.Itoa(root.Val))
	this.treeToString(root.Left)
	this.treeToString(root.Right)
}

func (this *Codec) stringToTree(maxV int, minV int, index *int) *TreeNode {
	root := TreeNode{Val: -1}
	if *index < len(this.val) {
		v, _ := strconv.Atoi(this.val[*index])
		if minV < v && v < maxV {
			root.Val = v
			*index += 1
			root.Left = this.stringToTree(v, minV, index)
			root.Right = this.stringToTree(maxV, v, index)
		}
	}

	if root.Val == -1 {
		return nil
	}
	return &root
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	this.val = strings.Split(data, ",")
	index := 0
	return this.stringToTree(math.MaxInt64, math.MinInt64, &index)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
