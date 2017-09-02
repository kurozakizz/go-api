package main

import(
	"net/http"
	"fmt"
)

func response(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	rw.Write([]byte("Hello World"))
}

func fail(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	rw.Write([]byte("Fail na"))
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/fail", fail)
	http.ListenAndServe(":8080", nil)
}