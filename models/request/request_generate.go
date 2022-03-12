package request

import "afkl/fumofuzzer/models/iterable"

var FUZZ_KEY_WORD = "FUZZ"

var UA_LIST = []string{
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.14) Gecko/2009091010",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.10) Gecko/2009042523",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.13; ) Gecko/20101203",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7b) Gecko/20040421",
}

type FuzzRequest struct {
	TargetUrl    string
	TargetHeader map[string]string
	TargetBody   string

	Iterator iterable.BaseIterator
}

func NewFuzzRequest(url string, header map[string]string, body string) *FuzzRequest {
	return &FuzzRequest{
		TargetUrl:    url,
		TargetHeader: header,
		TargetBody:   body,
	}
}
