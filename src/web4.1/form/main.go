package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		//slice := [] string {"football", "basketball", "tennis"}
		//a:=Slice_diff(r.Form["interest"],slice)
		//if a != nil {
		//	http.Error(w, "interest err", 506)
		//}

		slice_status := false
		slice := []string{"apple", "pear", "banane"}

		for _, v := range slice {

			if v == r.Form.Get("fruit") {
				slice_status = true
			}
		}

		if slice_status == false {
			http.Error(w, "fruit value err", 505)
		}

		if len(r.Form["username"][0]) == 0 {
			http.Error(w, "username can't nil", 501)
			panic("username can't nil")
		} else if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("username")); !m {
			http.Error(w, "username is zh", 504)
		}

		age, err := strconv.Atoi(r.Form.Get("age"))

		if err != nil {
			http.Error(w, "age type err", 502)
		} else if age > 100 {
			http.Error(w, "age can't dayu 100", 503)
		}

		fmt.Println(len(r.Form["username"][0]))
		fmt.Println(r.Form["username"])
		fmt.Println(len(r.Form["username"]) == 0)

		fmt.Println("username:", r.Form["username"])
		fmt.Println("password", r.Form["password"])
		fmt.Println("username", r.Form.Get("username"))
		fmt.Println("username", len(r.Form.Get("username")))
		fmt.Println("age", r.Form["age"])
		fmt.Println("fruit", r.Form.Get("fruit"))
	}
}
