package response

import (
	"github.com/go-resty/resty/v2"
)

type FuzzResponseCollector struct {
	FuzzResps   *FuzzResponses
	RespChannel chan *resty.Response
}

func NewFuzzResponseCollector(resps *FuzzResponses) FuzzResponseCollector {
	return FuzzResponseCollector{
		FuzzResps:   resps,
		RespChannel: make(chan *resty.Response),
	}
}

func (coll FuzzResponseCollector) Channel() chan *resty.Response {
	return coll.RespChannel
}

func (coll FuzzResponseCollector) ExecCollector() {
	go func() {
		for resp := range coll.Channel() {
			coll.FuzzResps.FuzzedResponses = append(coll.FuzzResps.FuzzedResponses, *resp)
		}
	}()
}
