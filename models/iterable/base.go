package iterable

import (
	"afkl/fumofuzzer/models/payload"
)

type BaseIterator interface {
	IsEnd() bool
	Scan() bool
	Value() []string
	Exec(Payloads []payload.Payload)
	Channel() chan []string
}

var ITER_MAP = map[string]BaseIterator{
	"chain":   NewChainIterator(),
	"zip":     NewZipIterator(),
	"product": NewProductIterator(),
}
