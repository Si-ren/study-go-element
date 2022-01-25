package main

import "fmt"

func main() {
	var court = make(chan int, 1)
	court <- 1
	//go func() {
	//	ball,ok := <-court
	//	fmt.Println(ball,ok)
	//}()
	ball, ok := <-court
	close(court)
	ball2, ok2 := <-court
	fmt.Println(ball, ok)
	fmt.Println(ball2, ok2)

}
