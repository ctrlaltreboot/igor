package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", echo)         // each request calls handler
	http.HandleFunc("/ean", ean)       // each request calls handler
	http.HandleFunc("/hotels", hotels) // each request calls handler
	log.Fatal(http.ListenAndServe("127.0.0.1:8088", nil))
}

// echo echos the Path component of the request URL r.
func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func ean(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func hotels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Ger(url)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v:", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
