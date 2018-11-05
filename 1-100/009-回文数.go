/* 题目
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

Example
输入: 121
输出: true

Example
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

Example
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/

/*
如果一个数小于零，则一定false
如果一个数大于0，但是它模10为0，则一定false

通过取个位数，对整个数进行反转
判断反转后的数与原始数是否相等
*/

func isPalindrome(x int) bool {
    
    if x < 0 || (x != 0 && x % 10 == 0) {
        return false
    }
    
    var src int = x
    var ret int = 0
    
    for x > 0 {
        ret = ret * 10 + x % 10
        x /= 10
    }
    
    return ret == src
}