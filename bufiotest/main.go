package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "test.txt"
	f, _ := os.Open(filename)
	defer f.Close()
	b1 := bufio.NewReader(f)
	// bufio.NewReader其实就是去调NewReaderSize(rd, defaultBufSize) 而这个默认的buffer是4096
	// 那么如果p的长度超过了4096，这个缓冲区就没用了。因为一次拿的值比缓冲区都大，不用频繁的调硬盘的io
	// 如果p的长度不到4096，那么这个缓冲器就有意义。因为数据先存到缓冲区，再从缓冲区拿，就减少了硬盘io
	p := make([]byte, 1024)
	n1, _ := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))
}
