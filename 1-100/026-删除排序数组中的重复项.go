/* 题目

https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/

给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

Example
给定数组 nums = [1,1,2], 
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 
你不需要考虑数组中超出新长度后面的元素。

Example
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。
*/

/* 思路
使用双指针，其中A指针正常遍历数组，B指针表示当前最后满足条件的下标

当数组为空或只有一个元素时，直接返回数组和数组的长度即可

当数组元素有2个或以上时，默认第一个为当前最后满足条件下标，故指针B为0
当指针A发现与B所指向元素不同的元素时，将该元素追加到满足条件元素最后，B指向下一个
当指针A发现与B所指向元素相同的元素时，将跳过该元素，B不动，A指向下一个

完成后，B指向最后一个正确元素的下标，所以长度为 B+1
*/

func removeDuplicates(nums []int) int {
    
    if len(nums) < 2 {
        return len(nums)
    }
    
    var cur int = 0
    
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[cur] {
            cur++
            nums[cur] = nums[i]
        }
    }
    
    return cur + 1
}