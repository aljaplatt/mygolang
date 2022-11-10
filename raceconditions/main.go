package main

import (
	"fmt"
	"sync"
)

func main() {
	// we need goroutine management - sync package can do this 
	// waitgroup need to be passed on to functions
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	// check on score 
	var score = []int{0}

	//* opt 1
	wg.Add(3) // or you can do all three together
	//wg.Add(1) - to waitgroup 
	// these are go routine iifys - 
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done() // signal waitgroup is done 
	}(wg, mut)
	//*wg.Add(1)  opt 2 would be to do this above each wg 
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		fmt.Println(score)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// Once the waitgroups have been passed in,
	// wait until all the waitgroups have returned back / finished job
	wg.Wait()
	fmt.Println(score)
}
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	// we need goroutine management - sync package can do this 
// 	// waitgroup need to be passed on to functions
// 	wg := &sync.WaitGroup{}
// 	mut := &sync.RWMutex{}

// 	// check on score 
// 	var score = []int{0}

// 	//* opt 1
// 	wg.Add(3) // or you can do all three together
// 	//wg.Add(1) - to waitgroup 
// 	// these are go routine iifys - 
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("One R")
// 		mut.Lock()
// 		score = append(score, 1)
// 		mut.Unlock()
// 		wg.Done() // signal waitgroup is done 
// 	}(wg, mut)
// 	//*wg.Add(1)  opt 2 would be to do this above each wg 
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Two R")
// 		mut.Lock()
// 		score = append(score, 2)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Three R")
// 		mut.Lock()
// 		score = append(score, 3)
// 		mut.Unlock()
// 		wg.Done()
// 	}(wg, mut)
// 	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
// 		fmt.Println("Three R")
// 		mut.RLock()
// 		fmt.Println(score)
// 		mut.RUnlock()
// 		wg.Done()
// 	}(wg, mut)

// 	// Once the waitgroups have been passed in,
// 	// wait until all the waitgroups have returned back / finished job
// 	wg.Wait()
// 	fmt.Println(score)
// }

// lock resourse when you use it, not globally 