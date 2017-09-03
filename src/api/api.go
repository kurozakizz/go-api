package main

import (
	"calculator"
	"net/http"
	"strconv"
)

func response(rw http.ResponseWriter, r *http.Request) {
	question := "1,2"
	sum := calculator.Add(question)
	sumText := strconv.Itoa(sum)
	rw.Write([]byte(question + " = " + sumText))
}

func fail(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Fail na"))
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/fail", fail)
	http.ListenAndServe(":9000", nil)
}
