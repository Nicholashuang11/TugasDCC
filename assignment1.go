package main

import (
    "fmt"
    "strings"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", hellowelcomehandler)
	http.HandleFunc("/upper",upperhandler )
    http.ListenAndServe(":8080", nil)
}
//endpoint 1: endpoint 1 akan menampilkan Hello, Welcome!. Good to see you
func hellowelcomehandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Welcome!. Good to see you")
}
//endpoint 2: endpoint 2 bisa mengubah text menjadi huruf kapital
func upperhandler(w http.ResponseWriter, r *http.Request){
    text := r.URL.Query().Get("text")

	
	if text == "" {
		fmt.Fprintf(w, "No text available")
		return
	}
	uppertext := strings.ToUpper(text)
	fmt.Fprintf(w, "%s", uppertext)
}
