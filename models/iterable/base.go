package iterable

import (
	"afkl/fumofuzzer/models/payload"
)

type BaseIterator interface {
	IsEnd() bool
	Scan() bool
	Value() []string
	Exec(Payloads []payload.Payload)
}
