package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/matineghbali/chat-server/chat"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	hub := chat.NewHub()
	go hub.Run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(hub, w, r)
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServ: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Panicln(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Found", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")

}
