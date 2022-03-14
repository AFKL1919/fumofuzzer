package format

import (
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
)

type Formatter interface {
	Exec(temp request.FuzzRequestTemplate, result response.FuzzResponses) string
}
