package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	loopSqrt()
	tpl, _ := template.ParseFiles("index.html")
	message := greeting("Code.education Rocks !!! + square loop")
	data := map[string]string{
		"Message": message,
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}

func sqrt(n float64) float64 {
	return n * n
}

func loopSqrt() {
	var x float64
	x = 0.0001
	for i := 0; i <= 1000000; i++ {
		x += sqrt(x)
	}
	fmt.Print("finish !")
}
