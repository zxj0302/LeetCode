/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
package main

func longestCommonPrefix(strs []string) string {
	comPrefix := uint8(0)
	for{
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == int(comPrefix) || strs[i][comPrefix] != strs[0][comPrefix]{
				return strs[0][:comPrefix]
			}
		}
		comPrefix++
	}
}
// @lc code=end

