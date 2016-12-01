package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	server    *httptest.Server
	body      io.Reader
	widgetUrl string
)

func init() {
	router := NewRouter()
	server = httptest.NewServer(router) //Creating new server with the widget handlers
	widgetUrl = server.URL
}

func TestSomething(t *testing.T) {
	// test stuff here...
}

func TestGreeting(t *testing.T) {
	request, err := http.NewRequest("GET", widgetUrl, body)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if strings.TrimSpace(string(data)) != "Greeting!" {
		t.Errorf("Greeting! expected |%s|", string(data))
	}
}
