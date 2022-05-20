/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
package main

func intToRoman(num int) string {
	s := make([]byte, 0)
	for num != 0 {
		switch {
			case num >= 1000:
				num -= 1000
				s = append(s, 'M')
			case num >= 900:
				num -= 900
				s = append(s, 'C')
				s = append(s, 'M')
			case num >= 500:
				num -= 500
				s = append(s, 'D')
			case num >= 400:
				num -= 400
				s = append(s, 'C')
				s = append(s, 'D')
			case num >= 100:
				num -= 100
				s = append(s, 'C')
			case num >= 90:
				num -= 90
				s = append(s, 'X')
				s = append(s, 'C')
			case num >= 50:
				num -= 50
				s = append(s, 'L')
			case num >= 40:
				num -= 40
				s = append(s, 'X')
				s = append(s, 'L')
			case num >= 10:
				num -= 10
				s = append(s, 'X')
			case num >= 9:
				num -= 9
				s = append(s, 'I')
				s = append(s, 'X')
			case num >= 5:
				num -= 5
				s = append(s, 'V')
			case num >= 4:
				num -= 4
				s = append(s, 'I')
				s = append(s, 'V')
			default:
				num -= 1
				s = append(s, 'I')
		}
	}
	return string(s)
}
// @lc code=end

