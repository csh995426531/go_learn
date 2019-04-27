package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func main(){
	mux := routes.New()
	mux.Get("/user/:uid", getUser)
	mux.Post("/user/:uid", updateUser)
	mux.Del("/user/:uid", deleteUser)
	mux.Put("/user", addUser)
	http.Handle("/", mux)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		panic(err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s", uid)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are post user %s", uid)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are del user %s", uid)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	a := r.Form.Get("***")
	fmt.Fprintf(w, "you are put user %s", a)
}
