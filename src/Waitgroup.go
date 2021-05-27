package main

import (
	"fmt"
"sync"
//"runtime"
)
var wg sync.WaitGroup

func main() {

	chnn := make(chan int)
	stop:= make (chan int)

	wg.Add(1)
	go func() {
		for i:=1;i<=100;i++ {
			fmt.Println("sending value:",i, "over channel")
			
			chnn<-i
		}
		fmt.Println("outside forloop of go routine 1")
		
		stop<-1
		
		wg.Done()
	}()
	
	
		for _:= range chnn{
			select {
				case <-stop : {fmt.Println("breaking out"); break}
				case channelvals:= <-chnn : fmt.Println("channel value:", channelvals)
				
			}
		}
close(chnn)
close(stop)
	wg.Wait()
	
}
