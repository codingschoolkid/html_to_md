package htmltomd

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type ParseOption func(*HtmlToMarkdownParser)
type ContentHandler func(string) string

type HtmlToMarkdownParser struct {
	contentHandler ContentHandler // 内容处理器，转换为markdown前对内容进行处理
	markdownConfig map[MarkdownContent]string
}

func WithParseMarkdownConfig(config map[MarkdownContent]string) ParseOption {
	return func(p *HtmlToMarkdownParser) {
		p.markdownConfig = config
	}
}

func WithContentHandler(contentHandler ContentHandler) ParseOption {
	return func(p *HtmlToMarkdownParser) {
		p.contentHandler = contentHandler
	}
}

func NewHtmlToMarkdownParser(options ...ParseOption) *HtmlToMarkdownParser {
	parser := &HtmlToMarkdownParser{
		markdownConfig: defaultMarkdownContentConfig,
	}
	for _, option := range options {
		option(parser)
	}
	return parser
}

func (p *HtmlToMarkdownParser) Parse(htmlContent string) (parseResult *string, err error) {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return
	}

	section := document.Find(p.markdownConfig[DocRoot])
	log.Printf("section: %v", section)

	return p.parseHelper(section)
}

func (p *HtmlToMarkdownParser) parseHelper(selection *goquery.Selection) (parseResult *string, err error) {
	stack := make([]*goquery.Selection, 10000)
	top := -1

	pushStack := func(selection *goquery.Selection) {
		top = top + 1
		if len(stack) < top {
			stack = append(stack, selection)
		} else {
			stack[top] = selection
		}
	}
	popStack := func() *goquery.Selection {
		result := stack[top]
		top = top - 1
		return result
	}
	isEmptyStack := func() bool {
		return top == -1
	}

	pushStack(selection)
	for !isEmptyStack() {
		selection = popStack()

		selection.Each(func(i int, child *goquery.Selection) {
			node := selection.Nodes[0]

			
			
			if node.Type == html.ElementNode {
				pushStack(child)
			} else if node.Type == html.TextNode {
				log.Printf("child.Text(): %v", child.Text())
				*parseResult = *parseResult + child.Text()
			}
		})
	}

	return
}
