/* 题目

https://leetcode-cn.com/problems/set-matrix-zeroes/description/

给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

Example
输入: 
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出: 
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]

Example
输入: 
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出: 
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]

*/

/* 思路

遍历数组，如果一行/列上有0，则该行/列第一个元素标记为0

再次遍历数组，通过标记将整行/列置零即可

*/

func setZeroes(matrix [][]int)  {

    var rowflag bool = false
    var colflag bool = false
    
    var row int = len(matrix)
    var col int = len(matrix[0])
    
    for j := 0; j < col; j++ {
        if matrix[0][j] == 0 {
            rowflag = true
        }
    }
    
    for i := 0; i < row; i++ {
        if matrix[i][0] == 0 {
            colflag = true
        }
    }
    
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    
    for i := 0; i < row; i++ {
        if matrix[i][0] == 0 {
            for j := 0; j < col; j++ {
                matrix[i][j] = 0
            }
        }
    }
    
    for j := 0; j < col; j++ {
        if matrix[0][j] == 0 {
            for i := 0; i < row; i++ {
                matrix[i][j] = 0
            }
        }
    }
    
    if rowflag {
        for i := 0; i < col; i++ {
            matrix[0][i] = 0
        }
    }
    
    if colflag {
        for i := 0; i < row; i++ {
            matrix[i][0] = 0
        }
    }
}