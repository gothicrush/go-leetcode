/* 题目
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

Example
输入: ["flower","flow","flight"]
输出: "fl"

Example
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/

/* 思路
如果输入字符串数组为nil或长度为0，则返回""

取第一个字符串作为基准，遍历整个字符串数组
通过 strings.HasPrefix 判断基准是否满足最长公共前缀
如果满足，则判断下一个字符串
如果不满足，则通过切片，缩少基准，不断重复判断
*/

func longestCommonPrefix(strs []string) string {
    
    if strs == nil || len(strs) == 0 {
        return ""
    }
    
    var ret string = strs[0]
    
    for i := 1; i < len(strs); i++ {
        for !strings.HasPrefix(strs[i], ret) {
            ret = ret[:len(ret)-1]
        }
    }
    
    return ret
}