package matcher

import (
	"log"
	"regexp"

	"github.com/go-resty/resty/v2"
)

type RegexpMatcher struct {
	Reg *regexp.Regexp
}

func (match *RegexpMatcher) SetRegexp(regexpString string) {
	reg, err := regexp.Compile(regexpString)
	if err != nil {
		log.Fatalln("Build Matcher Error")
	}

	match.Reg = reg
}

func (match *RegexpMatcher) Find(resps []resty.Response) [][]string {
	datas := make([][]string, 0)
	for _, resp := range resps {
		datas = append(datas, match.Reg.FindAllString(resp.String(), -1))
	}
	return datas
}

func (match *RegexpMatcher) Match(resps []resty.Response) []bool {
	datas := make([]bool, 0)
	for _, resp := range resps {
		datas = append(datas, match.Reg.MatchString(resp.String()))
	}
	return datas
}
