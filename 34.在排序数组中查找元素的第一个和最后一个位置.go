/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 *
 * https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
 *
 * algorithms
 * Medium (42.24%)
 * Likes:    1674
 * Dislikes: 0
 * Total Accepted:    531.3K
 * Total Submissions: 1.3M
 * Testcase Example:  '[5,7,7,8,8,10]\n8'
 *
 * 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
 *
 * 如果数组中不存在目标值 target，返回 [-1, -1]。
 *
 * 进阶：
 *
 *
 * 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 8
 * 输出：[3,4]
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,7,7,8,8,10], target = 6
 * 输出：[-1,-1]
 *
 * 示例 3：
 *
 *
 * 输入：nums = [], target = 0
 * 输出：[-1,-1]
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^9
 * nums 是一个非递减数组
 * -10^9
 *
 *
 */

// @lc code=start
package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}
	index1, index2 := -1, -1
	left, right := 0, len(nums)-1
	for middle := (left + right) / 2; left < right; middle = (left + right) / 2 {
		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	if nums[left] == target {
		index1 = left
	} else if left != len(nums)-1 && nums[left+1] == target {
		index1 = left + 1
	} else {
		return []int{-1, -1}
	}
	left, right = 0, len(nums)-1
	for middle := (left + right) / 2; left < right; middle = (left + right) / 2 {
		if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if nums[left] == target {
		index2 = left
	} else if left != 0 && nums[left-1] == target {
		index2 = left - 1
	} else {
		return []int{-1, -1}
	}
	return []int{index1, index2}
}

// @lc code=end
