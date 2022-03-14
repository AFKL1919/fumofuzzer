package matcher

import (
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Matcher interface {
	Find([]resty.Response) [][]string
	Match([]resty.Response) []bool
}

func SelectMatcher(s string) Matcher {
	data := strings.Split(s, ",")
	if len(data) != 2 {
		log.Fatalf("Unknow Matcher: %s\n", s)
	}

	matchType := data[0]
	matchValue := data[1]

	switch strings.ToLower(matchType) {
	case "regexp":
		match := new(RegexpMatcher)
		match.SetRegexp(matchValue)
		return match
	default:
		log.Fatalf("Unknow Matcher: %s\n", s)
	}

	return nil
}
