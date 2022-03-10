package loader

import (
	"bufio"
	"os"
)

type StdinPayloadLoader struct{}

func (loader *StdinPayloadLoader) Load(value string) []string {
	payloads := make([]string, 0)
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		payloads = append(payloads, scan.Text())
	}

	return payloads
}
