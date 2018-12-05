package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	client := ts.Client()
	res, err := client.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	greeting, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	fmt.Println(res.StatusCode)
	fmt.Println(res.Header.Get("Content-Type"))
	fmt.Println(string(greeting))
}
