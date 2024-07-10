package htmltomd

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func readFile(fileName string) string {
	if bytes, err := os.ReadFile(fileName); err == nil {
		return string(bytes)
	}
	return ""
}

func TestParseHtml(t *testing.T) {
	html := readFile("./testdata/data_01.html")
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("NewDocumentFromReader err: %v", err)
	}
	document.Find("ul li:not(:nth-last-child(-n+2))").Each(func(i int, selection *goquery.Selection) {
		fmt.Printf("%v", selection.Text())
		t.Logf("%v", selection.Text())
	})
}

func TestParseHtml2(t *testing.T) {
	html := readFile("./testdata/data_03.html")
	htmlParser := NewHtmlToMarkdownParser()

	result, err := htmlParser.Parse(html)
	if err != nil {
		t.Fatalf("Parse err: %v", err)
	}
	t.Logf("result: %v", *result)
}
