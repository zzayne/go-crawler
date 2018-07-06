package main

import (
	"github.com/zzayne/zayne-crawler/engine"
	parser "github.com/zzayne/zayne-crawler/parser/lianjia"
)

func main() {

	e := engine.SimpleEngine{}
	// req := engine.Request{
	// 	URL:       "https://www.douban.com/location/china/",
	// 	ParseFunc: parser.AreaListParser,
	// }
	req := engine.Request{
		URL:       "https://sz.lianjia.com/zufang/pg1",
		ParseFunc: parser.RentParser,
	}

	e.Run(req)
}
