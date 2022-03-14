package output

import (
	"afkl/fumofuzzer/models/output/format"
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
	"log"
	"os"
	"strings"
)

type Output struct {
	Write  *os.File
	Format format.Formatter
}

var FORMAT_MAP = map[string]format.Formatter{
	"json": new(format.JsonFormatter),
}

func NewOutput(out string, format string) *Output {
	var w *os.File
	if out == "" || strings.Compare(out, "stdout") == 0 {
		w = os.Stdout
	} else {
		fp, err := os.OpenFile(out, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			log.Fatalf("File Open Error: %s\n", err.Error())
		}

		w = fp
	}

	formatp, ok := FORMAT_MAP[strings.ToLower(format)]
	if !ok {
		formatp = FORMAT_MAP["json"]
	}

	return &Output{
		Write:  w,
		Format: formatp,
	}
}

func (output *Output) Start(temp request.FuzzRequestTemplate, result response.FuzzResponses) {
	_, err := output.Write.WriteString(output.Format.Exec(temp, result))
	if err != nil {
		log.Fatalf("File Write Error: %s\n", err.Error())
	}
}
