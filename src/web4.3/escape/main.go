package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/index", escape)
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func escape(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		t, _ := template.ParseFiles("index.gtpl")

		t.Execute(w, nil)
	} else {

		r.ParseForm()

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

	}
}
