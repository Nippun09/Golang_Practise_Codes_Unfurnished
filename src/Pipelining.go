package main

import (
	"fmt"
	"sync"
)

func generator(gchan chan<- int) {

	for i:=31; i<40; i++{
		gchan<-i
	}
	fmt.Println("done with generator")
 wg.Done()
 close(gchan)
}

func squarer(gchan <-chan int, pchan chan<- int) {

	for v:= range gchan{
		pchan<- v*v
	}
	fmt.Println("done with squarer")
	wg.Done()
	close(pchan)
}

func printer(pchan <-chan int) {

	for v:= range pchan{
		fmt.Println("squared value:", v)
	}
	fmt.Println("done with printer")
	wg.Done()
}

var wg sync.WaitGroup

func main() {

	

	gchan:=make(chan int)
	pchan:=make(chan int)

	wg.Add(3)
	go generator(gchan)

	go squarer(gchan, pchan)

	go printer(pchan)
	wg.Wait()
	// fmt.Println("closing gchan")
	// close(gchan)
	// fmt.Println("closing chan")
	// close(pchan)
}