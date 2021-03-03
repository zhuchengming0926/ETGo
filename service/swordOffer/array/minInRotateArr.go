/**
 * @File: minInRotateArr.go
 * @Author: zhuchengming
 * @Description:11-旋转数组中的最小数字
 * @Date: 2021/2/26 12:03
 */

package main

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。 输入一个非减排序的数组的一个旋转，输出旋转数组的最小元素。
例如数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转，该数组的最小值为1。 NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

三种情况：非递减
        （1）把前面0个元素搬到末尾，也就是排序数组本身，第一个就是最小值
        （2）一般情况二分查找，当high-low=1时，high就是最小值
        （3）如果首尾元素和中间元素都相等时，只能顺序查找
*/

func minArray(numbers []int) int {
	length := len(numbers)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return numbers[0]
	}
	//情况一：没有旋转
	if numbers[0] < numbers[length-1] {
		return numbers[0]
	}

	//情况二：有旋转
	low := 0
	high := length - 1
	for ; low < high; {

		mid := low + (high - low) / 2
		if numbers[low] == numbers[mid] && numbers[mid] == numbers[high] { //特殊情况，你不知道是跳动high还是low位置
			min := numbers[low]
			for i:=low+1; i <= high; i++ {
				if numbers[i] < min {
					min = numbers[i]
				}
			}
			return min
		}
		if numbers[low] <= numbers[mid] {
			low = mid
		} else if numbers[high] >= numbers[mid] {
			high = mid
		}
		if low + 1 == high { //因为有旋转的时候，最小值在low和high中间夹着，low一定比high小
			return numbers[high]
		}
	}
	return -1
}

//func main()  {
//	nums := []int{3,4,5,1,2}
//	fmt.Println(minArray(nums))
//}