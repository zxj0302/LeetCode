/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 *
 * https://leetcode.cn/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (45.32%)
 * Likes:    1559
 * Dislikes: 0
 * Total Accepted:    802.2K
 * Total Submissions: 1.8M
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
 * 
 * 请必须使用时间复杂度为 O(log n) 的算法。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: nums = [1,3,5,6], target = 5
 * 输出: 2
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: nums = [1,3,5,6], target = 2
 * 输出: 1
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: nums = [1,3,5,6], target = 7
 * 输出: 4
 * 
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 为 无重复元素 的 升序 排列数组
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
package main
func searchInsert(nums []int, target int) int {
	if nums[0] > target {
		return 0
	}else if nums[len(nums)-1] < target {
		return len(nums)
	}
	left, right := 0, len(nums) - 1
	for middle := (left + right) / 2; left <= right; middle = (left + right) / 2 {
		if nums[middle] == target {
			return middle
		}else if nums[middle] < target {
			left = middle + 1
		}else{
			right = middle - 1
		}
	}
	return left
}
// @lc code=end

