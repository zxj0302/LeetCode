/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start

package main
func isPalindrome(x int) bool {
	if x < 0{
		return false
	}

	s := make([]uint8, 0)
	for ;x != 0; x /= 10 {
		s = append(s, uint8(x % 10))
	}
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s)-1-i]{
			return false
		}
	}
	return true
}
// @lc code=end

