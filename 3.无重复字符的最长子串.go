/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
package main

func lengthOfLongestSubstring(s string) int {
	max, left, char2Pos := uint16(0), uint16(0), make(map[byte]uint16)
	for i, v := range s {
		if value, ok := char2Pos[byte(v)]; ok{
			if uint16(i) - left > max {
				max = uint16(i) - left
			}
			for start := left; start < value; start++{
				delete(char2Pos, byte(s[start]))
			}
			left = value + 1
		}
		char2Pos[byte(v)] = uint16(i)
	}
	if len(s) - int(left) > int(max) && len(char2Pos) > 0{
		return len(s) - int(left)
	}
	return int(max)
}
// @lc code=end

