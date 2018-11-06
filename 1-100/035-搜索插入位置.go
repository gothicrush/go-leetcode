/* 题目

https://leetcode-cn.com/problems/search-insert-position/description/

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
你可以假设数组中无重复元素。

Example
输入: [1,3,5,6], 5
输出: 2

Example
输入: [1,3,5,6], 2
输出: 1

Example
输入: [1,3,5,6], 7
输出: 4

Example
输入: [1,3,5,6], 0
输出: 0

*/

/* 思路

使用二分搜索即可

*/

func searchInsert(nums []int, target int) int {
     
    var left int = 0
    var right int = len(nums) - 1
    
    for left <= right {
        mid := (right - left) / 2 + left
        
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    
    return left
}