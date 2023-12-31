package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8900", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello port 8900")
}
