/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
package main

func romanToInt(s string) int {
	var result int16
	length := len(s)
	for i, v := range s {
		switch v {
		case 'I':
			if i + 1 < length && (s[i + 1] == 'V' ||  s[i + 1] == 'X'){
				result -= 1
			}else{
				result += 1
			}
		case 'V':
			result += 5
		case 'X':
			if i + 1 < length && (s[i + 1] == 'L' ||  s[i + 1] == 'C'){
				result -= 10
			}else{
				result += 10
			}
		case 'L':
			result += 50
		case 'C':
			if i + 1 < length && (s[i + 1] == 'D' ||  s[i + 1] == 'M'){
				result -= 100
			}else{
				result += 100
			}
		case 'D':
			result += 500
		case 'M':
			result += 1000
		}
	}
	return int(result)
}
// @lc code=end

