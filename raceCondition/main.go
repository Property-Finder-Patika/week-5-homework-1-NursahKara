package main

import (
	"fmt"
	"time"
)

type channelStructure struct {
	first  string
	second string
}

func main() {
	val := channelStructure{}
	c := make(chan channelStructure)
	go func1(c)
	go func2(c)

	func(c chan channelStructure) {
		//err := errors.New("Error!")
		for {
			select {
			case v, ok := <-c:
				if !ok {
					return
				}
				if v.first != "" {
					fmt.Printf("Received %v\n", v)
					val.first = v.first
				}
				if v.second != "" {
					fmt.Printf("Received %v\n", v)
					val.second = v.second
				}
				if val.first != "" && val.second != "" {
					close(c)
				}
			}
		}
	}(c)
	fmt.Printf("%v\n", val)
}

func func1(c chan<- channelStructure) {
	time.Sleep(1 * time.Second)
	c <- channelStructure{first: "Test 1"}
}

func func2(c chan<- channelStructure) {
	c <- channelStructure{second: "Test 2"}
}
