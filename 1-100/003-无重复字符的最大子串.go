/* 题目
给定一个字符串，找出不含有重复字符的最长子串的长度。

Example
输入: "abcabcbb"
输出: 3 
解释: 无重复字符的最长子串是 "abc"，其长度为 3。

Example
输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。

Example
输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*/

/* 思路
使用滑动窗口思想
当无重复元素时，右边移动
当有重复元素时，从左开始循环删除，直到删掉重复元素
*/

func lengthOfLongestSubstring(s string) int {
    
    var ret int = 0
    var m map[byte]bool = make(map[byte]bool)
    
    for i,j := 0,0; i < len(s) && j < len(s); {
        
        if _, ok := m[s[i]]; !ok {
            m[s[i]] = true
            i++
            ret = max(ret, i-j)
        } else {
            delete(m, s[j])
            j++
        }
    }
    
    return ret
}

func max(i int, j int) int {
    if i < j {
        return j
    } else {
        return i
    }
}