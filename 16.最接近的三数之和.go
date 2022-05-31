/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode.cn/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (45.66%)
 * Likes:    1143
 * Dislikes: 0
 * Total Accepted:    350.7K
 * Total Submissions: 768.4K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
 * 
 * 返回这三个数的和。
 * 
 * 假定每组输入只存在恰好一个解。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [-1,2,1,-4], target = 1
 * 输出：2
 * 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [0,0,0], target = 1
 * 输出：0
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 3 <= nums.length <= 1000
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 * 
 * 
 */

// @lc code=start
package main
import "sort"
import "math"
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt64
	for i := 0; i < len(nums) - 2; i++ {
		// 剪枝
		if nums[i] >= target && target > 0{
			if nums[i] + nums[i+1] + nums[i+2] - target >= abs(minDiff) {
				return minDiff + target
			}else{
				return nums[i] + nums[i+1] + nums[i+2]
			}
		}

		k := len(nums) -  1
		for j := i + 1; j < len(nums) - 1 && j < k; j++ {
			// 剪枝
			if nums[i] + nums[j] >= target && target > 0{
				if nums[i] + nums[j] + nums[j+1] - target >= abs(minDiff) {
					return minDiff + target
				}else{
					return nums[i] + nums[j] + nums[j+1]
				}
			}

			for ; k > j; k-- {
				diff := nums[k] + nums[j] + nums[i] - target
				if diff < 0{
					if abs(diff) < abs(minDiff) {
						minDiff = diff
					}
					if k < len(nums) - 1 && abs(nums[k+1] + nums[j] + nums[i] - target) < abs(minDiff) {
						minDiff = nums[k+1] + nums[j] + nums[i] - target
					}
					break
				}else if diff == 0 {
					return target
				}
			}
			if k == j && abs(nums[k+1] + nums[j] + nums[i] - target) < abs(minDiff){
				minDiff = nums[k+1] + nums[j] + nums[i] - target
			}
		}
	}
	return target + minDiff
}
// @lc code=end

