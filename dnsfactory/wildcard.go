package factory

import (
	"regexp"
	"strings"

	loglib "log"

	gintersect "github.com/yashtewari/glob-intersection"
)

type phrase struct {
	wildcard string
	regexp   *regexp.Regexp
}

func newPhrase(a1 string) *phrase {
	a := &phrase{
		wildcard: a1,
	}
	a.regexp = regexp.MustCompile(a.convertPhraseToRegexp())
	return a
}
func newPhraseSlice(l []string) []*phrase {
	al := []*phrase{}
	for _, a := range l {
		al = append(al, newPhrase(a))
	}
	return al
}

func mustWildcardIntersect(a1, a2 *phrase) bool {
	intersect, err := wildcardIntersect(a1, a2)
	if err != nil {
		loglib.Fatal("fatal wildcard intersection", err)
	}
	return intersect
}

func wildcardIntersect(a1, a2 *phrase) (bool, error) {
	intersect, err := gintersect.NonEmpty(a2.regexp.String(), a1.regexp.String())
	if err != nil {
		loglib.Print("error determining wildcard intersection",
			"Phrase1.wildcard", a1.wildcard,
			"Phrase1.regexp", a1.regexp,
			"Phrase2.wildcard", a2.wildcard,
			"Phrase2.regexp", a2.regexp)
		return true, err
	}
	return intersect, nil
}

func (a *phrase) convertPhraseToRegexp() string {
	// replace "." with "\." to match literal "." character
	regexpStr := strings.Replace(a.wildcard, ".", `\.`, -1)

	// replace "?" with "[A-Za-z0-9\-]" to match any one character excluding '.'
	regexpStr = strings.Replace(regexpStr, "?", `[A-Za-z0-9\-]`, -1)

	// replace "*" with "[A-Za-z0-9\-]+" to greedy match any character(s) excluding '.'
	regexpStr = strings.Replace(regexpStr, "*", `[A-Za-z0-9\-]+`, -1)

	return regexpStr
}
