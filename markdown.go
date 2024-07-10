package htmltomd

type MarkdownContent string

const (
	DocRoot MarkdownContent = "body"
	H1      MarkdownContent = "H1"
	H2      MarkdownContent = "H2"
	H3      MarkdownContent = "H3"
	H4      MarkdownContent = "H4"
	H5      MarkdownContent = "H5"
	H6      MarkdownContent = "H6"
	Code    MarkdownContent = "code"
)

var defaultMarkdownContentConfig = map[MarkdownContent]string{
	DocRoot: "body",
	H1:      "h1",
	H2:      "h2",
	H3:      "h3",
	H4:      "h4",
	H5:      "h5",
	H6:      "h6",
}
