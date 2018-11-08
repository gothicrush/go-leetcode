/* 题目

https://leetcode-cn.com/problems/plus-one/description/

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头。

Example
输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。

Example
输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

*/

/* 思路

情况一： 加了不进位，如 88 + 1 = 89
情况二： 加了进位，如 89 + 1 = 90
情况三： 全9然后加1，如 99 + 1 = 100

*/

func plusOne(digits []int) []int {
    
    if len(digits) <= 0 {
        return digits
    }
    
    for i := len(digits) - 1; i >= 0 ; i-- {
        if digits[i] != 9 {
            digits[i]++
            return digits
        } else {
            digits[i] = 0
        }
    }
    
    var ret []int = make([]int,len(digits)+1)
    
    ret[0] = 1
    
    return ret
    
}