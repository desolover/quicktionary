package quicktionary

import (
	"bufio"
	"io"
	"strings"

	"github.com/desolover/quicktionary/speech"
)

type Word struct {
	Value         string // word
	PartType      string // speech part
	Template      *speech.Template
	Bases         map[string]string
	Transcription struct {
		Variants     []string
		AccentedChar int // index
	}
	Morphological struct {
		Root   string
		Suffix string
	}
}

func (p *Parser) getWord(s string) (*Word, error) {
	r := bufio.NewReader(strings.NewReader(s))
	word := Word{}
	var reminder, template string
	var err error
	for {
		if reminder, template, err = readTag(r, reminder); err != nil {
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
		template = trimTag(template)

		if word.PartType = getPartType(template); word.PartType == "" {
			continue
		}
		hasTemplate := false
		for _, t := range p.Templates {
			if strings.HasPrefix(template, t.Title) {
				word.Template = &t
				hasTemplate = true
				break
			}
		}
		if !hasTemplate {
			continue
		}
		word.Bases = GetWordBasesFromTag(template)
		// TODO is need another cycle-return for getting another tags
		return &word, nil
	}
}

func getPartType(s string) string {
	if !strings.Contains(s, "ru") {
		return ""
	}
	for _, speechPart := range speech.SpeechParts {
		if strings.HasPrefix(s, speechPart) {
			return speechPart
		}
	}
	return ""
}
