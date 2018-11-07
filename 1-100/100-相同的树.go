/* 题目

https://leetcode-cn.com/problems/same-tree/description/

给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

*/

/* 思路

如果都为nil，则为true
如果一个为nil，另外一个不为nil，则返回false
如果两个都不为nil，但value不同，则返回false
如果两个都不为nil，但value相同，则递归判断左子树和右子树

*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    
    if p == nil && q == nil {
        return true
    }
    
    if p == nil || q == nil {
        return false
    }
    
    if p.Val != q.Val {
        return false
    }
    
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}