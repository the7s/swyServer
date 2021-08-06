package service

import (
	"github.com/PuerkitoBio/goquery"
	"swy-novel-server/app/model"
	"swy-novel-server/config"
	"swy-novel-server/library/utils"
)

type bookService struct{}

var Book = new(bookService)

func (bs bookService) GetBooks(webUrl string) []model.Book {

	var doc = utils.GetHtmlDoc(webUrl)

	var bl []model.Book

	// Find the review items
	doc.Find("#main #hotcontent div .item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		author := s.Find("dl dt span").Text()
		name := s.Find("dl dt a").Text()
		url, _ := s.Find("dl dt a").Attr("href")
		coverUrl, _ := s.Find(".image a img").Attr("src")
		desc := s.Find("dl dd").Text()

		var book = model.Book{
			Tag:      utils.SwyEncodeUrl(url),
			Name:     name,
			CoverUrl: config.Config.GetWebUrl() + coverUrl,
			Url:      config.Config.GetWebUrl() + url,
			Author:   author,
			Desc:     desc,
		}
		bl = append(bl, book)
	})
	return bl
}

func (bs bookService) GetBookDetail(webUrl string) model.BookDetail {

	var doc = utils.GetHtmlDoc(webUrl)

	var pEle = doc.Find("#wrapper .box_con")

	name := pEle.Find("#maininfo #info h1").Text()

	author := pEle.Find("#maininfo #info p").Eq(0).Text()

	status := pEle.Find("#maininfo #info p").Eq(1).Text()

	lastUpdateAt := pEle.Find("#maininfo #info p").Eq(2).Text()

	newChapter := pEle.Find("#maininfo #info p").Eq(3).Find("a").Text()
	newChapterUrl, _ := pEle.Find("#maininfo #info p").Eq(3).Find("a").Attr("href")

	desc := pEle.Find("#maininfo #intro p").Eq(0).Text()
	coverUrl, _ := pEle.Find("#sidebar #fmimg img").Attr("src")

	var bookDetail = model.BookDetail{
		Name:          name,
		CoverUrl:      config.Config.GetWebUrl() + coverUrl,
		Author:        utils.SwyParseColon(author),
		Desc:          desc,
		Status:        utils.SwyParseColon(status),
		LastUpdateAt:  utils.SwyParseColon(lastUpdateAt),
		NewChapter:    newChapter,
		NewChapterUrl: utils.SwyEncodeUrl(newChapterUrl),
	}
	return bookDetail
}