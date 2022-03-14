package response

import (
	"afkl/fumofuzzer/models/response/matcher"
	"afkl/fumofuzzer/models/response/sorter"

	"github.com/go-resty/resty/v2"
)

type FuzzResponses struct {
	FuzzedResponses []resty.Response

	Sorter  sorter.BaseSorter
	Matcher matcher.Matcher
}

func NewFuzzResponses(sort sorter.BaseSorter /*, match matcher.Matcher*/) *FuzzResponses {
	return &FuzzResponses{
		Sorter: sort,
	}
}

func (fuzz *FuzzResponses) Sort() {
	fuzz.Sorter.Sort(fuzz.FuzzedResponses)
}
