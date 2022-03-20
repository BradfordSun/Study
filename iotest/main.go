package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	bs := make([]byte, 2)
	for {
		// Read最多读len(p)长度的字节，而不是一定读这么长。它返回的n是实际读了多长。
		// 即使只读了一个字节，也不会返回io.EOF。只有彻底没得读了，n为0的时候，才返回io.EOF。
		n, err := f.Read(bs)
		fmt.Println(bs)
		fmt.Println(string(bs))
		fmt.Println(n, err)
		fmt.Println("--------------------------")
		if n == 0 || err == io.EOF {
			break
		}
	}
}
