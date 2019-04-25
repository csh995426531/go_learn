package main

import (
	"fmt"
	"net/http"
)

type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string
}

func main() {
	http.HandleFunc("/cookie", doCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}

func doCookie(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{Name: "username", Value: "value"}
	http.SetCookie(w, &cookie)

	r_cookie, _ := r.Cookie("username")

	fmt.Println(w, r_cookie)

}
