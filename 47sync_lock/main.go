package main

import (
	"sync"
	"fmt"
)

//死锁问题
var (
	x=0
	wg sync.WaitGroup
	lock sync.Mutex
)

func add(){
	for i:=0;i<5000 ;i++  {
		lock.Lock()//加锁
		x=x+1
		lock.Unlock()//解锁
	}
	wg.Done()
}

func main(){
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}