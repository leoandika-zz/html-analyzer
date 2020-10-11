package service

import "HTMLAnalyzer/model"

type Service interface {
	CheckHTMLFromURL(url string) (model.HTMLStructure, error)
}
