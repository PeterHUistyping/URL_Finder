// package main

// import "fmt"

// // func main() {
// // 	name := "Go Developers"
// // 	fmt.Println("Azure for", name)
// // }

package main

import (
	"fmt"
	"strings"
)

func main() {
	info_set := make(map[string]bool, 0) // implement a set
	info_new := make([]string, 0)
	url := make([]string, 0)
	url, info_set = Extract("https://news.ycombinator.com", info_set)
	count := len(url)
	//fmt.Printf("%v\n", url)
	for _, element := range url {
		//fmt.Println("At index", index, "value is", element)

		info_new, info_set = Extract(element, info_set)

		for _, url_new_element := range info_new {
			url = append(url, url_new_element)
			count++
			if count >= 100 {
				break
			}
		}
		if count >= 100 {
			break
		}
	}
	fmt.Println(strings.Join(url, "\n"))
}
