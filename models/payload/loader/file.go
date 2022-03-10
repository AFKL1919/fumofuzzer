package loader

import (
	"afkl/fumofuzzer/utils"
	"log"
)

type FilePayloadLoader struct{}

func (loader *FilePayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)

	fp, err := utils.Open(value)
	if err != nil {
		log.Fatalf("Load File [%s] Failed: %s\n", value, err)
	}
	defer fp.Close()

	for {
		payload, ok := fp.ReadLine()
		if !ok {
			break
		}
		payloads = append(payloads, payload)
	}

	return payloads
}
