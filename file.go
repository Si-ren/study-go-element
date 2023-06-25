package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
	data, err := os.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	err = os.WriteFile(destFile, data, 0666)
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

// 写入文件的几种方法
func openf1(f string) {
	f1, err := os.Create(f)
	if err != nil {
		fmt.Println("Cannt create file: ", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string("test"))
}

func openf2(f string) {
	f2, err := os.Create(f)
	if err != nil {
		fmt.Println("Cannt create file: ", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString("test")
	fmt.Printf("wrote %d bytes\n", n)
}

func openf3(f string) {
	f3, err := os.Create(f)
	if err != nil {
		fmt.Println("Cannt create file: ", err)
		return
	}
	defer f3.Close()
	w := bufio.NewWriter(f3)
	n, err := w.WriteString("test")
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
}

func openf4(f string) {
	s := []byte("Test Date")
	err := os.WriteFile(f, s, 0644)
	if err != nil {
		fmt.Println("Cannt write to file: ", err)

	}

}

func openf5(f string) {
	s := []byte("Test Date")
	f5, err := os.Create(f)
	if err != nil {
		fmt.Println("Cannt create file: ", err)
		return
	}
	defer f5.Close()
	n, err := io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println("Cannt write date to file: ", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}

// strings.NewReader读入数据写入别处
func stringsNewReader() {
	s := strings.NewReader("Only Strings")
	fmt.Println("r length: ", s.Len())
	n, err := s.WriteTo(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Wrote %d bytes to os.Stderr\n", n)
}
