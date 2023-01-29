package main

import (
	"fmt"
	"sync"
)

var data int
var mtx sync.Mutex

func updateData(wg *sync.WaitGroup, increment int) {
	defer wg.Done()
	mtx.Lock()
	data += increment
	mtx.Unlock()
}

func readData(wg *sync.WaitGroup) {
	defer wg.Done()
	mtx.Lock()
	fmt.Println("Data:", data)
	mtx.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go updateData(&wg, 10)
	go updateData(&wg, 20)
	go updateData(&wg, -5)
	go readData(&wg)
	go readData(&wg)
	wg.Wait()
	fmt.Println("Final Data:", data)
}
