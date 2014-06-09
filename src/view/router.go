package view

import (
	"net/http"
)

func Run() {
	http.HandleFunc("/exit", exitHandler)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":80", nil)
}
