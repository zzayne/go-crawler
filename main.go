package main

import (
	"fmt"

	"github.com/zzayne/zayne-crawler/fetcher"
)

func main() {
	content, err := fetcher.Fetch("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
