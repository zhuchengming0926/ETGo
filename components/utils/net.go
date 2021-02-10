/**************************************************************************
 *
 **************************************************************************/

/**
 * @File: net.go
 * @Author: zhuchengming@zuoyebang.com
 * @Description:
 * @Date: 2021/2/7 20:24
 */

package utils

import "net"

// 获取本机ip
func GetLocalIp() string {
	addrs, _ := net.InterfaceAddrs()
	var ip string = ""
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				if ip != "127.0.0.1" {
					return ip
				}
			}
		}
	}
	return "127.0.0.1"
}


