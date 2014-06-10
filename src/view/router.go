package view

import (
	"net/http"
)

func Run() {
	http.HandleFunc("/exit", exitHandler)
	http.HandleFunc("/", defaultHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Konnte Server nicht starten")
	}
}