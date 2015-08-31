package main

import (
	"fmt"
	"time"
)

type Api interface {
	Request() int
}

type Google struct {
}

type Facebook struct {
}

type Twitter struct {
}

func (google *Google) Request() int {
	fmt.Println("google request start")
	time.Sleep(3 * time.Second)
	fmt.Println("google request end")
	return 200
}

func (fb *Facebook) Request() int {
	fmt.Println("facebook request start")
	time.Sleep(1 * time.Second)
	fmt.Println("facebook request end")
	return 404
}

func (tw *Twitter) Request() int {
	fmt.Println("twitter request start")
	time.Sleep(2 * time.Second)
	fmt.Println("twitter request end")
	return 500
}

func main() {
	fmt.Println("main started")
	apis := []Api{
		&Google{},
		&Facebook{},
		&Twitter{},
	}

	outputChannel := make(chan int)
	for _, api := range apis {
		go func(channel chan<- int, api Api) {
			channel <- api.Request()
		}(outputChannel, api)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("return status %d\n", <-outputChannel)
	}
	close(outputChannel)
}
