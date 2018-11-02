package main

import (
	"go_crawler/car_prices/spiders"
	"go_crawler/car_prices/downloader"
	"github.com/PuerkitoBio/goquery"
	"logging"
	"go_crawler/car_prices/scheduler"
	"fmt"
	"time"
	"go_crawler/car_prices/model"
)

var (
	StartUrl = "/2sc/%s/a0_0msdgscncgpi1ltocsp1exb4/"
	BaseUrl = "https://car.autohome.com.cn"

	maxPage = 99
	cars []spiders.QcCar
)

func Start(url string, ch chan []spiders.QcCar)  {
	body := downloader.Get(BaseUrl + url)
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		logging.Error("Downloader.Get err:%v", err)
	}

	currentPage := spiders.GetCurrentPage(doc)
	nextPageUrl, _ := spiders.GetNextPageUrl(doc)

	if currentPage > 0 && currentPage <= maxPage {
		cars := spiders.GetCar(doc)
		logging.Debug("get cars list: %v", cars)
		// 获取 当前页面的车的信息
		ch <- cars
		if url := nextPageUrl; url != "" {
			// 当前城市的 下一页的 URL 信息 append 到 URLs  切片，select 监听看 ch ,上面的发送的当前cars 的信息，会触发
			// 进入 下一页面的 Start() 函数，而此时 传入的 url 为刚刚append 进URLs 切片的 下一页的 url地址
			scheduler.AppendUrl(url)
		}

		logging.Debug("get next url: %v", url)
	} else {
		logging.Debug("haved max page!!")
	}
}

func main()  {
	citys := spiders.GetCitys()
	for _, v := range citys {
		scheduler.AppendUrl(fmt.Sprintf(StartUrl, v.Pinyin))
	}

	start := time.Now()
	delayTime := time.Second * 6

	ch := make(chan []spiders.QcCar)

	LOOP:
		for {
			if url := scheduler.PopUrl(); url != "" {
				// 进入最后的一个 URLs 这个 切片的城市，获取车的信息
				go Start(url, ch)
			}

			select {
			case result := <- ch:
				//cars = append(cars, result...)
				model.AddCars(result)
				go Start(scheduler.PopUrl(), ch)
			case <- time.After(delayTime):
				logging.Warn("Timeout from channel..")
				break LOOP
			}
		}

	//if len(cars) > 0 {
	//	model.AddCars(cars)
	//}

	logging.Debug("Time: %s", time.Since(start) - delayTime)
}


