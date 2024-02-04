package service

import (
	"fmt"
	"sync"
)


func ChannelRoutine() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go great(ch, &wg)
	wg.Wait()

	ng := <- ch
	fmt.Println(ng)


	_, ok := <- ch 
	if ok {
		fmt.Println("channel is open")
	}else {
		fmt.Println("channel is close")
	}
}

func great(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "Hello!!"
	close(ch)

}


func MultiRoutine() {	
	var wg sync.WaitGroup
	wg.Add(2)
	go helloWorld(&wg)
	go complete(&wg)
	wg.Wait()
}


func helloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello world!")
}


func complete(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("complete work")
}




