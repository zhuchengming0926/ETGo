/**
 * @File: jishuFrontOushu.go
 * @Author: zhuchengming
 * @Description:奇数调整到偶数之前，相对位置不变
 * @Date: 2021/3/3 20:59
 */

/*
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前半部分，
所有的偶数位于数组的后半部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

  解题思路：

  首先，如果不考虑奇数和奇数，偶数和偶数的相对位置，那么我们有一种双指针解法来求解，类似于快排，维护两个指针，
第一个指针指向数组的第一个数字，第二个指针指向数组的最后一个数字。第一个指针向后移，第二个指针向前移，
如果第一个指针指向偶数，第二个指针指向的是奇数，则交换着两个数字，接着继续移动直到两指针相遇。

  上面的方法看似不错，但是对本题不适用，因为本题有相对位置不变的要求，直接交换会导致相对位置改变。
因此，我们采用下面的思路来解决本题。

  本题解法：对数组进行遍历，设置两个指针even和odd，even指向当前第一个偶数，odd从这个偶数之后开始查找，
找到第一个奇数，此时为了相对位置不变，不能直接交换even和odd，而是将从even到odd-1的元素都依次向后移一个位置，
将odd指向的那个奇数放到even的位置。然后再找下一个偶数，重复这一过程，最终就可以将奇数都放到偶数的前面，
并且保证了相对位置的不变。*/

package main

func exchange(nums []int) []int {
	length := len(nums)
	if length <= 0 {
		return nums
	}

	findOuShu := 0
	findJiShu := 0
	for ; findOuShu < length && findJiShu < length; {
		//从前到后找到第一个偶数
		for ; findOuShu < length && nums[findOuShu] % 2 != 0; {
			findOuShu++
		}
		if findOuShu >= length {
			return nums
		}
		if findJiShu < findOuShu + 1 {
			findJiShu = findOuShu + 1
		}

		for ; findJiShu < length && nums[findJiShu] % 2 == 0; {
			findJiShu++
		}
		if findJiShu >= length {
			return nums
		}
		//找到了奇数和偶数位置，不能直接交换，而是往后一位移动findOuShu和findJiShu-1这些数
		temp := nums[findJiShu]
		for i := findJiShu-1; i >= findOuShu; i-- {
			nums[i+1] = nums[i]
		}
		nums[findOuShu] = temp
		findJiShu++
		findOuShu++
	}
	return nums
}

//通过首尾夹逼法
func exchange2(nums []int) []int {
	length := len(nums)
	if length <= 0 {
		return nums
	}

	low := 0
	high := length - 1

	//从两边夹逼
	for low < high {
		for low < high && (nums[low] & 1) == 1 {
			low++
		}
		if low >= length {
			return nums
		}
		for low < high && (nums[high] & 1) == 0 {
			high--
		}
		if high <= low {
			return nums
		}
		//互换
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
	return nums
}