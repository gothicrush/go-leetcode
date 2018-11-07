/* 题目

https://leetcode-cn.com/problems/sort-colors/description/

给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

Example
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]

*/

/* 思路

用left表示最后一个零下一个位置
用right表示第一个二的前一个位置
用cur进行遍历
遇到0，则与left位置进行交换
遇到1，则前进
遇到2，则与right位置进行交换

*/

func sortColors(nums []int)  {
    
    var left int = 0
    var right int = len(nums) - 1
    var cur int = 0
    
    for cur <= right {
        if nums[cur] == 0 {
            swap(nums, cur, left)
            cur++
            left++
        } else if nums[cur] == 1 {
            cur++
        } else {
            swap(nums, cur, right)
            right--
        }
    }
}

func swap(nums []int, i int, j int) {
    temp := nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}