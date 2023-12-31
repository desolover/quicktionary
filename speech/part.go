package speech

import "strings"

var templateVariantsSplitters = []string{"<br />", "<br/>", "<br>", ", "}

func (t *Template) parseTemplate() {
	t.trimBrackets()

	ss := Split(t.Text, "|")
	if len(ss) == 0 {
		return
	}
	t.Parent = ss[0] // first element is the name of the parent-template

	for _, s := range ss {
		if strings.ContainsRune(s, '#') {
			// TODO isn't implemented yet
			continue
		}
		s = strings.TrimSpace(s)
		fields := strings.Split(s, "=")
		if len(fields) != 2 {
			continue
		}
		content := fields[1]
		if !(strings.Contains(content, "{{{") && strings.Contains(content, "}}}") && strings.Contains(content, "основа")) {
			continue
		}

		variant := Variant{
			Form: t.parseForm(strings.TrimSpace(fields[0])),
		}
		// TODO isn't completed:
		// "Он она оно" (2 лицо глаголов)
		if !variant.Form.Identified {
			continue
		}

		for _, splitter := range templateVariantsSplitters {
			if strings.Contains(content, splitter) {
				variant.Patterns = Split(content, splitter)
				// difficult cases like this:
				// {{{соотв-мн|[[{{{основа}}}ниматься|{{{основа}}}нима́ться]], [[{{{основа}}}ыматься|{{{основа}}}ыма́ться]]}}}
				break
			}
		}
		if variant.Patterns == nil {
			variant.Patterns = []string{content}
		}
		for i := range variant.Patterns {
			variant.Patterns[i] = clearExpectedFields(variant.Patterns[i])
			variant.Patterns[i] = strings.TrimSpace(variant.Patterns[i])
		}
		t.Variants = append(t.Variants, variant)

		//if strings.Contains(content, " ") {
		//	 continue
		//}
	}
}

func clearExpectedFields(s string) string {
	if !strings.Contains(s, "|") {
		return s
	}
	s = strings.Replace(s, "|{{{1}}}", "", 1)
	s = strings.Replace(s, "|{{{2}}}", "", 1)
	return strings.Replace(s, "|{{{3}}}", "", 1)
}

func (t *Template) trimBrackets() {
	// all page source
	// sometimes concluded in double-figure-brackets,
	// sometimes concluded in single
	if strings.HasPrefix(t.Text, "{{") {
		t.Text = strings.TrimPrefix(t.Text, "{{")
	} else if strings.HasPrefix(t.Text, "{") {
		t.Text = strings.TrimPrefix(t.Text, "{")
	}
	if strings.HasSuffix(t.Text, "}}") {
		t.Text = strings.TrimSuffix(t.Text, "}}")
	} else if strings.HasSuffix(t.Text, "}") {
		t.Text = strings.TrimSuffix(t.Text, "}")
	}
}
