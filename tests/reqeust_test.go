package main

import (
	"afkl/fumofuzzer/models/payload"
	"log"
	"testing"

	resty "github.com/go-resty/resty/v2"
)

func TestResty(t *testing.T) {
	client := resty.New()
	resp, _ := client.
		R().
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.7113.93 Safari/537.36").
		Get("http://127.0.0.1:9999")
	log.Println(resp)
}

func TestNewPayload(t *testing.T) {
	var p *payload.Payload
	p = payload.NewPayload("file,./assets/payload.txt,md5")
	// log.Println(p)

	p = payload.NewPayload("list,a-b-c-1-2-3")
	log.Println(p)

	p = payload.NewPayload("range,0-z")
	log.Println(p)
}

func TestSome(t *testing.T) {
	test := map[string]string{
		"aaa": "aaa",
	}

	log.Println(test["bbb"])
}
