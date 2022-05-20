/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
package main

var dataMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}
	for _, digit := range digits {
		result = cartesian(result, byte(digit))
	}
	return result
}

func cartesian(s []string, b byte) []string {
	ret := []string{}
	for _, v := range s {
		for _, c := range dataMap[b]{
			ret = append(ret, v + string(c))
		}
	}
	return ret
}
// @lc code=end

