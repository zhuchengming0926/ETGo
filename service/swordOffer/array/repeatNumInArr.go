/**
 * @File: repeatNumInArr.go
 * @Author: zhuchengming
 * @Description:数组中重复的数
 * @Date: 2021/3/4 17:44
 */

/*
在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。
也不知道每个数字重复几次。请找出数组中任意一个重复的数字。
例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是第一个重复的数字2。
*/

package main

func repeatNum(nums []int) int {
	//第一种方法：
	//扫描数组，每次判断numbers[numbers[i]]是否大于0，若大于0，则
	//把它减去length，若小于0表示这个数numbers[i]就是一个重复数字
	//因为是重复数字，肯定在上一次碰到这个数字时把对应位置的数字减去length
	//了，下次再碰到这个数，那么对于位置的数字肯定小于0了
	length := len(nums)
	if length <= 1 {
		return -1
	}
	for i:=0; i<length; i++{
		index := nums[i]
		if index < 0 { //表示数组中存在一个值为i，只有这样nums[i]才会变成负数
			index = index + length
		}
		if nums[index] >= 0 {//表示index下标上的值还没被改变，此时把它变成负数
			nums[index] = nums[index] - length
		} else {//表示index下标上的值已经被改变过，说明前边已经存在一个和index一样的值，导致nums[index]改变了
			return index
		}
	}
	return -1
}