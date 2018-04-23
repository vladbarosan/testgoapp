package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"net/http"
	"os"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	http.HandleFunc("/", f)
	fmt.Println("Starting default static web app, port is %s...", port)
	http.ListenAndServe(port, nil)
}

func f(w http.ResponseWriter, r *http.Request) {

	basicColors := mapset.NewSet()
	basicColors.Add("Red")
	basicColors.Add("Blue")
	basicColors.Add("Green")
	if basicColors.Contains("Green") {
		fmt.Fprintf(w, "Yay! 'Green' is a supported color")
	} else {
		fmt.Fprintf(w," What a disappointment! 'Green' is not a basic color")
	}
}
