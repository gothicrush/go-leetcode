/* 题目

https://leetcode-cn.com/problems/longest-palindromic-substring/description/

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

Example
输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。

Example
输入: "cbbd"
输出: "bb"
*/

/* 思路
遍历每一个字母，每一个字母从中间向两边搜素最大回文串
*/

func longestPalindrome(s string) string {
    
    if len(s) == 0 {
        return s
    }
    
    var result string
    
    for i := 0; i < len(s); i++ {
        process(s,i,i,&result)
        process(s,i,i+1,&result)
    }
    
    return result
}

func process(s string, left int, right int, result *string) {
    
    for left >=0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    
    var temp string = string(s[left+1:right])
    
    if len(temp) > len(*result) {
        *result = temp
    }
}