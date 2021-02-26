/**
 * @File: httpRequestParam.go
 * @Author: zhuchengming
 * @Description:http各类请求参数组装
 * @Date: 2021/2/26 16:32
 */

package main

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"net/url"
	"reflect"
)

//四种POST提交方式查看：https://imququ.com/post/four-ways-to-post-data-in-http.html

//各种请求参数获取方式
//http://www.360doc.com/content/19/0731/15/33093582_852188067.shtml
//https://blog.csdn.net/guyan0319/article/details/84674627

//get--------请求
func GetParams(requestParam interface{}) (encodeData string, err error) {
	v := url.Values{}
	if data, ok := requestParam.(map[string]string); ok {
		for key, value := range data {
			v.Add(key, value)
		}
	} else if data, ok := requestParam.(map[string]interface{}); ok {
		for key, value := range data {
			var vStr string
			switch value.(type) {
			case string:
				vStr = value.(string)
			default:
				if tmp, err := jsoniter.Marshal(value); err != nil {
					return encodeData, err
				} else {
					vStr = string(tmp)
				}
			}
			v.Add(key, vStr)
		}
	} else {
		return encodeData, errors.New("unSupport RequestBody type")
	}
	encodeData, err = v.Encode(), nil
	return  encodeData, err

	/*
	get参数使用，就是拼接在url后边
	urlData, _ := GetParams(params)
	finalURl := fmt.Sprintf("%s?%s", urlStr, urlData)

	//请求
	req, _ := http.NewRequest("GET", finalURl,nil)
	res, _ := http.DefaultClient.Do(req)

	//读取返回
	aa, _ := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()

	fmt.Println(string(aa))
	*/
}

//post-----请求 Content-Type:application/x-www-form-urlencoded
func PostUrlEncodedParam(requestParamMap map[string]interface{}) string {
	//先转换为map[string]string
	strMapParams := EncodeParamMap(requestParamMap)

	//组装参数方式一：
	//var r http.Request
	//r.ParseForm() //解析请求所有参数，放入r.Form中
	//for key, val := range strMapParams {
	//	r.Form.Add(key, val)
	//}
	//bodyStr := strings.TrimSpace(r.Form.Encode())
	//return bodyStr

	//组装参数方式二：
	bodyStr, _ := GetParams(strMapParams)
	return bodyStr

	/*
	POST参数使用：
	bodyStr := PostUrlEncodedParam(requestParam)
	request, _ := http.NewRequest("POST", urlStr, strings.NewReader(bodyStr))
	resp, _ := http.DefaultClient.Do(request)
	*/
}


// 将map[string]interface{}结构的param转换为map[string]string
func EncodeParamMap(maps map[string]interface{}) map[string]string {
	res := make(map[string]string)
	for k, v := range maps {
		if reflect.TypeOf(v).Kind() == reflect.String {
			value, ok := v.(string)
			if !ok {
				continue
			}
			if IsJsonMap(value) || IsJsonSlice(value) {
				str, err := jsoniter.Marshal(value)
				if err == nil {
					res[k] = string(str)
				}
			} else {
				res[k] = value
			}
		} else {
			str, err := jsoniter.Marshal(v)
			if err == nil {
				res[k] = string(str)
			}
		}
	}
	return res
}

func IsJsonString(s string) bool {
	var js string
	return jsoniter.Unmarshal([]byte(s), &js) == nil

}

func IsJsonMap(s string) bool {
	var js map[string]interface{}
	return jsoniter.Unmarshal([]byte(s), &js) == nil
}

func IsJsonSlice(s string) bool {
	var js []interface{}
	return jsoniter.Unmarshal([]byte(s), &js) == nil
}