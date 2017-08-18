package ean

import (
	"fmt"
	"github.com/ctrlaltreboot/igor/helper"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	b, err := helper.Fetch("http://127.0.0.1:5092/ean")
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", b)
}
