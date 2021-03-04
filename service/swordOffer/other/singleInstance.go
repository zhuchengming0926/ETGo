/**
 * @File: singleInstance.go
 * @Author: zhuchengming
 * @Description:单例模式
 * @Date: 2021/3/4 16:01
 */

package main

import "sync"

//懒汉模式：非线程安全。当正在创建时，有线程来访问此时ins = nil就会再创建，单例类就会有多个实例了
//type SingInstance struct {}
//var instance *SingInstance
//func GetInstance() *SingInstance {
//	if instance == nil {
//		instance = new(SingInstance)
//	}
//	return instance
//}

//饿汉模式:如果singleton创建初始化比较复杂耗时时，加载时间会延长。
//type SingInstance struct {}
//var instance = &SingInstance{}
//func GetInstance() *SingInstance {
//	return instance
//}

//懒汉加锁:每次都要判断，效率不高
//type Singleton struct {}
//var instance *Singleton
//var mu sync.Mutex
//
//func GetInstance() *Singleton {
//	mu.Lock()
//	defer mu.Unlock()
//	if instance == nil {
//		instance = &Singleton{}
//	}
//	return instance
//}

//双重锁，避免每次访问都要加锁
//type Singleton struct {}
//var instance *Singleton
//var mu sync.Mutex
//
//func GetInstance() *Singleton {
//	if instance == nil {
//		mu.Lock()
//		defer mu.Unlock()
//		if instance == nil {
//			instance = &Singleton{}
//		}
//	}
//	return instance
//}

//sync.Once实现
type Singleton struct {}
var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}