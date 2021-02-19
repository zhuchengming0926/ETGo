/**
 * @File: request.go
 * @Author: zhuchengming
 * @Description:
 * @Date: 2021/2/18 20:08
 */

package dtoAlgorithm

type SortReq struct {
	Data   []int `json:"data"`  //原始待排数据
	Type   int   `json:"type"`  //1，快速排序 2，二分查找，3，插入排序
	Order  int   `json:"order"` //排序，默认是0升序，1为降序
}

