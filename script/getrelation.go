/**
 * @File: kpstaff.go
 * @Author: zhuchengming
 * @Description: application/x-www-form-urlencoded形式的post请求
 * @Date: 2021/2/26 15:19
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var urlStr string = "http://kp.zuoyebang.cc/kpstaff/api/getrelation"

func testKpstaff()  {
	var params = map[string]interface{}{
		"staffUid": strconv.Itoa(2556824543),
		"studentUids":[]uint64{2191726571},
	}

	//获取请求参数
	bodyStr := PostUrlEncodedParam(params)
	fmt.Println(bodyStr)

	//请求
	request, _ := http.NewRequest("POST", urlStr, strings.NewReader(bodyStr))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := http.DefaultClient.Do(request)

	aa, _ := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	fmt.Println(string(aa))
}

func main()  {
	testKpstaff()
}

