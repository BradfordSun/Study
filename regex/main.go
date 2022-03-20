package main

import (
	"fmt"
	"regexp"
)

const text = "My email is cesu@nwcdcloud.cn@阿迪斯厉害 123n"

const text2 = `
My email is cesu@nwcdcloud.cn@阿迪斯厉害 123n
email1 is 123zxc@nippqwe.org
email2 is                   asdnl129@xnckj.cn
email3 is ddd@adsa.com.cn
`

func main() {
	// 点匹配任何一个字符。+就是一个或多个。*是一个或多个
	// 点本身要用\来转义。但是\本身会被go语言认为是转义，所以要两个\
	//re := regexp.MustCompile(".+@.+\\..+")
	// 或者用``这样里面就不转义了
	// 用[a-zA-Z0-9]来表示任何数字或字母
	// 点在方括号里不用加\
	// 用()可以把我们想要的提取出来。得用FindAllStringSubmatch
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindString(text)
	fmt.Println(match)
	// 传-1代表所有的匹配
	match1 := re.FindAllString(text2, -1)
	fmt.Println(match1)
	match2 := re.FindAllStringSubmatch(text2, -1)
	fmt.Println(match2)

}
