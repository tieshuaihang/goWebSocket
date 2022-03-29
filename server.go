package main

import "net/http"

func wsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello websocket"))
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
