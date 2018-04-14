package main

import (
    "net/http"
)

func main() {
	r := NewRouter()
	http.ListenAndServe(":80", r)
}
