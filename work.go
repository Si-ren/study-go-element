package main

import (
	"github.com/goinaction/code/chapter7/patterns/work"
	"log"
	"time"
)

type namePrinter struct {
	name string
}

func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)
}
