package model

//House ...
type House struct {
	//编码
	Code string
	Img  string
	Name string
	//卖点
	Merit string
	//房型
	Model string
	//楼层
	Floor string
	//朝向
	Toward string
	//地铁
	Metro string
	//小区
	Plot string
	//小区链接
	PlotURL string
	//发布时间
	Datetime string
	Price    string
	//区域，如宝安区
	Region string
	//位置 如市民中心
	Location string
	//面积
	Area string
}
