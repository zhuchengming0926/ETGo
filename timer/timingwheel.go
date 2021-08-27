package timer

import (
	"log"
	"sync"
	"time"
)

const wheel_cnt uint8 = 5                                                                   //时间轮数量5个
var element_cnt_per_wheel = [wheel_cnt]uint32{256, 64, 64, 64, 64}                          //每个时间轮的槽(元素)数量。在 256+64+64+64+64 = 512 个槽中，表示的范围为 2^32
var right_shift_per_wheel = [wheel_cnt]uint32{8, 6, 6, 6, 6}                                //当指针指向当前时间轮最后一位数，再走一位就需要向上进位。每个时间轮进位的时候，使用右移的方式，最快实现进位。这里是每个轮的进位二进制位数
var base_per_wheel = [wheel_cnt]uint32{1, 256, 256 * 64, 256 * 64 * 64, 256 * 64 * 64 * 64} //记录每层时间轮的间隔时间单位(秒)
var mutex sync.Mutex                                                                        //加锁
var rwmutex sync.RWMutex
var newest [wheel_cnt]uint32                           //每个时间轮当前指针所指向的位置
var timewheels [5][]*Node                              //定义5个时间轮
var TimerMap map[string]*Node = make(map[string]*Node) //保存待执行的计时器，方便按链表节点指针地址直接删除定时器

type Timer struct {
	Name        string            //定时器名称
	Inteval     uint32            //时间间隔，即以插入该定时器的时间为起点，Inteval秒之后执行回调函数DoSomething()。例如进程插入该定时器的时间是2015-04-05 10:23:00，Inteval=5，则执行DoSomething()的时间就是2015-04-05 10:23:05。
	DoSomething func(interface{}) //自定义事件处理函数，需要触发的事件
	Args        interface{}       //上述函数的输入参数
}

func SetTimer(name string, inteval uint32, handler func(interface{}), args interface{}) {
	if inteval <= 0 {
		return
	}
	var bucket_no uint8 = 0
	var offset uint32 = inteval
	var left uint32 = inteval
	for offset >= element_cnt_per_wheel[bucket_no] { //偏移量大于当前时间轮容量，则需要向高位进位
		offset >>= right_shift_per_wheel[bucket_no] //计算高位的值。偏移量除以低位的进制。比如低位当前是256，则右移8个二进制位，就是除以256，得到的结果是高位的值。
		var tmp uint32 = 1
		if bucket_no == 0 {
			tmp = 0
		}
		left -= base_per_wheel[bucket_no] * (element_cnt_per_wheel[bucket_no] - newest[bucket_no] - tmp)
		bucket_no++
	}
	if offset < 1 {
		return
	}
	if inteval < base_per_wheel[bucket_no]*offset {
		return
	}
	left -= base_per_wheel[bucket_no] * (offset - 1)
	//left是赋值给具体timer，比如A的interval的，当A被插入到一个时间轮的具体位置后，从当前时间开始算起走到这个A所在格子，本就要经过一段时间
	//当走到后，因为高级时间轮一格表示一个时间段，也就是即使走到这个高级时间轮的A所在格子，也不能马上执行A，还要看它的interval是否为0，
	//是0才表示到它了，可以执行了，所以left表示一个timer任务插入到具体时间轮的某个格子后，当走到这个格子后，还需要经过多少秒才能执行
	pos := (newest[bucket_no] + offset) % element_cnt_per_wheel[bucket_no] //通过类似hash的方式，找到在时间轮上的插入位置

	var node Node
	node.SetData(Timer{name, left, handler, args})

	rwmutex.RLock()
	TimerMap[name] = timewheels[bucket_no][pos].InsertHead(node) //插入定时器
	rwmutex.RUnlock()
	//fmt.Println("pos ", bucket_no, pos, tmp)
}

func step() {
	//var dolist list.List
	{
		rwmutex.RLock()
		//遍历所有桶
		var bucket_no uint8 = 0
		for bucket_no = 0; bucket_no < wheel_cnt; bucket_no++ {
			newest[bucket_no] = (newest[bucket_no] + 1) % element_cnt_per_wheel[bucket_no] //当前指针递增1
			//fmt.Println(newest)
			var head *Node = timewheels[bucket_no][newest[bucket_no]] //返回当前指针指向的槽位置的表头
			var firstElement *Node = head.Next()
			for firstElement != nil { //链表不为空
				if value, ok := firstElement.Data().(Timer); ok { //如果element里面确实存储了Timer类型的数值，那么ok返回true，否则返回false。
					inteval := value.Inteval
					doSomething := value.DoSomething
					args := value.Args
					if nil != doSomething { //有遇到函数为nil的情况，所以这里判断下非nil
						if 0 == bucket_no || 0 == inteval {
							//dolist.PushBack(value) //执行自定义处理函数
							go doSomething(args)
						} else {
							SetTimer(value.Name, inteval, doSomething, args) //重新插入计时器
						}
					}
					Delete(firstElement) //删除定时器
				}
				firstElement = head.Next() //重新定位到链表第一个元素头
			}
			if 0 != newest[bucket_no] { //指针不是0，还未转回到原点，跳出。如果回到原点，则说明转完了一圈，需要向高位进位1，则继续循环入高位步进一步。
				break
			}
		}
		rwmutex.RUnlock()
	}
}

func Run() {
	var i int = 0
	for {
		go step()
		i++
		log.Printf("第%ds", i)
		//间隔时间inteval=1s
		time.Sleep(1 * time.Second)
	}
}

func init() { //初始化
	var bucket_no uint8 = 0
	for bucket_no = 0; bucket_no < wheel_cnt; bucket_no++ {
		var i uint32 = 0
		for ; i < element_cnt_per_wheel[bucket_no]; i++ {
			timewheels[bucket_no] = append(timewheels[bucket_no], new(Node))
		}
	}
}


/*
使用示例
package main

import (
	"log"
	"runtime"
	"timer_server/timer"
)

func callback1(args interface{}) {
	//只执行一次的事件
	if values, ok := args.([]string); ok {
		var str1 string = values[0]
		var str2 string = values[1]
		log.Println("callback1(" + str1 + "," + str2 + ")")
	} else {
		log.Println("callback1()")
	}
}

func callback2(args interface{}) {
	//每次在当前时间点之后5s插入一个定时器，这样就能形成每隔5秒调用一次callback2回调函数，可以用于周期性事件
	timer.SetTimer("callback2", 5, callback2, args)
	log.Println("callback2")
}

func main() {
	// cpu多核
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 定时器1，传入两个参数
	timer.SetTimer("callback1", 3, callback1, []string{"hello", "world"})
	// 移除定时器
	//timer.Delete(timer.TimerMap["callback2"])
	//运行计时器
	go timer.Run()
	// 定时器2，不传参数
	timer.SetTimer("callback2", 6, callback2, nil)
	select {
	}
}
*/