package main

import (
	"fmt"
	"github.com/astaxie/goredis"
	"net/http"
)

const (
	GO_WEB_NAME string = "GO_WEB_NAME"
	GO_LIST     string = "go_list"
)

func main() {
	http.HandleFunc("/redis", doRedis)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}

}

func doRedis(w http.ResponseWriter, r *http.Request) {
	var client goredis.Client

	//字符串操作
	client.Set(GO_WEB_NAME, []byte("test go"))
	value, _ := client.Get(GO_WEB_NAME)
	fmt.Println("value:", string(value))
	client.Del(GO_WEB_NAME)

	//list操作
	vals := []string{"a", "b", "c"}

	for _, temp := range vals {
		client.Rpush(GO_LIST, []byte(temp))
	}

	results, _ := client.Lrange(GO_LIST, 0, 3)

	for key, temp := range results {
		fmt.Println("key:", string(key), " _ value:", string(temp))
	}

	client.Del(GO_LIST)
}
