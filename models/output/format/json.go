package format

import (
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
	"encoding/json"
)

type JsonFormatter struct{}

func (formatter JsonFormatter) Exec(temp request.FuzzRequestTemplate, result response.FuzzResponses) string {
	var maps []map[string]interface{}
	matches, isSetMatcher := result.Match()
	result.Sort()

	for respID, resp := range result.FuzzedResponses {
		map0 := map[string]interface{}{
			"url":    resp.Request.URL,
			"status": resp.StatusCode(),
			"body":   resp.String(),
			"header": resp.Header(),
			"size":   resp.Size(),
		}

		if isSetMatcher {
			var match []bool
			for _, matchOne := range matches {
				match = append(match, matchOne[respID])
			}
			map0["matches"] = match
		}

		maps = append(maps, map0)
	}

	var payloads []string
	for _, payload := range temp.Payloads {
		payloads = append(payloads, payload.Original)
	}

	data := map[string]interface{}{
		"url":      temp.TargetUrl,
		"payloads": payloads,
		"result":   maps,
	}

	rb, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(rb)
}
