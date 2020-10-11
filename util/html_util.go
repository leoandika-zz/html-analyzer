package util

import (
	"HTMLAnalyzer/model"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"time"
)

func GetHtmlTitle(HTMLString string) (title string) {
	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data != "title" {
				continue
			}

			tt := z.Next()

			if tt == html.TextToken {
				t := z.Token()
				title = t.Data
				return
			}
		}
	}
}

func CountHeadingLevel(HTMLString string) (result model.HeadingStructure) {
	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data != "h1" && t.Data != "h2" && t.Data != "h3" && t.Data != "h4" && t.Data != "h5" && t.Data != "h6" {
				continue
			}

			switch t.Data {
			case "h1":
				result.H1Count++
			case "h2":
				result.H2Count++
			case "h3":
				result.H3Count++
			case "h4":
				result.H4Count++
			case "h5":
				result.H5Count++
			case "h6":
				result.H6Count++
			}
		}
	}
}

func CountLinks(HTMLString string, url string) (countInternal int64, countExternal int64, countInaccessible int64) {

	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken || tt == html.EndTagToken:
			t := z.Token()

			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						if strings.Contains(attr.Val, url) {
							countInternal++
						} else {
							countExternal++
						}
						//if the link need more than 1 second, it'll be considered inaccessible. Must be optimized.
						resp, err := client.Get(attr.Val)
						if err != nil || resp.StatusCode != http.StatusOK {
							countInaccessible++
						}
					}
				}
			}
		}
	}
}

func CheckLoginForm(HTMLString string) (isExist bool) {
	r := strings.NewReader(HTMLString)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()

			if t.Data == "input" {
				for _, attr := range t.Attr {
					if attr.Key == "type" {
						if strings.Contains(attr.Val, "password") {
							isExist = true
							return
						}
					}
				}
			}
		}
	}
}