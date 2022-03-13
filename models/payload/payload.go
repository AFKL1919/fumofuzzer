package payload

import (
	"afkl/fumofuzzer/models/payload/filter"
	"afkl/fumofuzzer/models/payload/loader"
	"log"
	"strings"
)

var PAYLOAD_LOADER_MAP = map[string]loader.PayloadLoader{
	"list":  new(loader.ListPayloadLoader),
	"file":  new(loader.FilePayloadLoader),
	"range": new(loader.RangePayloadLoader),
	"stdin": new(loader.StdinPayloadLoader),
}

var PAYLOAD_FILTER_MAP = map[string]filter.PayloadFilter{
	"none": new(filter.NonePayloadFilter),
	"md5":  new(filter.Md5PayloadFilter),
}

type Payload struct {
	Value []string

	Loader loader.PayloadLoader
	Filter filter.PayloadFilter
}

func (p *Payload) Load(value string) {
	for _, data := range p.Loader.Load(value) {
		p.Value = append(p.Value, p.Filter.Encode(data))
	}
}

func selectFilter(filterType string) filter.PayloadFilter {
	if filter, ok := PAYLOAD_FILTER_MAP[filterType]; ok {
		return filter
	} else {
		return PAYLOAD_FILTER_MAP["none"]
	}
}

func NewPayload(sourceData string) Payload {

	var (
		loaderValue    string
		loaderType     string
		payloadp       Payload
		payloadpFitler filter.PayloadFilter
	)

	split := strings.Split(sourceData, ",")
	splitLen := len(split)

	switch splitLen {
	case 1:
		payloadpFitler = selectFilter("")
		loaderType = strings.ToLower(split[0])

	case 2:
		loaderValue = split[1]
		payloadpFitler = selectFilter("")
		loaderType = strings.ToLower(split[0])

	case 3:
		loaderValue = split[1]
		loaderType = strings.ToLower(split[0])
		payloadpFitler = selectFilter(split[2])

	default:
		log.Fatalf("Syntax Error, Too Many Commas: %s\n", sourceData)
	}

	if loader, ok := PAYLOAD_LOADER_MAP[loaderType]; ok {
		payloadp = Payload{
			Loader: loader,
			Filter: payloadpFitler,
		}
	} else {
		log.Fatalf("Unknown Error in Paser: %s\n", sourceData)
	}

	payloadp.Load(loaderValue)
	return payloadp
}
