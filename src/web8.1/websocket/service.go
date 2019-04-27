package main

import (
	"fmt"
	"github.com/golang/net/websocket"
	"log"
	"net/http"
)

func main(){
	http.Handle("/",websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Echo(ws *websocket.Conn)  {

	var err error

	for {
		var reply  string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Receive back form client:"+reply)

		msg := "Receive:"+reply
		fmt.Println("Sending to client:"+msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}
