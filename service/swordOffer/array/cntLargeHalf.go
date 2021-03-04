/**
 * @File: cntLargeHalf.go
 * @Author: zhuchengming
 * @Description:数组中重复次数超过一半的数
 * @Date: 2021/3/3 21:34
 */

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
 例如：输入如下所示的一个长度为9的数组{1,2,3,2,2,2,5,4,2}。
由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。如果不存在则输出0。

解题思路：

  本题有以下三种方法可解：

  方法一：首先对数组进行排序，在一个有序数组中，次数超过一半的必定是中位数，那么可以直接取出中位数，
然后遍历数组，看中位数是否出现次数超过一半，这取决于排序的时间复杂度，最快为O(nlogn)。

  方法二：遍历数组，用 HashMap 保存每个数出现的次数，这样可以从map中直接判断是否有超过一半的数字，
这种算法的时间复杂度为O(n)，但是这个性能提升是用O(n)的空间复杂度换来的。

  方法三（最优解法）：根据数组特点得到时间复杂度为O(n)的算法。根据数组特点，数组中有一个数字出现的次数超过数组长度的
一半，也就是说它出现的次数比其他所有数字出现的次数之和还要多。因此，我们可以在遍历数组的时候设置两个值：
一个是数组中的数result，另一个是出现次数times。当遍历到下一个数字的时候，如果与result相同，则次数加1,不同则次数减一，
当次数变为0的时候说明该数字不可能为多数元素，将result设置为下一个数字，次数设为1。这样，当遍历结束后，
最后一次设置的result的值可能就是符合要求的值（如果有数字出现次数超过一半，则必为该元素，否则不存在），
因此，判断该元素出现次数是否超过一半即可验证应该返回该元素还是返回0。这种思路是对数组进行了两次遍历，复杂度为O(n)。
*/

package main

import "fmt"

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return nums[0]
	}

	curValue := nums[0]
	curCount := 1
	for i:= 1; i < length; i++ {
		fmt.Println(curValue, "=>", curCount)
		if nums[i] != curValue {
			curCount--
		} else {
			curCount++
		}
		if curCount == 0 {
			curValue = nums[i]
			curCount = 1
		}
	}
	curCount = 0
	for i := 0; i < length; i++ {
		if nums[i] == curValue {
			curCount++
		}
	}
	if curCount > length /2 {
		return curValue
	}
	return -1
}

//func main()  {
//	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
//	fmt.Println(majorityElement(nums))
//}