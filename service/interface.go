package service

import "html-analyzer/model"

type Service interface {
	CheckHTMLFromURL(url string) (model.HTMLStructure, error)
}
