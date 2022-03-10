package loader

import (
	"fmt"
	"log"
	"strings"
)

type RangePayloadLoader struct{}

func (loader *RangePayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)
	index := strings.Index(value, "-")

	if len(value[:index]) != 1 || len(value[index+1:]) != 1 {
		log.Fatalf("Syntax Error in Range List Loader: %s", value)
	}

	var (
		start = int(value[:index][0])
		end   = int(value[index+1:][0])
	)

	for i := start; i <= end; i++ {
		payloads = append(payloads, fmt.Sprintf("%c", i))
	}

	return payloads
}
