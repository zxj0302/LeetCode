/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (35.39%)
 * Likes:    4761
 * Dislikes: 0
 * Total Accepted:    965.2K
 * Total Submissions: 2.7M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = []
 * 输出：[]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [0]
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * -10^5 
 * 
 * 
 */

// @lc code=start
package main

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums) - 2; i++ {
		if nums[i] > 0{
			break
		}
		if i > 0 && nums[i-1] == nums[i]{
			continue
		}
		k := len(nums) - 1
		for j := i + 1; j < len(nums) - 1; j++ {
			if j > i + 1 && nums[j-1] == nums[j]{
				continue
			}
			for ;k > j;k--{
				if k < len(nums) - 1 && nums[k] == nums[k + 1]{
					continue
				}
				if nums[k] == -nums[j]-nums[i]{
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}else if nums[k] < -nums[j]-nums[i]{
					break
				}
			}
		}
	}
	return result
}
// @lc code=end

