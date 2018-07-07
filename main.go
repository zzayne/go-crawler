package main

import (
	"strconv"

	parser "github.com/zzayne/go-crawler/parser/lianjia"
	"github.com/zzayne/go-crawler/engine"
)

func main() {

	e := engine.SimpleEngine{}

	var reqList []engine.Request

	//链家的租房无明显入口，每页内容也无好提取的下一页url，所以先手动提供120个页面入口来爬取内容
	for i := 1; i < 120; i++ {
		req := engine.Request{
			URL:       "https://sz.lianjia.com/zufang/pg" + strconv.Itoa(i),
			ParseFunc: parser.RentParser,
		}
		reqList = append(reqList, req)
	}

	e.Run(reqList...)
}
