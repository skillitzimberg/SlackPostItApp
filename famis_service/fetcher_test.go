package main

import "testing"

func TestFetcher_Execute(t *testing.T) {
	fetcher := Fetcher{}
	tok, _, err := fetcher.Login("apptree", "moreAPPS", "https://st-ccsd.accruenttest.net")
	if err != nil {
		t.Fail()
		panic(err)
	}
	println(tok)
}
