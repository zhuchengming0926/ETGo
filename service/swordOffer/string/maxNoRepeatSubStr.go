/**
 * @File: maxNoRepeatSubStr.go
 * @Author: zhuchengming
 * @Description:最大无重复子串
 * @Date: 2021/3/4 20:17
 */

package string

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度

思路；
1，双重for循环，暴力法
2，窗口滑动
找到重复的 让left指针跳到重复位置的下一个 right指针不动
直到没有重复的时候right才往前进
s[left:right]大于当前长度才+1
*/

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}

	left := 0
	right := 0
	maxLength := 0
	for right < length {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				left = i+1
				break
			}
		}
		right++
		if len(s[left:right]) > maxLength {
			maxLength = len(s[left:right])
		}
	}
	return maxLength
}