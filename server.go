package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
        "math"
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

		if strings.Contains(junk, "State") {
                        //send back an action immediately
			message = get_action(message)
			err = c.WriteMessage(mt, message)
                     
                        //update A and A1 
                        update_lcs()
 
			if err != nil {
				log.Println("write:", err)
				break
			} //end of if on write
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

        //fix this so it is coordinated
        canvas_width = 480
        canvas_height = 300
        canvas_max_distance = math.Hypot(canvas_width,canvas_height)
        get_parameters()
        population = make(map[string]Classifier)
        reward_minus_one = 0.0
        my_time = 0.0
        sum_reward = 0.0

	http.HandleFunc("/talk", talk)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("\nlistening on 8081")
	log.Fatal(http.ListenAndServe(*addr, nil))
} //end of main
