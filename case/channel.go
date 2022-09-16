package main

import (
	"fmt"
	"os"
	"time"
)

var dog = make(chan struct{})
var cat = make(chan struct{})
var fish = make(chan struct{})
var count = 0

func stop() {
	if count == 100 {
		os.Exit(0)
	}
}

func Dog() {
	for {
		<-fish
		fmt.Println("dog")
		count++
		dog <- struct{}{}
		stop()
	}

}

func Cat() {
	for {
		<-dog
		fmt.Println("cat")
		count++
		cat <- struct{}{}
		stop()
	}
}

func Fish() {
	for {
		<-cat
		fmt.Println("fish")
		count++
		fish <- struct{}{}
		stop()
	}

}

func main() {
	go Dog()
	go Cat()
	go Fish()
	fish <- struct{}{}

	time.Sleep(10 * time.Second)
}
