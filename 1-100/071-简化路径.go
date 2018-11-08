/* 题目

https://leetcode-cn.com/problems/simplify-path/description/

给定一个文档 (Unix-style) 的完全路径，请进行路径简化。

Example
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"

边界情况:

你是否考虑了 路径 = "/../" 的情况？
在这种情况下，你需返回 "/" 。
此外，路径中也可能包含多个斜杠 '/' ，如 "/home//foo/" 。
在这种情况下，你可忽略多余的斜杠，返回 "/home/foo" 。

*/

/* 思路

先用 strings.Split 以 / 为标准进行切分
用栈解决，一共有 字母，..，.，以及""三种情况

遇到 字母 则入栈
遇到 . 以及 "" 则不动
遇到 .. 则出栈

*/

func simplifyPath(path string) string {

    var src []string = strings.Split(path,"/")
    
    var stack []string
    
    for _, item := range src {
        if item == ".." {
            if len(stack) > 0 {
                stack = stack[0:len(stack)-1]
            }
        } else if item != "." && item != "" && item != ".." {
            stack = append(stack, item)
        }
    }
    
    var ret string = ""
    
    for len(stack) > 0 {
        ret = "/" + stack[len(stack)-1] + ret
        stack = stack[0:len(stack)-1]
    }
    
    if ret == "" {
        return "/"
    }
    
    return ret
}