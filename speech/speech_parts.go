package speech

import "strings"

const (
	SpeechPartNoun         = "сущ"
	SpeechPartVerb         = "гл"
	SpeechPartAdjective    = "прил"
	SpeechPartPronoun      = "мест"
	SpeechPartNumeral      = "числ"
	SpeechPartParticiple   = "прич"
	SpeechPartPreposition  = "prep"
	SpeechPartConjunction  = "conj"
	SpeechPartAdverb       = "adv"
	SpeechPartParticle     = "part"
	SpeechPartInterjection = "interj"
	SpeechPartSurname      = "Фам"
)

const Prefix = "ru"

var SpeechParts = []string{SpeechPartNoun, SpeechPartVerb, SpeechPartAdjective, SpeechPartPronoun, SpeechPartNumeral, SpeechPartParticiple,
	SpeechPartPreposition, SpeechPartConjunction, SpeechPartAdverb, SpeechPartParticle, SpeechPartInterjection, SpeechPartSurname}

var (
	speechParts   = []string{SpeechPartNumeral, SpeechPartParticiple, SpeechPartSurname}
	ruSpeechParts = []string{SpeechPartVerb, SpeechPartAdjective, SpeechPartPronoun, SpeechPartPreposition, SpeechPartConjunction, SpeechPartAdverb, SpeechPartParticle, SpeechPartInterjection, SpeechPartNoun}
)

func (t *Template) detectSpeechPart() {
	for _, part := range speechParts {
		if strings.HasPrefix(t.Title, part) {
			t.SpeechPart = part
		}
	}
	for _, part := range ruSpeechParts {
		if strings.HasPrefix(t.Title, part+" "+Prefix) {
			t.SpeechPart = part
		}
	}
}

func (t *Template) isPureSpeechPart() bool {
	return t.SpeechPart == SpeechPartPreposition || t.SpeechPart == SpeechPartConjunction || t.SpeechPart == SpeechPartAdverb ||
		t.SpeechPart == SpeechPartParticle || t.SpeechPart == SpeechPartInterjection
}
