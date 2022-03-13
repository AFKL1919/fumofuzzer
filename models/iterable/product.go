package iterable

import (
	"afkl/fumofuzzer/models/payload"
	"log"
)

type ProductIterator struct {
	isEnd bool
	token []string

	channel chan []string
}

func NewProductIterator() *ProductIterator {
	return &ProductIterator{
		isEnd:   false,
		channel: make(chan []string),
	}
}

func (iter *ProductIterator) product(sets ...[]string) {
	lens := func(i int) int { return len(sets[i]) }
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		iter.channel <- r
	}
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func (iter *ProductIterator) Exec(Payloads []payload.Payload) {
	if len(Payloads) == 0 {
		log.Fatalf("Payloads Num is 0")
	}
	go func() {
		values := make([][]string, 0)
		defer close(iter.channel)
		for _, payload := range Payloads {
			values = append(values, payload.Value)
		}
		iter.product(values...)
	}()
}

func (iter *ProductIterator) IsEnd() bool {
	return iter.isEnd
}

func (iter *ProductIterator) Scan() bool {
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

func (iter *ProductIterator) Value() []string {
	return iter.token
}

func (iter *ProductIterator) Channel() chan []string {
	return iter.channel
}
