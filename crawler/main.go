package main

import (
	"crawler/engine"
	"crawler/knowledgecenter/parser"
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
	engine.Run(engine.Request{
		Url:        "https://aws.amazon.com/cn/premiumsupport/knowledge-center/",
		Parserfunc: parser.ParseUrlList,
	})
	//resp, err := http.Get("https://aws.amazon.com/cn/premiumsupport/knowledge-center/")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code ", resp.StatusCode)
	//	return
	//}
	//
	////e := determineEncoding(resp.Body)
	////utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	////all, err := ioutil.ReadAll(utf8Reader)
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}
	//printArticleList(all)
	////fmt.Printf("%s\n", all)
}

// 点不用转义
//func printArticleList(contents []byte) {
//	re := regexp.MustCompile(`<a href="(/cn/premiumsupport/knowledge-center/[^"]+)"[^>]*>([^<]+)</a>`)
//	matches := re.FindAllSubmatch(contents, -1)
//	for _, m := range matches {
//		fmt.Printf("Title: %s, URL: https://aws.amazon.com%s\n", m[2], m[1])
//	}
//	fmt.Println(len(matches))
//}
