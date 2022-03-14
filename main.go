package main

import (
	"afkl/fumofuzzer/assets"
	"afkl/fumofuzzer/models/iterable"
	"afkl/fumofuzzer/models/output"
	"afkl/fumofuzzer/models/payload"
	"afkl/fumofuzzer/models/request"
	"afkl/fumofuzzer/models/response"
	"afkl/fumofuzzer/models/response/matcher"
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
		sort    = app.StringOpt("s sort", "", "Set the sort by what.")
		matches = app.StringsOpt("m matches", []string{}, "Set the match by what.")
		out     = app.StringOpt("o output", "stdout", "Set the output to where.")
		formatp = app.StringOpt("f farmat", "json", "Set the output format.")
		// timeout    = app.IntOpt("timeout", 20, "Request timeout.")

		threads = app.IntOpt("threads", 10000, "Threads number.")

		iterator     = app.StringOpt("i iterator", "chain", "Set the iterator.")
		payloadsList = app.StringsOpt("p payloads", []string{}, "Set the payloads.")

		fumo = app.StringOpt("fumo", "fumo", "fumo")
	)
	app.Spec = "-v | --fumo=<fumo> | -t=<target> [--threads=<threads>] [-X=<method>] [-i=<iterator>] [-s=<sort>] [-m=<matches>] [-o=<output>] [-p=<payloadsList>]..."
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

			sortp, ok := sorter.SORTER_MAP[*sort]
			if !ok {
				sortp = nil
			}

			var matchesp []matcher.Matcher
			for _, match := range *matches {
				matchesp = append(matchesp, matcher.SelectMatcher(match))
			}

			resps := response.NewFuzzResponses(sortp, matchesp)

			coll := response.NewFuzzResponseCollector(resps)
			coll.ExecCollector()

			pool.Submit(temp, coll)
			pool.Wait()
			pool.Close()

			output.NewOutput(*out, *formatp).Start(*temp, *resps)
		} else {
			log.Println(assets.NobodySeeingKoishi)
		}
	}
	app.Run(os.Args)
}
