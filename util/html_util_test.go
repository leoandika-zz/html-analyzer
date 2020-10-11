package util

import (
	"HTMLAnalyzer/model"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetHtmlTitle_NoTitle_ReturnEmptyString(t *testing.T) {
	htmlString := "<html><head></head><body></body></html>"
	expected := ""

	actual := GetHtmlTitle(htmlString)

	assert.Equal(t, expected, actual)
}

func TestGetHtmlTitle_TitleExist_ReturnTitle(t *testing.T) {
	htmlString := "<html><head><title>Unit Testing HTML</title></head><body></body></html>"
	expected := "Unit Testing HTML"

	actual := GetHtmlTitle(htmlString)

	assert.Equal(t, expected, actual)
}

func TestCountHeadingLevel_NoHeadingAtAll_ReturnZeroHeadingData(t *testing.T) {
	htmlString := "<html><head></head><body></body></html>"
	expected := model.HeadingStructure{}

	actual := CountHeadingLevel(htmlString)

	reflect.DeepEqual(expected, actual)
}

func TestCountHeadingLevel_HTMLWithSomeHeading_ReturnCorrectHeadingCount(t *testing.T) {
	htmlString := "<html>" +
		"<head></head>" +
		"<body>" +
		"<h1>this is h1</h1>" +
		"<h2>this is h2</h2>" +
		"<h3>this is h3</h3>" +
		"<h4>this is h3</h4>" +
		"<h4>this is h3</h4>" +
		"<h5>this is h3</h5>" +
		"<h6>this is h3</h6>" +
		"</body>" +
		"</html>"
	expected := model.HeadingStructure{
		H1Count: 1,
		H2Count: 1,
		H3Count: 1,
		H4Count: 2,
		H5Count: 1,
		H6Count: 1,
	}

	actual := CountHeadingLevel(htmlString)

	reflect.DeepEqual(expected, actual)
}

func TestCountLinks_NoLinkAtAll_ReturnZeroValue(t *testing.T) {
	url := "www.unittest.com"
	htmlString := "<html><head></head><body></body></html>"
	expectedInternal := int64(0)
	expectedExternal := int64(0)
	expectedInaccessible := int64(0)

	actualInternal, actualExternal, actualInaccessible := CountLinks(htmlString, url)

	assert.Equal(t, expectedExternal, actualExternal)
	assert.Equal(t, expectedInternal, actualInternal)
	assert.Equal(t, expectedInaccessible, actualInaccessible)
}

func TestCountLinks_SomeExternalAndInternalLink_ReturnCorrectValue(t *testing.T) {
	url := "google.com"
	htmlString := "<html>" +
		"<head></head>" +
		"<body>" +
		"<a href=\"https://web.whatsapp.com\"></a>" +
		"<a href=\"https://image.google.com\"></a>" +
		"<a href=\"http://notawebpage.com\"></a>" +
		"</body>" +
		"</html>"
	expectedInternal := int64(1)
	expectedExternal := int64(2)
	expectedInaccessible := int64(1)

	actualInternal, actualExternal, actualInaccessible := CountLinks(htmlString, url)

	assert.Equal(t, expectedExternal, actualExternal)
	assert.Equal(t, expectedInternal, actualInternal)
	assert.Equal(t, expectedInaccessible, actualInaccessible)
}

func TestCheckLoginForm_NoLoginFormOnPage_ShouldReturnFalse(t *testing.T) {
	htmlString := "<html>" +
		"<head></head>" +
		"<body></body>" +
		"</html>"
	expected := false

	actual := CheckLoginForm(htmlString)

	assert.Equal(t, expected, actual)
}

func TestCheckLoginForm_LoginFormExistOnPage_ShouldReturnTrue(t *testing.T) {
	htmlString := "<html>" +
		"<head></head>" +
		"<body>" +
		"<form>" +
		"<input type=\"text\" class=\"username\">" +
		"<input type=\"password\" class=\"password\">" +
		"</form>" +
		"</body>" +
		"</html>"
	expected := true

	actual := CheckLoginForm(htmlString)

	assert.Equal(t, expected, actual)
}