package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//	打开文件
	file, err := os.Open("e:/go-file.txt")
	defer file.Close()

	if err != nil {
		fmt.Println("open file error", err)
	}

	//	输出文件
	fmt.Printf("file=%v\n", file)

	//创建一个 *Reader,默认带缓冲，缓冲大小为4096
	reader := bufio.NewReader(file)
	//循环读取问价内容
	for {
		str, err := reader.ReadString('\n')
		//读到eof（end of file）就报错
		if err == io.EOF {
			break
		}
		fmt.Printf("%v", str)
	}

	//ioutil.ReadFile一次性把整个文件读入内存
	file01 := "e:\\go-file.txt"
	readFile, err := ioutil.ReadFile(file01)
	if err != nil {
		fmt.Println("read file err：", err)
		return
	}
	//要用字符串输出，不然ioutil读入的都是字节，即输出为：123,546等
	fmt.Printf("%s", readFile)

	//使用OpenFile创建文件并写入
	//const (
	//	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	//	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	//	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	//	// The remaining values may be or'ed in to control behavior.
	//	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	//	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	//	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	//	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	//	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	//)
	fileWrite := "e:\\test-gofile.txt"
	//如果文件存在，使用O_CREATE会报错
	filewrite, err := os.OpenFile(fileWrite, os.O_CREATE|os.O_WRONLY, 0666)
	defer filewrite.Close()
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	str := "hello Siri!!!\n"
	writer := bufio.NewWriter(filewrite)
	writer.WriteString(str)
	//由于writer是带缓存的，因此WriteString方法是先写入内容中的
	//所以需要调用Flush方法，将内存中的数据真正写入文件中
	//否则文件中无数据
	writer.Flush()

	fmt.Println()

	//copy文件
	sourceFile := "e:\\go-file.txt"
	var destFile string = "e:\\to-File.txt"
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	err = ioutil.WriteFile(destFile, data, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

}

// PathExists 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
