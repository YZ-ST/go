// main.go
package main

import (
	"fmt"
	"GO/medium/modules" // 匯入指令的模組
)

func main() {
	fmt.Println("Start crawling...")

	url := "https://hitripod.com/http-restful-api-with-golang/"

	// 呼叫模組的 Request() 函式
	res, err := crawler.Request(url)
	if err != nil {
		fmt.Println("Error in HTTP request:", err)
		return
	}

	// 呼叫模組的 ParseResponse() 函式
	body, err := crawler.ParseResponse(res)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	fmt.Println(body)
}
