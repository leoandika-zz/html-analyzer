package service

import (
	"HTMLAnalyzer/model"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHTMLAnalyzerService_CheckHTMLFromURL_EmptyHTMLBody(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`<html></html>`))
	}))
	defer mockServer.Close()
	expected := model.HTMLStructure{
		HTMLVersion: "<html></html>",
	}

	service := NewHTMLAnalyzerService(mockServer.Client())

	actual, err := service.CheckHTMLFromURL(mockServer.URL)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("not equal. got expected : %v, actual %v", expected, actual)
	}
}

func TestHTMLAnalyzerService_CheckHTMLFromURL_ReturnCorrectHTMLAnalysis(t *testing.T) {
	htmlBody := "<html><head><title>This is a title</title></head><body>" +
		"<h1>heading</h1><h2>heading2</h2><h2>heading2</h2><h3>heading3</h3>" +
		"<h4>heading4</h4>" +
		"<h4>heading4</h4>" +
		"<h4>heading4</h4>" +
		"<h5>heading5</h5>" +
		"<h5>heading5</h5>" +
		"<a href=\"notawebpage.com\"></a>" +
		"<form>" +
		"<input type=\"text\" class=\"username\">" +
		"<input type=\"password\" class=\"password\">" +
		"</form>" +
		"</body></html>"
	mockServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(htmlBody))
	}))
	defer mockServer.Close()
	expected := model.HTMLStructure{
		HTMLVersion: "<html><head><title>This is a title</title></head><body>" +
			"<h1>heading</h1><h2>heading2</h2><h2>heading2</h2><h3>heading3</h3>" +
			"<h4>heading4</h4>" +
			"<h4>heading4</h4>" +
			"<h4>heading4</h4>" +
			"<h5>heading5</h5>" +
			"<h5>heading5</h5>" +
			"<a href=\"notawebpage.com\"></a>" +
			"<form>" +
			"<input type=\"text\" class=\"username\">" +
			"<input type=\"password\" class=\"password\">" +
			"</form>" +
			"</body></html>",
		PageTitle: "This is a title",
		HeadingCount: model.HeadingStructure{
			H1Count: 1,
			H2Count: 2,
			H3Count: 1,
			H4Count: 3,
			H5Count: 2,
			H6Count: 0,
		},
		InternalLinkCount:     0,
		ExternalLinkCount:     1,
		InaccessibleLinkCount: 1,
		LoginFormExist:        true,
	}

	service := NewHTMLAnalyzerService(mockServer.Client())

	actual, err := service.CheckHTMLFromURL(mockServer.URL)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("not equal. got expected : %v, actual %v", expected, actual)
	}
}
