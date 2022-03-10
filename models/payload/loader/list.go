package loader

import "strings"

type ListPayloadLoader struct{}

func (loader *ListPayloadLoader) Load(value string) []string {
	return strings.Split(value, "-")
}
