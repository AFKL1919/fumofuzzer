package main

import (
	"afkl/fumofuzzer/models/iterable"
	"afkl/fumofuzzer/models/payload"
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
	"afkl/fumofuzzer/models/response/sorter"
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
	var p payload.Payload
	p = payload.NewPayload("file,./assets/payload.txt,md5")
	log.Println(p)

	p = payload.NewPayload("list,a-b-c-1-2-3")
	log.Println(p)

	p = payload.NewPayload("range,0-z")
	log.Println(p)
}

func TestChainIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Value: []string{"1", "2", "3"},
		},
		{
			Value: []string{"a", "b", "c"},
		},
	}
	c := iterable.NewChainIterator()
	log.Println("inited iter")
	c.Exec(payloads)
	log.Println("execed iter")

	log.Println("start scan")
	for c.Scan() {
		log.Println("scan get one")
		log.Println(c.Value())
		log.Println("scan put one")
	}
	log.Println("all end...")
}

func TestZipIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Value: []string{"1", "2", "3"},
		},
		{
			Value: []string{"a", "c"},
		},
		{
			Value: []string{"7", "8", "9"},
		},
	}
	c := iterable.NewZipIterator()
	log.Println("inited iter")
	c.Exec(payloads)
	log.Println("execed iter")

	log.Println("start scan")
	for c.Scan() {
		log.Println("scan get one")
		log.Println(c.Value())
		log.Println("scan put one")
	}
	log.Println("all end...")
}

func TestProductIteratorExec(t *testing.T) {
	payloads := []payload.Payload{
		{
			Value: []string{"1", "2", "3"},
		},
		{
			Value: []string{"a", "b", "c"},
		},
		{
			Value: []string{"7", "8", "9"},
		},
	}
	c := iterable.NewProductIterator()
	log.Println("inited iter")
	c.Exec(payloads)
	log.Println("execed iter")

	log.Println("start scan")
	for c.Scan() {
		log.Println("scan get one")
		log.Println(c.Value())
		log.Println("scan put one")
	}
	log.Println("all end...")
}

func TestSome(t *testing.T) {
	pool := request.InitRequestPool(10000)

	payloads := []payload.Payload{
		payload.NewPayload("list,1-22-333"),
		payload.NewPayload("list,aaaa-bbbbbb-ccccccc"),
	}
	iter := iterable.NewChainIterator()
	temp := request.NewFuzzRequestTemplate(
		"POST",
		"http://127.0.0.1:9000/FUZ0Z/",
		[]string{
			"Cookie: 233=233;",
		},
		"a=1",
		payloads,
		iter,
	)

	sort := new(sorter.SizeSorter)
	resps := response.NewFuzzResponses(sort)

	coll := response.NewFuzzResponseCollector(resps)
	coll.ExecCollector()

	pool.Submit(temp, coll)
	pool.Wait()
	pool.Close()

	log.Println("=====================1st======================")
	for _, resp := range resps.FuzzedResponses {
		log.Printf("%d, ", resp.Size())
	}

	log.Println("=====================2rd======================")
	resps.Sort()
	for _, resp := range resps.FuzzedResponses {
		log.Printf("%d, ", resp.Size())
	}
}
