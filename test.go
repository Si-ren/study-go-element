package main

func main() {
	ch := make(chan int)
	go func(chan int) {
		ch <- 2
	}(ch)
	<-ch

}
