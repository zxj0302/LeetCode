/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode.cn/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (48.74%)
 * Likes:    1105
 * Dislikes: 0
 * Total Accepted:    265.4K
 * Total Submissions: 544K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
 * 输出：[1,2,3,6,9,8,7,4,5]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
 * 输出：[1,2,3,4,8,12,11,10,9,5,6,7]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == matrix.length
 * n == matrix[i].length
 * 1 
 * -100 
 * 
 * 
 */

// @lc code=start
package main
func spiralOrder(matrix [][]int) []int {
	i, j := 0, -1
	m := len(matrix)
	n := len(matrix[0])
	result := make([]int, m * n)
	
	for count := 0; count != m * n; count++ {
		switch{
			case (j < n - 1 && matrix[i][j + 1] != -101) && (i == 0 || matrix[i - 1][j] == -101):
				result[count] = matrix[i][j + 1]
				matrix[i][j + 1] = -101
				j++
			case i < m - 1 && matrix[i + 1][j] != -101:
				result[count] = matrix[i + 1][j]
				matrix[i + 1][j] = -101
				i++
			case j > 0 && matrix[i][j - 1] != -101:
				result[count] = matrix[i][j - 1]
				matrix[i][j - 1] = -101
				j--
			case i > 0 && matrix[i - 1][j] != -101:
				result[count] = matrix[i - 1][j]
				matrix[i - 1][j] = -101
				i--
		}
	}
	return result
}
// @lc code=end

