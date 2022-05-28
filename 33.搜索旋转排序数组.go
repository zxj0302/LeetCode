/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (43.55%)
 * Likes:    2096
 * Dislikes: 0
 * Total Accepted:    540.1K
 * Total Submissions: 1.2M
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 整数数组 nums 按升序排列，数组中的值 互不相同 。
 * 
 * 在传递给函数之前，nums 在预先未知的某个下标 k（0 ）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ...,
 * nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如，
 * [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
 * 
 * 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1], target = 0
 * 输出：-1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^4 
 * nums 中的每个值都 独一无二
 * 题目数据保证 nums 在预先未知的某个下标上进行了旋转
 * -10^4 
 * 
 * 
 * 
 * 
 * 进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？
 * 
 */

// @lc code=start
package main

func search(nums []int, target int) int {
	k := findk(nums)
	if target < nums[0]{
		temp := binarySearch(nums[k+1:], target)
		if temp != -1{
			temp += k+1
		}
		return temp
	}else if target > nums[0]{
		temp := binarySearch(nums[:k+1], target)
		return temp
	}else{
		return 0
	}
}

func findk(nums []int) int {
	left, right := 0, len(nums)-1
	for {
		middle := (left + right) / 2
		if nums[(middle+1)%len(nums)] <= nums[middle]{
			return middle
		}
		if nums[middle] < nums[0]{
			right = middle-1
		}else{
			left = middle+1
		}
	}
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right{
		middle := (left + right) / 2
		if nums[middle] == target{
			return middle
		}else if nums[middle] < target{
			left = middle + 1
		}else{
			right = middle - 1
		}
	}
	return -1
}
// @lc code=end

