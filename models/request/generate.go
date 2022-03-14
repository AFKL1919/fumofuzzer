package request

import (
	"afkl/fumofuzzer/models/iterable"
	"afkl/fumofuzzer/models/payload"
	"afkl/fumofuzzer/models/response"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"

	"github.com/go-resty/resty/v2"
)

var FUZZ_KEY_WORD = "FUZ%dZ"

var UA_LIST = []string{
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.14) Gecko/2009091010",
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.10) Gecko/2009042523",
	"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.13; ) Gecko/20101203",
	"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.7b) Gecko/20040421",
}

type FuzzRequestTemplate struct {
	Method string

	HasFuzzUrl bool
	TargetUrl  string

	HasFuzzHeaders bool
	TargetHeaders  map[string]string

	HasFuzzBody bool
	TargetBody  string

	Payloads []payload.Payload
	Iterator iterable.BaseIterator

	WG sync.WaitGroup
}

func hasFuzzString(value string) bool {
	return strings.Contains(value, "FUZ")
}

func NewFuzzRequestTemplate(
	method string, url string, headersString []string, body string,
	payloads []payload.Payload, iter iterable.BaseIterator,
) *FuzzRequestTemplate {

	var (
		headersMap     = make(map[string]string)
		hasFuzzUrl     = hasFuzzString(url)
		hasFuzzBody    = hasFuzzString(body)
		hasFuzzHeaders = false
	)

	for _, header := range headersString {
		hasFuzzHeaders = hasFuzzString(header)
		h := strings.Split(header, ": ")
		if len(h) == 2 {
			headersMap[h[0]] = h[1]
		} else {
			log.Printf("Format error HEADER(Has been ignored): %s\n", header)
			continue
		}
	}

	return &FuzzRequestTemplate{
		Method: method,

		HasFuzzUrl: hasFuzzUrl,
		TargetUrl:  url,

		HasFuzzHeaders: hasFuzzHeaders,
		TargetHeaders:  headersMap,

		HasFuzzBody: hasFuzzBody,
		TargetBody:  body,

		Payloads: payloads,
		Iterator: iter,
	}
}

type FuzzRequest struct {
	Data      []string
	Request   *resty.Request
	collector response.FuzzResponseCollector
}

func (fuzz FuzzRequestTemplate) GenerateFuzzRequest(FuzzData []string, coll response.FuzzResponseCollector) *FuzzRequest {
	url := fuzz.TargetUrl
	body := fuzz.TargetBody
	headers := fuzz.TargetHeaders
	FuzzDataNum := len(FuzzData)

	for i := 0; i < FuzzDataNum; i++ {
		key := fmt.Sprintf(FUZZ_KEY_WORD, i)
		replace := func(v string) string {
			return strings.ReplaceAll(
				v,
				key,
				FuzzData[i],
			)
		}

		if fuzz.HasFuzzUrl && strings.Contains(url, key) {
			url = replace(url)
			continue
		}

		if fuzz.HasFuzzBody && strings.Contains(body, key) {
			body = replace(body)
			continue
		}

		if fuzz.HasFuzzHeaders {
			for k, v := range headers {
				fuzzKey := strings.Contains(k, key)
				fuzzValue := strings.Contains(v, key)
				if fuzzKey {
					v = replace(v)
				} else if fuzzValue {
					k = replace(k)
				}

				if fuzzKey || fuzzValue {
					headers[k] = v
					break
				}
			}
			continue
		}
	}

	request := resty.New().R()
	request.URL = url
	request.Method = fuzz.Method
	request.SetHeaders(headers)
	request.SetBody(body)
	request.SetHeader("User-Agent", UA_LIST[rand.Intn(4)])

	return &FuzzRequest{
		Data:      FuzzData,
		Request:   request,
		collector: coll,
	}
}
