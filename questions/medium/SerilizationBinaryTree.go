// 449. 序列化和反序列化二叉搜索树
// 2020.6.30 9:10 by kencin

package medium

import (
	"sort"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	// 1. 把二叉搜索树按大小排布，然后字符串存储。 问题：可能在反序列化时位置不一样
	// 2. 图形化二叉搜索树来存储，及第一排是什么，第二排是什么，第三排是什么。。。
	// 决定采用第二种方式

	if root == nil {
		return ""
	}
	//
	//// 定义一个切片，用来存储每一排的数据
	//// 切片为int类型
	//var allNode [][] *TreeNode
	//
	//allNode = make([][]*TreeNode, 100)
	//
	//allNode[0][0] = root
	//
	//curNode := root
	//curIndex := 0
	//
	//if curNode.Left != nil{
	//	allNode[0][0] = root
	//}

	// 上面的统统不要，想叉了，按照上面的做法有点复杂
	// 根据前序遍历和中序遍历即可构建出一棵树，而对于二叉搜索树来说，中序遍历即是排序结果，因此我们只需要序列化前序遍历即可
	// 前序遍历可以采用递归也可以采用栈存储的方式
	// 这里采用栈存储的方式吧

	// 以下是前序遍历
	var stack []*TreeNode

	var valueList []string

	curNode := root

	for curNode != nil || len(stack) != 0 {
		if curNode != nil {
			valueList = append(valueList, strconv.Itoa(curNode.Val))
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			curNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curNode = curNode.Right
		}
	}
	return strings.Join(valueList, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	stringValueList := strings.Split(data, ",")

	var intValueList []int

	for _, i := range stringValueList {
		theInt, _ := strconv.Atoi(i)
		intValueList = append(intValueList, theInt)
	}

	pre_o := make([]int, len(intValueList))
	copy(pre_o, intValueList)

	// 排序intValueList使之成为中序遍历
	sort.Ints(intValueList)

	return buildTree(pre_o, intValueList)

}

func buildTree(pre_o []int, in_o []int) *TreeNode {
	if pre_o == nil {
		return nil
	}
	mid := pre_o[0]
	midIndexOfInO := indexOf(mid, in_o)
	if midIndexOfInO == -1 {
		panic("mid index of in_o not found")
	}
	root := &TreeNode{Val: mid, Left: nil, Right: nil}
	root.Left = buildTree(pre_o[1:midIndexOfInO+1], in_o[:midIndexOfInO])
	root.Right = buildTree(pre_o[midIndexOfInO+1:], in_o[midIndexOfInO+1:])
	return root
}

func indexOf(value int, list []int) int {
	for k, v := range list {
		if value == v {
			return k
		}
	}
	return -1
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
