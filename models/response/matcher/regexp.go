package matcher

import (
	"log"
	"regexp"
)

type RegexpMatcher struct {
	Reg *regexp.Regexp
}

func NewRegexpMatcher(regexpString string) RegexpMatcher {
	reg, err := regexp.Compile(regexpString)
	if err != nil {
		log.Fatalln("Build Matcher Error")
	}

	return RegexpMatcher{
		Reg: reg,
	}
}
