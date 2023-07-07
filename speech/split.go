package speech

import "strings"

type Tag struct {
	open, close string
}

// Some tags can include some another tags, example:
// Tag "{{" includes tag "{".
// So, included tags should be declared after including tags for correct processing.
// Like this:
var tags = []Tag{{open: "{{{", close: "}}}"}, {open: "{{", close: "}}"}, {open: "[[", close: "]]"}, {open: "{", close: "}"}}

func index(s string, sep string, closeTagID int) (tagID int, isOpen bool, index int, length int) {
	for i := range s {
		if strings.HasPrefix(s[i:], sep) {
			return -1, false, i, len(sep)
		} else if closeTagID != -1 && strings.HasPrefix(s[i:], tags[closeTagID].close) {
			return closeTagID, false, i, len(tags[closeTagID].close)
		}
		for tagID, tag := range tags {
			if strings.HasPrefix(s[i:], tag.open) {
				return tagID, true, i, len(tag.open)
			}
		}
	}
	return -1, false, -1, 0
}

// Split works like strings.Split, but considering escape-tags (separators inside them are ignored).
func Split(s string, sep string) []string {
	var tagIDs []int
	n := strings.Count(s, sep) + 1
	a := make([]string, n)
	n--
	i := 0
	for i < n {
		currentTagID := -1
		if len(tagIDs) > 0 {
			currentTagID = tagIDs[len(tagIDs)-1]
		}

		tagID, isOpen, pos, length := index(s, sep, currentTagID)
		if pos < 0 {
			break
		}
		if tagID >= 0 {
			if isOpen {
				// opening-tag was found
				tagIDs = append(tagIDs, tagID)
			} else if len(tagIDs) > 0 && tagIDs[len(tagIDs)-1] == tagID {
				// suitable closing-tag was found
				tagIDs = tagIDs[:len(tagIDs)-1]
			}
		}

		if len(tagIDs) == 0 && tagID == -1 {
			// separator was found
			a[i] += s[:pos]
			i++
		} else {
			// tag was found or separator was escaped
			a[i] += s[:pos+length]
		}
		s = s[pos+length:]

	}

	a[i] += s
	return a[:i+1]
}
