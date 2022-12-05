// package main

// import "fmt"

// // func main() {
// // 	name := "Go Developers"
// // 	fmt.Println("Azure for", name)
// // }

package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Extract(webUrl string, info_set map[string]bool) ([]string, map[string]bool) {
	// webUrl := "https://news.ycombinator.com"
	response, err := http.Get(webUrl)
	if err != nil {
		log.Fatalln(err)
		// handle error
	} else if response.StatusCode == 200 {
		// fmt.Println("Success!")
	} else {
		log.Fatalln("Failed")
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// document.Find("tr.athing")
	// document.Find("tr.athing").Each(func(index int, selector *goquery.Selection) {
	// 	/* Process selector here */
	// 	title := selector.Find("td.title").Text()
	// 	link, _ := selector.Find("a.titlelink").Attr("href") // blank identifier or found
	// 	// var link string
	// 	// link, _ = selector.Find("a.titlelink").Attr("href") // blank identifier

	// 	type Information struct {
	// 		link  string
	// 		title string
	// 	}
	// 	info := make([]Information, 0)

	// 	info = append(info, Information{
	// 		title: title,
	// 		link:  link,
	// 	})

	// 	// info := make([]string, 0)
	// 	// info = append(info, link)
	// 	fmt.Println(info)
	// })
	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "https")
		// In the predicate function : make sure that link has the https prefix.
	}
	info := make([]string, 0)
	// info_set := make(map[string]bool, 0) // implement a set
	document.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
		link, _ := tag.Attr("href")
		_, ok := info_set[link]
		if link != webUrl && !ok {
			info_set[link] = true
			info = append(info, link)
		}
		// linkText := tag.Text()
		// fmt.Printf("%s \n", link)

	})
	return info, info_set
}
