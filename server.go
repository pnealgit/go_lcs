package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
)


var addr = flag.String("addr", "localhost:8081", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func talk(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		} //end of if on message err

		junk := string(message)
//fmt.Println("JUNK: ",junk)

		if strings.Contains(junk, "state") {
			message = do_lcs( message)
                        fmt.Println("back - message is ",message)
			err = c.WriteMessage(mt, message)
                       
			if err != nil {
				log.Println("write:", err)
				break
			} //end of if on write
		}

		if strings.Contains(junk, "num_episodes") {
			fmt.Println("NUM EPISODES!!")
		}

		outmap := make(map[string]string)
		outmap["status"] = "ok"
		message, err = json.Marshal(outmap)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		} //end of if on write
	} //end of for loop
} //end of talk

func main() {
	flag.Parse()
	log.SetFlags(0)
        get_parameters()
        p.Classifiers = nil

	http.HandleFunc("/talk", talk)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("\nlistening on 8081")
	log.Fatal(http.ListenAndServe(*addr, nil))
} //end of main
