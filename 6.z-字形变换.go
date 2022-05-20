/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
package main

func convert(s string, numRows int) string {
	if numRows >= len(s)|| numRows == 1{
		return s
	}

	result := make([][]byte, numRows)
	circular := 2 * numRows - 2
	for i, v := range s {
		j := i % circular
		if j >= numRows{
			j = circular - j
		}
		result[j] = append(result[j], byte(v))
	}

	res := ""
	for _, v := range result{
		res += string(v)
	}
	return res
}
// @lc code=end

