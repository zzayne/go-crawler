package main

import (
	"strconv"

	"github.com/zzayne/zayne-crawler/engine"
	parser "github.com/zzayne/zayne-crawler/parser/lianjia"
)

func main() {

	e := engine.SimpleEngine{}
	// req := engine.Request{
	// 	URL:       "https://www.douban.com/location/china/",
	// 	ParseFunc: parser.AreaListParser,
	// }

	var reqList []engine.Request
	for i := 1; i < 120; i++ {
		req := engine.Request{
			URL:       "https://sz.lianjia.com/zufang/pg" + strconv.Itoa(i),
			ParseFunc: parser.RentParser,
		}
		reqList = append(reqList, req)
	}

	e.Run(reqList...)
}
