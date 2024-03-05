package main

import (
	"GoGuiWithFyne/Section_06_GoldWatcher_Part_03/repository"
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
)

var testApp Config

func TestMain(m *testing.M) {
	testApp.App = test.NewApp()
	testApp.MainWindow = testApp.App.NewWindow("")
	testApp.HttpClient = client
	testApp.DB = repository.NewTestRepository()
	os.Exit(m.Run())
}

var jsonToReturn = `
{
	"ts": 1709204325596,
	"tsj": 1709204315633,
	"date": "Feb 29th 2024, 05:58:35 am NY",
	"items": [
		{
		"curr": "USD",
		"xauPrice": 2030.8375,
		"xagPrice": 22.3579,
		"chgXau": -3.8475,
		"chgXag": -0.1031,
		"pcXau": -0.1891,
		"pcXag": -0.459,
		"xauClose": 2034.685,
		"xagClose": 22.461
		}
	]
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func newTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = newTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
