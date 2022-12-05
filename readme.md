# URL_Finder

## Requirement
Implement a very simple web crawler. It should accept a single starting URL, such as https://news.ycombinator.com, as its input. It should download the web page available at the input URL and extract the URLs of other pages linked to from the HTML source code. Although there are several types of link in HTML, just looking at the href attribute of tags will be sufficient for this task. 
It should then attempt to download each of those URLs in turn to find even more URLs, and then download those, and so on. The program should stop after it has discovered 100 unique URLs and print them (one URL per line) as its output. 

## Brief Introduction
Solving the problem in Python and Go

## Python
### Dependencies 
```
import requests
# pip install requests
from bs4 import BeautifulSoup
# Beautiful Soup(bs4) is a Python library for pulling data out of HTML and XML files. 
```

## Go
Using goquery
### Dependencies
Type it in the terminal:
```
go get github.com/PuerkitoBio/goquery
```
import
```
"github.com/PuerkitoBio/goquery"
```
