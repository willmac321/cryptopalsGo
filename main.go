package main

import (
	challenge "cryptoGo/challenge/set1"
	utils "cryptoGo/util"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("running challenges")

	strings:=utils.ReadFile("./data/challenge8.txt")
		res := challenge.IdentifyAES128HexFromList(strings, 16)
		fmt.Printf("%x\n", res)
}

func threadExample() {
	log.Println("Starting...")
	var wg sync.WaitGroup
	done := make(chan bool)

	wg.Add(1)
	go func() {
		defer wg.Done()
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
