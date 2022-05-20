/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package main
func isValid(s string) bool {
	if len(s) % 2 == 1{
		return false
	}
	stack := make([]byte, 0, len(s) / 2)
	for _, c := range s{
		if len(stack) == 0 {
			stack = append(stack, byte(c))
			continue
		}
		if len(stack) > len(s) / 2{
			return false
		}
		switch stack[len(stack)-1]{
			case '(':
				switch c{
					case ')':
						stack = stack[:len(stack)-1]
					case '{', '[', '(':
						stack = append(stack, byte(c))
					default:
						return false
				}
			case '[':
				switch c{
					case ']':
						stack = stack[:len(stack)-1]
					case '{', '[', '(':
						stack = append(stack, byte(c))
					default:
						return false
				}
			case '{':
				switch c{
					case '}':
						stack = stack[:len(stack)-1]
					case '{', '[', '(':
						stack = append(stack, byte(c))
					default:
						return false
				}
			default:
				return false
		}
	}
	if len(stack) == 0 {
		return true
	}else{
		return false
	}
}

// @lc code=end

