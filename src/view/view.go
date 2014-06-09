package view

import(
	"fmt"
	"net/http"
	"log"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>URL: %s<br>", r.URL.Path[1:])
	fmt.Fprintf(w, "<a href=\"/exit\">exit</a></html>")
}

func exitHandler(w http.ResponseWriter, r *http.Request) {
	log.Fatal("Server shutdown")
}