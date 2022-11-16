package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"

	"os"
	"time"
)

func produce(ch chan struct{}, sysch chan os.Signal) {
	for {
		select {
		case <-sysch:
			sysch <- syscall.SIGINT
			fmt.Println("Produce exit")

			return
		default:
			ch <- struct{}{}
			time.Sleep(1 * time.Second)
			fmt.Println("This is produce")

		}

	}
}

func consume(ch chan struct{}, sysch chan os.Signal) {
	for {
		select {
		case <-sysch:
			sysch <- syscall.SIGINT
			fmt.Println("Consume exit")
			return
		default:
			<-ch
			time.Sleep(1 * time.Second)
			fmt.Println("This is consume")

		}
	}
}

func main() {
	//一个channel控制退出，一个控制生产消费
	ch := make(chan struct{})
	sysch := make(chan os.Signal, 1)
	go produce(ch, sysch)
	go consume(ch, sysch)
	//优雅退出
	signal.Notify(sysch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for sg := range sysch {
		switch v := sg.(type) {
		default:
			logrus.Infof("receive signal '%v', start graceful shutdown", v.String())
			sysch <- syscall.SIGINT
			time.Sleep(5 * time.Second)
			os.Exit(1)
		}

	}

}
