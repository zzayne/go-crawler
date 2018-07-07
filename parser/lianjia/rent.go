package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/zzayne/go-crawler/engine"
	"github.com/zzayne/go-crawler/model"
)

//RentParser ...
func RentParser(doc *goquery.Document) (engine.ParseResult, error) {
	var result engine.ParseResult
	var ID, URL string

	var house model.House
	titleSec := doc.Find(".content-wrapper .title-wrapper .title")

	house.Name = titleSec.Find("main").Text()
	house.Merit = titleSec.Find("sub").Text()

	infoSec := doc.Find(".content-wrapper .overview")
	house.Img = infoSec.Find(".imgContainer img").AttrOr("src", "")

	house.Price = infoSec.Find(".price .total").Text() + infoSec.Find(".price .unit span").Text()
	roomStrList := doc.Find(".zf-room p")
	if roomStrList.Length() == 9 {
		house.Area = roomStrList.Eq(0).Text()
		house.Model = roomStrList.Eq(1).Text()
		house.Floor = roomStrList.Eq(2).Text()
		house.Toward = roomStrList.Eq(3).Text()
		house.Metro = roomStrList.Eq(4).Text()
		house.Plot = roomStrList.Eq(5).Find("a").Text()
		house.PlotURL = roomStrList.Eq(5).Find("a").AttrOr("href", "")
		if addr := roomStrList.Eq(6).Find("a"); addr.Length() == 2 {
			house.Region = addr.Eq(0).Text()
			house.Location = addr.Eq(1).Text()
		}
		house.Datetime = roomStrList.Eq(7).Text()
		house.Datetime = roomStrList.Eq(7).Text()
		house.Code = roomStrList.Eq(8).Text()
	}

	//匹配到的ID内容,链家编号：105101392982,处理后得到实际ID
	ID = doc.Find(".houseRecord .houseNum").Text()
	ID = strings.Replace(ID, "链家编号：", "", -1)
	URL = "https://sz.lianjia.com/zufang/" + ID + ".html"
	result.Items = append(result.Items, engine.Item{
		ID:      ID,
		URL:     URL,
		Type:    "rent",
		Payload: house,
	})

	return result, nil
}
