# go-crawler 并发式网站爬虫

## 2022-04-09 更新

最近整理 repo 时，发现这个项目还有人 star，作为当初的练手项目，现在来看这个项目还是很不成熟。先升级到 go mod，有空再尝试整理维护。

## 介绍

作为并发编程练手项目，这个Repo是在慕课网学完[Google 资深工程师深度讲解 Go 语言](https://coding.imooc.com/class/180.html)后（个人觉得课程很不错，适合有语法基础的初学者进一步学习），自己尝试实践完成的爬虫实战练习。

## 目的

完成一个通用的爬虫基础框架，熟悉golang并发编程的基础，逐步深入golang分布式开发学习。


## 思路

原课程讲解了一个分布式的爬虫项目实践过程，十分有参考价值。所以本项目的实现也大体参照该课程的结构，有以下关键组成：

- [x] Engine，爬虫任务下发和结果处理
- [x] Scheduler，任务队列调度管理
- [x] Fetcher, 获取请求路径的html内容
- [x] Model，保存内容的实体
- [x] Parser，fetcher得到html解析器，解析后返回数据和下一步请求信息。
- [x] Persist，爬取信息保存实现

除了使用`Parser`模块来处理特定页面的信息解析规则外，其他模块都是抽象的通用模块，所以如果想爬取其他网站，只需要在`Parser`里添加新的解析规则，其他的模块处理逻辑不变。


## 依赖

该项目依赖以下组件和软件

- [github.com/PuerkitoBio/goquery](github.com/PuerkitoBio/goquery),提供类似`jquery`的`dom`元素选择方法，提高爬取内容的解析效率，相比原课程的正则匹配来提取内容，这种实现更加清晰优雅
- [gopkg.in/olivere/elastic.v5](https://gopkg.in/olivere/elastic.v5),`Elasticsearch`是一个分布式 RESTful 搜索和分析引擎,因为个人理解不深，不做过多介绍，项目里是做爬取结果内容存储和提供RESTful接口读取使用
- docker，这个不用多说，这里暂时用作运行`Elasticsearch`，相关命令`docker run -d -p 9200:9200  elasticsearch`


## 使用

暂时以爬取链家深圳租房信息为例，进行项目开发和调试，后期可根据实际需求来添加应对目标网站反爬虫的机制，或者提高爬取速度的分布式开发。

在`$GOPATH/src/github.com/zzayne`下存放该项目，可运行查看结果输出。如果`Elasticsearch`运行在`9200`默认端口，可`GET`请求`http://localhost:9200/house_info/rent/_search?`后得到json数据。

![image](https://note.youdao.com/yws/public/resource/d4aea1ddaafd92526a8e7ff70a3586ab/xmlnote/07B258CB79AB4E8A8EC854B4479F8EB0/4129)


## 其他

本项目仅为学习实践项目，他用产生一切后果与本人无关。我的邮箱：thezhangwen@outlook.com，欢迎一起交流学习。


