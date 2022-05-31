/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 *
 * https://leetcode.cn/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (60.77%)
 * Likes:    979
 * Dislikes: 0
 * Total Accepted:    299K
 * Total Submissions: 492.4K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的每个数字在每个组合中只能使用 一次 。
 * 
 * 注意：解集不能包含重复的组合。 
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 输出:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 * 
 * 示例 2:
 * 
 * 
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 输出:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 * 
 * 
 * 提示:
 * 
 * 
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 * 
 * 
 */

// @lc code=start
package main
import "sort"
func combinationSum2(candidates []int, target int) (ans [][]int) {
    sort.Ints(candidates)
    var freq [][2]int
    for _, num := range candidates {
        if freq == nil || num != freq[len(freq)-1][0] {
            freq = append(freq, [2]int{num, 1})
        } else {
            freq[len(freq)-1][1]++
        }
    }

    var sequence []int
    var dfs func(pos, rest int)
    dfs = func(pos, rest int) {
        if rest == 0 {
            ans = append(ans, append([]int(nil), sequence...))
            return
        }
        if pos == len(freq) || rest < freq[pos][0] {
            return
        }

        dfs(pos+1, rest)

        most := min(rest/freq[pos][0], freq[pos][1])
        for i := 1; i <= most; i++ {
            sequence = append(sequence, freq[pos][0])
            dfs(pos+1, rest-i*freq[pos][0])
        }
        sequence = sequence[:len(sequence)-most]
    }
    dfs(0, target)
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
// @lc code=end

