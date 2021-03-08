package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tcnksm/go-httpstat"
)

func main() {
	DestinationV1()
	DestinationV2()
}

func DestinationV1() {
	req, err := http.NewRequest(http.MethodGet, "https://api-tracking-staging.sicepat.com/v1/customer/destination", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("api-key", "23a987a9acc11ad3bfb615f3b35c3a59")

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	cli := http.DefaultClient
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	// end := time.Now()

	// Show the results
	log.Printf("Server processing v1: %d ms", int(result.ServerProcessing/time.Millisecond))
	log.Printf("Content transfer v1: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
}

func DestinationV2() {
	req, err := http.NewRequest(http.MethodGet, "https://api-tracking-staging.sicepat.com/v2/customer/destination", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("api-key", "23a987a9acc11ad3bfb615f3b35c3a59")

	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	cli := http.DefaultClient
	res, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	// end := time.Now()

	// Show the results
	log.Printf("Server processing v2: %d ms", int(result.ServerProcessing/time.Millisecond))
	log.Printf("Content transfer v2: %d ms", int(result.ContentTransfer(time.Now())/time.Millisecond))
}
