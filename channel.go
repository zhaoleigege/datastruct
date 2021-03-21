package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("here")
	wg := &sync.WaitGroup{}
	stopChan := make(chan struct{})
	resultChan := make(chan string)

	go func() {

		for i := 0; i < 5; i++ {
			(func(index int) {
				go consumer(index, resultChan, wg)
			})(i)
		}
		wg.Wait()
		close(stopChan)
	}()

	result := make([]string, 0)
Loop:
	for {
		select {
		case <-stopChan:
			break Loop
		default:

		}

		str, ok := <-resultChan
		if !ok {
			return
		}

		result = append(result, str)
	}

	close(resultChan)
	fmt.Println(result)
}

func consumer(index int, rc chan<- string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		wg.Done()
	}()

	time.Sleep(1 * time.Second)
	rc <- fmt.Sprintf("consumer-%d", index)
}
