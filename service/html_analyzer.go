package service

import (
	log "github.com/sirupsen/logrus"
	"html-analyzer/model"
	"html-analyzer/util"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type HTMLAnalyzerService struct {
	Client *http.Client
}

func NewHTMLAnalyzerService(client *http.Client) Service {
	return &HTMLAnalyzerService{
		Client: client,
	}
}

func (H *HTMLAnalyzerService) CheckHTMLFromURL(url string) (model.HTMLStructure, error) {
	var result model.HTMLStructure
	var title string
	var headingCount model.HeadingStructure
	var internalCount, externalCount, inaccessibleCount int64
	var isLoginFormExist bool

	var wg sync.WaitGroup
	wg.Add(4)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Errorf("error in creating http request. err : %s", err.Error())
		return result, err
	}

	resp, err := H.Client.Do(request)
	if err != nil {
		log.Errorf("error in make HTTP Call (GET) to %s. err : %s", url, err.Error())
		return result, err
	}
	defer resp.Body.Close()

	//Get htmlVersion body
	htmlByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("fail to read response body")
		return result, err
	}

	htmlVersion := string(htmlByte)

	go func() {
		//Get title
		title = util.GetHtmlTitle(htmlVersion)
		wg.Done()
	}()

	go func() {
		//Get Heading Count
		headingCount = util.CountHeadingLevel(htmlVersion)
		wg.Done()
	}()

	go func() {
		//Get all kind of links
		//modify url to only contain the last url without protocol and www
		//e.g https://www.google.com will be changed to google.com, therefore images.google.com also considered internal link
		modifiedUrl := strings.Replace(url, "https://www.", "", -1)
		modifiedUrl = strings.Replace(modifiedUrl, "http://www.", "", -1)
		internalCount, externalCount, inaccessibleCount = util.CountLinks(htmlVersion, modifiedUrl)
		wg.Done()
	}()

	go func() {
		//Check if there's login form
		isLoginFormExist = util.CheckLoginForm(htmlVersion)
		wg.Done()
	}()

	wg.Wait()

	result = model.HTMLStructure{
		HTMLVersion:           htmlVersion,
		PageTitle:             title,
		HeadingCount:          headingCount,
		InternalLinkCount:     internalCount,
		ExternalLinkCount:     externalCount,
		InaccessibleLinkCount: inaccessibleCount,
		LoginFormExist:        isLoginFormExist,
	}

	return result, nil
}
