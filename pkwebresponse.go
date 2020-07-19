package main

import ("net/http"
"fmt"
"time"
"sync"
)

var (wg sync.WaitGroup
total=make(map[string]int))
type Demo struct {
	ch chan time.Duration
	url string
totalduration time.Duration
amount int
}

func (demo *Demo)gethttp() (code int)  {
	defer wg.Done()
	defer func() {
		if err:=recover();err!=nil{fmt.Println("frank",demo.url,err)
	demo.ch<- 1*time.Millisecond
			demo.amount++}
	}()

	starttime:=time.Now()
	rsp,err:=http.Get(demo.url)
	if err!=nil{
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	endtime:=time.Since(starttime)
	demo.ch<-endtime
	demo.totalduration+=(endtime)
	demo.amount++
	return rsp.StatusCode
}


func Newdemo(url string)*Demo  {
	return &Demo{ch:make(chan time.Duration,Retriesamounts),url:url}
}
func main() {

b,s,bi:=Newdemo("https://www.baidu.com/"),Newdemo("https://www.sina.com/"),Newdemo("https://www.bilibili.com/")
	baidu, sina,bili := b.ch,s.ch,bi.ch
	for i := 0; i < Retriesamounts ; i++ {
		wg.Add(3)
		time.Sleep(1*time.Millisecond)


		go b.gethttp()
		go s.gethttp()
		go bi.gethttp()

		select {
		case c := <-baidu:
			fmt.Println("baidu received with code ", c)
			total["baidu"]++

		case c := <-sina:
			fmt.Println("sina received with code ", c)
			total["sina"]++

		case c := <-bili:
			fmt.Println("bili received with code ", c)
			total["bili"]++

		}


	}

wg.Wait()

//	select {
//
//}
/*
for循环遍历直到所有的goroutine都结束,不过最好用sync.WaitGroup来实现
*/
//for len(b.ch)+total["baidu"]!=Retriesamounts || len(bi.ch)+total["bili"]!=Retriesamounts ||len(s.ch)+total["sina"]!=Retriesamounts {
//time.Sleep(time.Millisecond)
//}

fmt.Println(len(b.ch),len(bi.ch),len(s.ch))

	fmt.Println(total)
	fmt.Println("sina :",s.totalduration,s.amount,"baidu ",b.totalduration,b.amount,"bili ",bi.totalduration,bi.amount)
}