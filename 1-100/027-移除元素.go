/* 题目

https://leetcode-cn.com/problems/remove-element/description/

给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

Example
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

Example
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
注意这五个元素可为任意顺序。
你不需要考虑数组中超出新长度后面的元素。

*/

/* 思路

用双指针解决，A指针正常遍历数组，B指针表示当前最后一个正确元素下标

当A指针遇到非指定元素时，追加到正确元素后，B指针移动

*/

func removeElement(nums []int, val int) int {
    
    var cur int = 0
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            nums[cur] = nums[i]
            cur++
        }
    }
    
    return cur
}