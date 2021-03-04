/**
 * @File: strToInt.go
 * @Author: zhuchengming
 * @Description:字符串转整数
 * @Date: 2021/3/4 16:04
 */

package string

import "math"

//0-9的ascii码值是 48-57  “-”号是45，“+”号是43
//func myAtoi(s string) int {
//	s = strings.TrimSpace(s)
//	if len(s) == 0 {
//		return 0
//	}
//	flag := true
//	sum := 0
//	maxInt := (1 << 31) - 1
//	minInt := -(1 << 31)
//
//	for i, x := range s {
//		//处理正负号
//		if i == 0 &&  int(x) == 45 {
//			flag = false
//			continue
//		}
//		if i == 0 && int(x) == 43 {
//			continue
//		}
//		//处理其它数字
//		if int(x) < 48 || int(x) > 57 {
//			break
//		}
//		sum = sum * 10 + int(x) - 48
//		if flag && sum > maxInt {
//			return maxInt
//		}
//
//		if !flag {
//			temp := -sum
//			if temp < minInt {
//				return minInt
//			}
//		}
//	}
//	if !flag {
//		sum = -sum
//	}
//	return sum
//}

//不需要记住ascii码
func myAtoi(s string) int {
	var value int64
	var maxValue = int64(math.MaxInt32)
	var minValue = int64(math.MinInt32)

	var flag = true        // 前导空格
	var isNegative = false // -
begin:
	for i := 0; i < len(s); i++ {
		// 读入字符串并丢弃无用的前导空格
		switch s[i] {
		case ' ':
			if flag {
				continue
			}
			break begin
		case '-':
			if flag {
				isNegative = true
				flag = false
				continue
			}
			break begin
		case '+':
			if flag {
				flag = false
				continue
			}
			break begin
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			flag = false
			value = value*10 + int64(s[i]-'0')
			if isNegative == false && value > maxValue {
				value = maxValue
				break begin
			}

			if isNegative && value > -minValue {
				value = -minValue
				break begin
			}
		default:
			break begin
		}
	}

	if isNegative {
		value = -value
	}

	return int(value)
}
