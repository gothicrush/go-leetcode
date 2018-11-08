/* 题目

https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/

给定一个二叉树，返回它的中序 遍历。

Example
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

*/

/* 思路

用递归法直接解决

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    
    var ret []int
    
    inorder(root, &ret)
    
    return ret
}

func inorder(node *TreeNode, ret *[]int) {
    
    if node == nil {
        return
    }
    
    inorder(node.Left, ret)
    *ret = append(*ret, node.Val)
    inorder(node.Right, ret)
}