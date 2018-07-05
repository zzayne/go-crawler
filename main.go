package main

import (
	"fmt"
	"regexp"

	"github.com/zzayne/zayne-crawler/fetcher"
)

func main() {
	contents, err := fetcher.Fetch("https://www.douban.com/location/shenzhen/events/week-all")
	if err != nil {
		panic(err)
	}

	parser(contents)
}

func parser(contents []byte) {
	activityRe := `<span itemprop="summary">([^>]+)</span>`
	regx := regexp.MustCompile(activityRe)
	matches := regx.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("sumarry title:%s\n", string(m[1]))
	}
}
