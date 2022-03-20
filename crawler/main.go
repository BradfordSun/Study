package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 如果html网页的编码不是utf-8，可以使用charset.DetermineEncoding来判断下是什么类型，然后在main函数里用transform转换成utf8

//func determineEncoding(r io.Reader) encoding.Encoding {
//	// 用bufio装一下前1024个字节，用于判断类型，以免直接传给后面DetermineEncoding这1024字节就用不了了
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//	return e
//}

func main() {
	resp, err := http.Get("https://aws.amazon.com/cn/premiumsupport/knowledge-center/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}
