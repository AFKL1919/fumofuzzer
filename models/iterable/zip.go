package iterable

import (
	"afkl/fumofuzzer/models/payload"
	"log"
	"math"
)

type ZipIterator struct {
	isEnd bool
	token []string

	channel chan []string
}

func NewZipIterator() *ZipIterator {
	return &ZipIterator{
		isEnd:   false,
		channel: make(chan []string),
	}
}

func findMinLenghtPayloadValue(Payloads []payload.Payload) int {
	tmp := math.MaxInt
	for _, payload := range Payloads {
		if len(payload.Value) < tmp {
			tmp = len(payload.Value)
		}
	}
	return tmp
}

func (iter *ZipIterator) Exec(Payloads []payload.Payload) {
	payloadValueLenght := findMinLenghtPayloadValue(Payloads)
	if payloadValueLenght == 0 {
		log.Fatalf("All Payloads Has zero-length Payload")
	}

	go func() {
		defer close(iter.channel)
		for i := 0; i < payloadValueLenght; i++ {
			var data []string
			for _, payload := range Payloads {
				data = append(data, payload.Value[i])
			}
			iter.channel <- data
		}
	}()
}

func (iter *ZipIterator) IsEnd() bool {
	return iter.isEnd
}

func (iter *ZipIterator) Scan() bool {
	if iter.IsEnd() {
		return false
	}

	if data, ok := <-iter.channel; ok {
		iter.token = data
		return true
	} else {
		iter.isEnd = true
		return false
	}
}

func (iter *ZipIterator) Value() []string {
	return iter.token
}
