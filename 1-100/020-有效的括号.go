/* 题目

https://leetcode-cn.com/problems/valid-parentheses/description/

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。

注意空字符串可被认为是有效字符串。

Example
输入: "()"
输出: true

Example
输入: "()[]{}"
输出: true

Example
输入: "(]"
输出: false

Example
输入: "([)]"
输出: false

Example
输入: "{[]}"
输出: true
*/

/* 思路
用栈解决
遇到 ( 则入栈 )；遇到 [ 则入栈 ]；遇到 { 则入栈 }
遇到 ) 则出栈，判断出栈元素是否是 )，不是则返回false
遇到 ] 则出栈，判断出栈元素是否是 ]，不是则返回false
遇到 } 则出栈，判断出栈元素是否是 }，不是则返回false

最后判断栈是否为空，不为空则返回false，否则返回true
*/

func isValid(s string) bool {
    
    var stack []string
    
    for _, ru := range s {
        ch := string(ru)
        if ch == "(" {
            stack = append(stack, ")")
        } else if ch == "[" {
            stack = append(stack, "]")
        } else if ch == "{" {
            stack = append(stack, "}")
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != ch {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0
}