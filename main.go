package main

import (
	"github.com/zzayne/zayne-crawler/engine"
	parser "github.com/zzayne/zayne-crawler/parser/douban"
)

func main() {

	e := engine.SimpleEngine{}
	req := engine.Request{
		URL:       "https://www.douban.com/location/china/",
		ParseFunc: parser.AreaListParser,
	}
	
	e.Run(req)
}
