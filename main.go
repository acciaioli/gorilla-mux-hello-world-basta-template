package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{{ .endpoint }}", Handler)
	http.Handle("/", r)
	fmt.Println("Starting up on {{ .port }}")
	log.Fatal(http.ListenAndServe(":{{ .port }}", nil))
}

func Handler(w http.ResponseWriter, req *http.Request) {	
	{{- if .json}}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"text": "{{ .responseText}}"}`)
	{{- else}}
	fmt.Fprintln(w, "{{ .responseText}}")
	{{- end}}
}
