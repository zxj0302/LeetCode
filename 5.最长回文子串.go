/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
package main

func longestPalindrome(s string) string {
	matrix := make([][]bool, len(s))
	for i := range matrix {
		matrix[i] = make([]bool, i + 1)
	}
	for i := 0; i < len(s); i++{
		matrix[i][i] = true
	}
	for i := 1; i < len(s); i++{
		if s[i] == s[i - 1]{
			matrix[i][i - 1] = true
		}else{
			matrix[i][i - 1] = false
		}
	}
	for i := 2; i < len(s); i++{
		for j := i; j < len(s); j++{
			if s[j] == s[j - i] && matrix[j - 1][j - i + 1]{
				matrix[j][j - i] = true
			}else{
				matrix[j][j - i] = false
			}
		}
	}

	for i := len(s) - 1; i >= 0; i--{
		for j := 0; j < len(s) - i; j++{
			if matrix[i + j][j]{
				return s[j : j + i + 1]
			}
		}
	}
	return s
}

// @lc code=end

