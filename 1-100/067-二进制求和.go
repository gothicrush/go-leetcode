/* 题目

https://leetcode-cn.com/problems/add-binary/description/

给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。

Example
输入: a = "11", b = "1"
输出: "100"

Example
输入: a = "1010", b = "1011"
输出: "10101"

*/

/* 思路

从右向左求和即可，注意最后可能会有进位

*/

func addBinary(a string, b string) string {
    
    var i int = len(a) - 1
    var j int = len(b) - 1
    
    var ret []byte
    var advance int = 0
    
    for i >= 0 || j >= 0 {
        
        sum := advance
        
        if i >= 0 {
            sum = sum + int(a[i]) - 48
            i--
        }
        
        if j >= 0 {
            sum = sum + int(b[j]) - 48
            j--
        }
        
        ret = append(ret, byte(sum % 2 + 48))
        advance = sum / 2
    }
    
    if advance == 1 {
        ret = append(ret, byte(1 + 48))
    }
    
    for left,right := 0,len(ret)-1; left <= right; {
        temp := ret[left]
        ret[left] = ret[right]
        ret[right] = temp
        
        left++
        right--
    }
    
    return string(ret)
}