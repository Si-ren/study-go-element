package main

import (
	"context"
	"fmt"
)

// context 一次只能附带一个值

func main() {

	// c1 := context.Background()
	c1 := context.WithValue(context.Background(), "key1", "value1")
	c2 := context.WithValue(c1, "key2", "value2")
	c3 := context.WithValue(c2, "key3", "value3")
	c4 := context.WithValue(c2, "key4", "value4")
	fmt.Println(c1.Value("key1"), c1.Value("key11"), c1.Value("key2"), c1.Value("key3"))
	fmt.Println(c2.Value("key1"), c2.Value("key2"), c2.Value("key3"))
	fmt.Println(c3.Value("key1"), c3.Value("key2"), c3.Value("key3"))
	fmt.Println(c4.Value("key1"), c4.Value("key2"), c4.Value("key3"), c4.Value("key4"))
	/*
		value1 <nil> <nil>
		value1 value2 <nil>
		value1 value2 value3
		value1 value2 <nil> value4
	*/
}
