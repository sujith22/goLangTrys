package main

import ("fmt"
"math/rand"
"sync"
"strconv"
"errors"
)

var intChannel = make(chan int)
var errorChannel = make(chan error)

const totalValues = 1000

/*** 
* create multiple go routine producers for two channels
* create 2 go routine consumer to read from the channels
*/


func main()  {
	mainWg := sync.WaitGroup{}
	mainWg.Add(2)
	go create_consumers(&mainWg)
	go create_producers(&mainWg)
	mainWg.Wait()
}

func create_consumers(mainWg *sync.WaitGroup) {
	defer mainWg.Done()

	wg := sync.WaitGroup{} 
	defer wg.Wait()
	
	// integer consumer
	wg.Add(1)
	go func (wg *sync.WaitGroup) {
		defer wg.Done()
		sum := 0
		for {
			val,ok1 := <- intChannel
			if ok1 {
				sum = sum + val
			} else {
				break;
			}
		}
		fmt.Println("FINAL SUM", sum)
	}(&wg)

	// error consumer
	wg.Add(1)
	go func (wg *sync.WaitGroup) {
		errors := []error{}
		defer wg.Done()
		for {
			val,ok1 := <- errorChannel
			if ok1 {
				errors = append(errors, val)
			} else {
				break;
			}
		}
		fmt.Println("FINAL ERRORS ", errors)
	}(&wg)
}

func create_producers(mainWg *sync.WaitGroup) {
	defer mainWg.Done()
	defer close(intChannel)
	defer close(errorChannel)

	wg := sync.WaitGroup{}
	defer wg.Wait()

	for i:=0; i<totalValues; i++ {
		wg.Add(1)
		go createRand(&wg)
	}
}

func createRand (wg *sync.WaitGroup) {
	defer wg.Done()
	i := rand.Intn(1000)
	fmt.Println("i value is ", i)
	if i%10 == 0 {
		str := errors.New(strconv.Itoa(i) + " is divisble by 10")
		errorChannel <- str
	} else {
		intChannel <- i
	}
}
