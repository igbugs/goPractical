package main

import (
	"github.com/PuerkitoBio/goquery"
	"go_crawler/douban_spider/model"
	"go_crawler/douban_spider/parse"
	"logging"
	"strings"
)

var baseUrl = "https://movie.douban.com/top250"

func add(movies []parse.DoubanMovie) {
	for index, movie := range movies {
		if err := model.DB.Create(&movie).Error; err != nil {
			logging.Error("db.Create index: %d, err: %v", index, err)
		}
	}
}

func start() {
	var movies []parse.DoubanMovie

	pages := parse.GetPages(baseUrl)
	for _, page := range pages {
		doc, err := goquery.NewDocument(strings.Join([]string{baseUrl, page.Url}, ""))
		if err != nil {
			logging.Error("Join Url failed, err: %v", err)
		}
		movies = append(movies, parse.ParseMovies(doc)...)
	}
	logging.Debug("movies: %v, %v", movies, len(movies))
	add(movies)
}

func main() {
	logging.Debug("start spider...")
	start()
	defer model.DB.Close()
}
