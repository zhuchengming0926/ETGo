/**
 * @File: sort.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/18 20:14
 */

package svAlgorithm

import (
	"ETGo/components/defines"
	"ETGo/components/dto/dtoAlgorithm"
	"errors"
	"fmt"
)

var VarousSortFuncMap = map[int]func(data []int, order int) (*dtoAlgorithm.SortRes, error) {
	defines.KSORT     :Ksort,
	defines.HEAPSORT  :HeapSort,
	defines.BUBBLESORT:BubbleSort,
	defines.MERGESORT :MergeSort,
}

func VariousSort(req *dtoAlgorithm.SortReq) (*dtoAlgorithm.SortRes, error) {
	if v, ok := VarousSortFuncMap[req.Type]; ok {
		return v(req.Data, req.Type)
	} else if req.Type == 2 {
		temp := []int{Bsearch(req.Data, 2)}
		return &dtoAlgorithm.SortRes{Ret:temp}, nil
	}
	return nil, errors.New("not exist corresponding function")
}

//1--快速排序
func Ksort(data []int, order int) (*dtoAlgorithm.SortRes, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is null")
	}

	QuickSort(data, 0, length - 1)
	return &dtoAlgorithm.SortRes{Ret:data}, nil
}

//快速排序基准操作-找到指定枢轴的最终位置
func OneKSortOpration(data []int, low, high int) int {
	if len(data) <= 1 {
		return low
	}
	//默认data[low]为枢轴
	key := data[low]
	for ;low < high; {
		for ;low < high && data[high] >= key; {
			high--
		}
		if low < high {
			data[low] = data[high]
			low++
		}
		for ;low < high && data[low] <= key; {
			low++
		}
		if low < high {
			data[high] = data[low]
			high--
		}
	}
	//跳出循环时，low和high相等
	data[low] = key
	return low
}

func QuickSort(data []int, low, high int)  {
	if low < high {
		standard := OneKSortOpration(data, low, high)
		QuickSort(data, low, standard - 1)
		QuickSort(data, standard + 1, high)
	}
}

//2--二分查找,建立在有序数组上
func Bsearch(data []int, target int) int {
	printArray(data)
	length := len(data)
	if length == 0 {
		return -1
	}
	low := 0
	high := length - 1

	for ;low <= high; {
		mid := (high + low) / 2
		fmt.Println(mid)
		if data[mid] == target {
			return mid
		} else if data[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}


//3--堆排序
func HeapSort(data []int, order int) (*dtoAlgorithm.SortRes, error) {
	length := len(data)
	if length <= 1 {
		return &dtoAlgorithm.SortRes{Ret:data}, nil
	}
	genHeap(data, order) //构建一个初始堆
	printArray(data)

	for i := 0; i < length; i++ {
		//最终排序时，都是交换0号位置，再调整0号位置
		data[length - i - 1], data[0] = data[0], data[length - i - 1]
		adjustOneNodeToHeap(data, 0, order, length - i - 1)
		printArray(data)
	}
	return &dtoAlgorithm.SortRes{Ret:data}, nil
}

//堆排序基准操作之一--首先构建一个堆
func genHeap(data []int, order int)  {
	//找到第一个非叶子节点下标
	length := len(data)
	pivotIdx :=  int(length / 2) - 1

	//从这个叶子节点下标到下标为0的根节点都调整完就构造了一个堆
	for i := pivotIdx; i>= 0; i-- {
		//针对每个节点调整成堆
		adjustOneNodeToHeap(data, i, order, length)
	}
}

//堆排序基准操作之二--从该节点往下调整成堆，实际调用该方法时是从低层到高层开始调整，所以调整后能保持整个结构
//length参数是在最终堆排序时需要交换堆顶和当前长度的最后一位交换，交换后不需要再判断了
func adjustOneNodeToHeap(data []int, adjustNodeIdx, order, length int)  {
	//用一个变量保留adjustNodeIndex的原有值
	//能根据adjustNodeIdx是否有变化来判断是否有交换操作

	for ;true; {
		Lchild := 2 * adjustNodeIdx + 1
		if Lchild >= length || Lchild < 0 {
			break
		}
		tempIdx := Lchild
		Rchild := Lchild + 1
		if Rchild < length && data[Lchild] < data[Rchild] {
			tempIdx = Rchild
		}

		if data[tempIdx] > data[adjustNodeIdx] { //交换
			data[tempIdx], data[adjustNodeIdx] = data[adjustNodeIdx], data[tempIdx]
		} else {
			break
		}
		adjustNodeIdx = tempIdx
	}
}

//4--冒泡排序
func BubbleSort(data []int, order int) (*dtoAlgorithm.SortRes, error) {
	length := len(data)
	if length <= 1 {
		return &dtoAlgorithm.SortRes{Ret:data}, nil
	}

	for i:=0; i < length - 1; i++ { //外层循环是指待排的一个数所在index,只要排了length-1个数，那么就都确定位置了
		for j:=i+1; j < length; j++ { //对每个位置和其后边的数据做对比，拿到最小的值放在这
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
				printArray(data)
			}
		}
	}
	return  &dtoAlgorithm.SortRes{Ret:data}, nil
}

//5--归并排序
func MergeSort(data []int, order int) (*dtoAlgorithm.SortRes, error) {
	length := len(data)
	low := 0
	high := length - 1;

	if low < high {
		temp := make([]int, length)
		mergesort(data, low, high, temp)
	}
	return &dtoAlgorithm.SortRes{Ret:data}, nil
}

//归并排序基本操作是对同一数组的相邻两个有序段合并成一个有序段，先用一个临时数组存储
//data[first,mid]data[mid+1,last]
//然后将排好的那段覆盖原数组的对应段即可,这样只要申请一次长度为n的数组即可
func MergeTwoSortedArray(data []int, first, mid, last int, temp []int)  {
	i := first;
	m := mid
	j := mid + 1;
	n := last

	k := 0
	for ; i<=m && j<=n; {
		if data[i] < data[j] {
			temp[k] = data[i]
			i++
		} else {
			temp[k] = data[j]
			j++
		}
		k++
	}

	//两段长度可能不一样
	if i <= m {
		for ;i <=m; i++ {
			temp[k] = data[i]
			k++
		}
	} else if j <= n {
		for ; j<= n; j++ {
			temp[k] = data[j]
			k++
		}
	}

	//将排序好的覆盖原来段,k正好就是长度
	printArray(temp)
	for t := 0; t < k; t++ {
		data[first+t] = temp[t]
	}
}

func mergesort(data []int, first, end int, temp []int)  {
	if first < end {
		mid := (first + end) / 2
		mergesort(data, first, mid, temp)
		mergesort(data, mid + 1, end, temp)
		MergeTwoSortedArray(data, first, mid, end, temp)
	}
}

func printArray(data []int)  {
	for i:=0; i<len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
	fmt.Println()
}