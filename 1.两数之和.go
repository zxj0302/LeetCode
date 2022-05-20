/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
package main

func twoSum(nums []int, target int) []int {

	m := make(map[int]int, len(nums))

	for i, n := range nums {
		v, ok := m[target - n]
		if ok{
			return []int{i, v}
		}
		m[n] = i
	}

	return nil
}
// @lc code=end

