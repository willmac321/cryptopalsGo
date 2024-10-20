package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("running challenges")
}

func threadExample() {
	log.Println("Starting...")
	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
		doStuff()
	}()

	wg.Add(1)
	// this goroutine could be moved into the other worker and done via a timestamp calc on that loop, but wanted to look at signalling for goroutines
	go func() {
		defer wg.Done()
		log.Println("thread start")
		time.Sleep(1 * time.Second)
		// signal done to other goroutine
		done <- true
		log.Println("thread stop")
	}()

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done

	log.Println("Process completed")
}
