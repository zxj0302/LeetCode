/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 *
 * https://leetcode.cn/problems/4sum/description/
 *
 * algorithms
 * Medium (39.31%)
 * Likes:    1236
 * Dislikes: 0
 * Total Accepted:    314.2K
 * Total Submissions: 799.3K
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a],
 * nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
 * 
 * 
 * 0 <= a, b, c, d < n
 * a、b、c 和 d 互不相同
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 * 
 * 
 * 你可以按 任意顺序 返回答案 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,0,-1,0,-2,2], target = 0
 * 输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [2,2,2,2,2], target = 8
 * 输出：[[2,2,2,2]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 
 * 
 */

// @lc code=start
package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if target > 0 && nums[i] >= target {
			break
		}
		for j := i+1; j < len(nums)-2; j++ {
			if j > i + 1 && nums[j] == nums[j-1]{
				continue
			}
			if target > 0 && nums[i] + nums[j] >= target {
				break
			}
			k := j + 1
			r := len(nums)-1
			for ; k < r; k++ {
				if k > j+1 && nums[k] == nums[k-1]{
					continue
				}
				for ; k < r; r-- {
					if nums[i] + nums[j] + nums[k] + nums[r] > target {
						continue
					}else if nums[i] + nums[j] + nums[k] + nums[r] == target{
						result = append(result, []int{nums[i], nums[j], nums[k], nums[r]})
					}
					break
				}
			}
		}
	}
	return result
}
// @lc code=end

