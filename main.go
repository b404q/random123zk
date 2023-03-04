package main

import "net/http"

func main() {
	servedir := http.FileServer(http.Dir("sitko"))
	err := http.ListenAndServe(":80", servedir)
	if err != nil {
		panic(err)
	}
}
