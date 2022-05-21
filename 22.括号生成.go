/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.45%)
 * Likes:    2643
 * Dislikes: 0
 * Total Accepted:    507K
 * Total Submissions: 654.5K
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：n = 1
 * 输出：["()"]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= n <= 8
 * 
 * 
 */

// @lc code=start
package main

func generateParenthesis(n int) []string {
	result := []string{}
	for _, c := range generate(0, 2*n){
		result = append(result, string(c))
	}
	return result
}

func generate(diff, n int) [][]byte{
	result := [][]byte{}
	if n == 1{
		return [][]byte{{')'}}
	}
	if diff == 0{
		for _, c := range generate(1, n-1){
			result = append(result, append([]byte{'('}, c...))
		}
	}else{
		for _, c := range generate(diff-1, n-1){
			result = append(result, append([]byte{')'}, c...))
		}
		if (n-diff)/2 > 0 {
			for _, c := range generate(diff+1, n-1){
				result = append(result, append([]byte{'('}, c...))
			}
		}
	}
	return result
}
// @lc code=end

