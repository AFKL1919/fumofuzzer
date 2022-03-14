package response

import (
	"afkl/fumofuzzer/models/response/matcher"
	"afkl/fumofuzzer/models/response/sorter"

	"github.com/go-resty/resty/v2"
)

type FuzzResponses struct {
	FuzzedResponses []resty.Response

	Sorter   sorter.Sorter
	Matchers []matcher.Matcher
}

func NewFuzzResponses(sortp sorter.Sorter, matchesp []matcher.Matcher) *FuzzResponses {
	return &FuzzResponses{
		Sorter:   sortp,
		Matchers: matchesp,
	}
}

func (fuzz *FuzzResponses) IsSetSorter() bool {
	return fuzz.Sorter != nil
}

func (fuzz *FuzzResponses) IsSetMatcher() bool {
	return len(fuzz.Matchers) != 0
}

func (fuzz *FuzzResponses) Sort() bool {
	if fuzz.IsSetSorter() {
		fuzz.Sorter.Sort(fuzz.FuzzedResponses)
		return true
	} else {
		return false
	}
}

func (fuzz *FuzzResponses) Match() (matches [][]bool, isSetMatcher bool) {
	if fuzz.IsSetMatcher() {
		for _, matcher := range fuzz.Matchers {
			matches = append(matches, matcher.Match(fuzz.FuzzedResponses))
		}
		isSetMatcher = true
	} else {
		matches = [][]bool{}
		isSetMatcher = false
	}
	return
}
