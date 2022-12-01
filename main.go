package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiStatus(w http.ResponseWriter, r *http.Request) {

}

func apiBase64(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		if k == "complex" {
			for _, t := range v {
				plain, err := base64.StdEncoding.DecodeString(t)
				if err != nil {
					fmt.Fprintf(w, "%v\n", err)
				} else {
					fmt.Fprintf(w, "%s\n", plain)
				}
			}
		} else if k == "plain" {
			for _, t := range v {
				fmt.Fprintf(w, "%s\n", base64.StdEncoding.EncodeToString([]byte(t)))
			}
		}
	}
}

func main() {
	port := os.Getenv("APIPORT")
	if port == "" {
		port = "80"
	}
	http.HandleFunc("/v1/base64", apiBase64)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
