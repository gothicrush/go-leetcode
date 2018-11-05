/* 题目
罗马数字转整数，规则看 012
*/

/* 思路
定义一个罗马字母转整数的函数

注意：对于两个罗马字母，如果 左边 < 右边，则它们的实际组合起来代表一个值，该值为 右边 - 左边

从下标1开始遍历字符串：
如果当前下标的罗马字母小于左边罗马字母，则直接将其对应的整数加入到统计值中
如果当前下标的罗马字母大于左边罗马字母，则它们的值为 右 - 左 - 左，加入统计值中
需要再减去一个左的原因是在上一次的遍历中，已经将左边的值计入总值中
*/

func romanToInt(s string) int {
    
    var ret = symbolToInt(string(s[0]))
    
    for i := 1; i < len(s); i++ {
        if symbolToInt(string(s[i])) > symbolToInt(string(s[i-1])) {
            ret = ret + (symbolToInt(string(s[i])) - 2 * symbolToInt(string(s[i-1])))
        } else {
            ret = ret + symbolToInt(string(s[i]))
        }
    }
    
    
    return ret
}

func symbolToInt(s string) int {
    var ret = 0
    
    switch s {
    case "M": ret = 1000
    case "D": ret = 500
    case "C": ret = 100
    case "L": ret = 50
    case "X": ret = 10
    case "V": ret = 5
    case "I": ret = 1
    }
    
    return ret
}