package parse

import (
	"github.com/PuerkitoBio/goquery"
	"logging"
	"strconv"
	"strings"
	"regexp"
)

type DoubanMovie struct {
	Title    string
	Subtitle string
	Other    string
	Desc     string
	Year     string
	Area     string
	Tag      string
	Star     string
	Comment  string
	Quote    string
}

type Page struct {
	Page int
	Url  string
}

func GetPages(url string) []Page {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		logging.Error("NewDocument failed, err: %v", err)
	}

	return ParsePages(doc)
}

func ParsePages(doc *goquery.Document) (pages []Page) {
	pages = append(pages, Page{
		Page: 1,
		Url:  "",
	})

	doc.Find("#content > div > div.article > div.paginator > a").Each(
		func(i int, s *goquery.Selection) {
			page, _ := strconv.Atoi(s.Text())
			url, _ := s.Attr("href")

			pages = append(pages, Page{
				Page: page,
				Url:  url,
			})
		})

	return pages
}

func ParseMovies(doc *goquery.Document) (movies []DoubanMovie) {
	doc.Find("#content > div > div.article > ol > li").Each(
		func(i int, s *goquery.Selection) {
			title := s.Find(".hd a span").Eq(0).Text()

			subtitle := s.Find(".hd a span").Eq(1).Text()
			logging.Debug("before subtitle:'%s'",subtitle)
			subtitle = strings.TrimSpace(subtitle)
			subtitle = strings.TrimLeft(subtitle, "/")
			logging.Debug("after subtitle:'%s'",subtitle)

			other := s.Find(".hd a span").Eq(2).Text()
			logging.Debug("before other:'%s'",other)
			other = strings.TrimSpace(other)
			other = strings.TrimLeft(other, "/")
			logging.Debug("after other:'%s'",other)

			desc := strings.TrimSpace(s.Find(".bd p").Eq(0).Text())
			descInfo := strings.Split(desc, "\n")
			desc = descInfo[0]

			movieDesc := strings.Split(descInfo[1], "/")
			year := strings.TrimSpace(movieDesc[0])
			area := strings.TrimSpace(movieDesc[1])
			tag := strings.TrimSpace(movieDesc[2])

			star := s.Find(".bd .star .rating_num").Text()

			comment := strings.TrimSpace(s.Find(".bd .star span").Eq(3).Text())
			compile := regexp.MustCompile("[0-9]")
			comment = strings.Join(compile.FindAllString(comment, -1), "")

			quote := s.Find(".quote .inq").Text()

			movie := DoubanMovie{
				Title:    title,
				Subtitle: subtitle,
				Other:    other,
				Desc:     desc,
				Year:     year,
				Area:     area,
				Tag:      tag,
				Star:     star,
				Comment:  comment,
				Quote:    quote,
			}

			logging.Debug("i: %d, movie: %v", i, movie)
			movies = append(movies, movie)
		})
	return movies
}
