package main

import "testing"

const testValue1 = `{"Username":"apptree","Password":"moreAPPS","Url":"https://st-ccsd.accruenttest.net","Endpoint":"/MobileWebServices/apis/360facility/v1/assets", "ChunkSize": 100}`
const testValue2 = `{"Username":"apptree","Password":"moreAPPS","Url":"https://st-ccsd.accruenttest.net","Endpoint":"/MobileWebServices/apis/360facility/v1/assets", "Id": 2}`

func TestFetcher_Execute(t *testing.T) {
	fetcher := Fetcher{}
	val, err := fetcher.ExecuteJson(testValue1)
	if err != nil {
		panic(err)
	}
	print(val)
}

func TestFetcher_ExecuteSingle(t *testing.T) {
	fetcher := FetchSingle{}
	val, err := fetcher.ExecuteJson(testValue2)
	if err != nil {
		panic(err)
	}
	print(val)
}
