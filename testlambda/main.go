package main

import (
	"fmt"
	"os"
	"time"
)

var timeCheck = time.Tick(time.Minute)

func main() {
	fmt.Println("Hello World")
	indexName := "knowledge-center-" + time.Now().Format("2006-01-02-15-04-05")
	fmt.Println(indexName)
	<-timeCheck
	fmt.Println("123")
	os.Exit(0)
}
