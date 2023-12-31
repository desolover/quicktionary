package quicktionary

import (
	"bufio"
	"encoding/xml"
	"io"
	"sort"
	"strings"

	"github.com/desolover/quicktionary/speech"
)

type Parser struct {
	Templates []speech.Template
}

func NewParser(templates []speech.Template) *Parser {
	return &Parser{Templates: templates}
}
func ParseTemplatesSource(r io.Reader) ([]speech.Template, error) {
	reader := bufio.NewReader(r)

	var t []speech.Template
	for {
		value, err := readPage(reader)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var page Page
		if err := xml.Unmarshal([]byte(value), &page); err != nil {
			return nil, err
		}

		if !(strings.HasPrefix(page.Title, "Шаблон:") && (strings.Contains(page.Title, "ru") || strings.Contains(page.Title, "числ"))) {
			continue
		}
		isSpeechPartTemplate := false
		for _, part := range speech.SpeechParts {
			if strings.Contains(page.Title, part) {
				isSpeechPartTemplate = true
				break
			}
		}
		if !isSpeechPartTemplate {
			continue
		}

		t = append(t, speech.Template{
			Title: strings.TrimPrefix(page.Title, "Шаблон:"),
			ID:    page.ID,
			Text:  page.Revision.Text,
		})
	}

	sort.Slice(t, func(i, j int) bool { return t[i].Title < t[j].Title })
	return t, nil
}
