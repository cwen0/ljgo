package util

import (
	"html/template"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func ParseMarkdown(markdown string) template.HTML {
	return template.HTML(blackfriday.Run([]byte(markdown)))
}
