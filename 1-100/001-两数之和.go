/* 题目
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

Example
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

/* 思路
数组：由下标索引到值
哈希表：由值索引到下标
target - nums[i] 为目标值
在哈希表中查看目标值的对应下标
返回两个下标i和j即可
*/

func twoSum(nums []int, target int) []int {
    
    var ret []int
    var m map[int]int = make(map[int]int, 10)
    
    for index, item := range nums {
        m[item] = index
    }
    
    for index, item := range nums {
        var value int = target - item
        
        if i, ok := m[value]; ok && i != index {
            ret = append(ret, index)
            ret = append(ret, i)
            break
        }
    }
    
    return ret
}