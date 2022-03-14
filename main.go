package main

import (
	"afkl/fumofuzzer/assets"
	"afkl/fumofuzzer/models/iterable"
	"afkl/fumofuzzer/models/payload"
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
	"afkl/fumofuzzer/models/response/sorter"
	"log"
	"os"

	cli "github.com/jawher/mow.cli"
)

var VERSION = "FUMO-0.0.1"

func main() {
	app := cli.App("FumoFuzzer", "OMG! The FUMO! ᗜˬᗜ")
	app.Version("v version", VERSION)
	var (
		target  = app.StringOpt("t target", "", "Set the target url.")
		method  = app.StringOpt("X method", "GET", "Set the method.")
		headers = app.StringsOpt("H Headers", []string{}, "Set the headers.")
		body    = app.StringOpt("d body", "", "Set the send body.")
		// timeout    = app.IntOpt("timeout", 20, "Request timeout.")

		threads = app.IntOpt("threads", 10000, "Threads number.")

		iterator     = app.StringOpt("i iterator", "chain", "Set the iterator.")
		payloadsList = app.StringsOpt("p payloads", []string{}, "Set the payloads.")

		fumo = app.StringOpt("fumo", "fumo", "fumo")
	)
	app.Spec = "-v | --fumo=<fumo> | -t=<target> [--threads=<threads>] [-X=<method>] [-i=<iterator>] [-p=<payloadsList>]..."
	app.Action = func() {
		if *fumo == "fumo" {
			pool := request.InitRequestPool(*threads)

			var payloads []payload.Payload
			for _, data := range *payloadsList {
				payloads = append(payloads, payload.NewPayload(data))
			}

			iter, ok := iterable.ITER_MAP[*iterator]
			if !ok {
				iter = iterable.ITER_MAP["chain"]
			}

			temp := request.NewFuzzRequestTemplate(
				*method,
				*target,
				*headers,
				*body,
				payloads,
				iter,
			)

			sort := new(sorter.SizeSorter)
			resps := &response.FuzzResponses{
				Sorter: sort,
			}

			coll := response.NewFuzzResponseCollector(resps)
			coll.ExecCollector()

			pool.Submit(temp, coll)
			pool.Wait()
			pool.Close()

		} else {
			log.Println(assets.NobodySeeingKoishi)
		}
	}
	app.Run(os.Args)
}
