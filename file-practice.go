package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

//统计一个文本中各个字符数量
func main() {
	fileName := "e:\\go-file.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read file err:", err)
			break
		}
		//遍历str,进行统计
		for _, v := range str {
			switch {
			case v > 'a' && v < 'z':
				fallthrough //穿透，至下一个case
			case v > 'A' && v < 'Z':
				count.ChCount++
			case v == ' ' || v == '\n':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}
	fmt.Println(count)

}
