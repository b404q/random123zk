package main

import (
	"fmt"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", index)

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`aa`))
}
