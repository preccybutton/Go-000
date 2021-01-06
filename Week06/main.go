package main

import (
	"github.com/robfig/cron"
	"net/http"
	"sync/atomic"
	"time"
)

type circulLinkeList struct{
	//count int64
	totalTime int64
	next *circulLinkeList
}

type HytrixWindows struct {
	length int
	num int
	maxTuntun int64
	head *circulLinkeList
	tail *circulLinkeList
}



//统计总请求数
var LiveTask int64

//元素是请求消耗的时间
var infoChan chan int64

var brokeflag int64	//熔断标志

func (hs *HytrixWindows) append(count int64) {
	if hs.num <= hs.length-1{
		if hs.num == 0{
			hs.tail = &circulLinkeList{totalTime:count}
			hs.head = hs.tail
		}else{
			hs.tail.next = &circulLinkeList{totalTime:count}
			hs.tail = hs.tail.next
			if hs.num == hs.length - 1{
				hs.tail.next = hs.head
			}
		}
		hs.num++
	}else if hs.num == hs.length{
		hs.head.totalTime = count
		hs.head, hs.tail = hs.head.next,hs.tail.next
		//hs.tail = hs.tail.next
	}
}

func (hs *HytrixWindows) countTime() {
	for tmp := range infoChan{
		d := hs.tail
		if d.totalTime == 0{
			atomic.StoreInt64(&LiveTask, 0)
		}
		if atomic.LoadInt64(&LiveTask) /1000 > hs.maxTuntun{	//判断吞吐量是否大于过去10秒内最大吞吐
			atomic.StoreInt64(&brokeflag, 1)
			go func() {
				time.Sleep(1*time.Second)
				atomic.StoreInt64(&brokeflag, 0)
			}()
		}else {
			d.totalTime += tmp
		}
	}
}

func main(){
	infoChan = make(chan int64, 100)
	hw := &HytrixWindows{length:10}
	go hw.countTime()	//统计该窗口内总共消耗时间
	initCron("* * * * *", hw)	//每秒钟统计过去10s
	http.HandleFunc("/", FirstHandlec)
	http.ListenAndServe("127.0.0.0:8000", nil)
}

func FirstHandlec (w http.ResponseWriter, r *http.Request){
	timeStart := time.Now().UnixNano()/ int64(time.Millisecond)
	defer func() {
		atomic.AddInt64(&LiveTask, -1)
		infoChan <- time.Now().UnixNano()/ int64(time.Millisecond) - timeStart
	}()
	atomic.AddInt64(&LiveTask,1)
	if atomic.LoadInt64(&brokeflag) == 1{	//熔断标志立即返回
		return
	}
	w.Write([]byte("first"))
}

func initCron(interval string, hystrix *HytrixWindows){
	var max int64
	cronJob := cron.New()
	cronJob.AddFunc(interval, func() {
		d := hystrix.head
		for i := 0; i < hystrix.length; i++{
			if d == nil{
				break
			}
			if max < d.totalTime{
				max = d.totalTime
			}
			d = d.next
		}
		hystrix.maxTuntun = max
		hystrix.append(0)
	})
	cronJob.Run()
}
