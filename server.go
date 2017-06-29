package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
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

func getRandomFloat32(min float32, max float32) float32 {
	return 0.0 + (rand.Float32() * (max - min)) + min
}

func getRandomInt(min int, max int) int {
	return rand.Intn(max-min) + min
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

		log.Printf("recv: %s", message)
		junk := string(message)
		if strings.Contains(junk, "sensor") {
			message = do_generates(message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			} //end of if on write
                } 
		if strings.Contains(junk, "reward") {
			do_updates(message)
                } 

		outmap := make(map[string]string)
		outmap["status"] = "OK"
		message, err = json.Marshal(outmap)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		} //end of if on write
	} //end of for loop
} //end of talk

var actions = [8]string{"EE","SE","SS","SW","WW","NW","NN","NE"}
var action_angles =  map[string]float64{}

type Classifier struct {
    Id         int
    Condition  string
    Action     string
    Prediction float64
    Epsilon    float64
    Fitness    float64
    Experience float64
    Numerosity float64
    As         float64
    Ga_clicks  float64
}
type Classifier_set struct {
    Cl   []Classifier
}

var cset = []Classifier{}

var start_time = int32


    const NUMBER_CLASSIFIERS int = 1000;
    const BIGP float64 = 0.0;
    const BETA float64 = .15 //range 0.1-0.2 -- learning rate
    const ALPHA float64 = 0.1 //normally
    const EPSILON_0 float64 = 10.0 // ten percent of rho...
    const RHO int = 1000 //max reward ??
    const GAMMA float64 = .71 //from many problems discount factor
    const THETA_CLICKS int = 30 //range 25-50
    const CHI float64 = .5 //range .5 - 1.0 crossover probability
    const MU  float64 = .03 //mutation probability in selection
    const THETA_MNA int =  8 //to make sure every action is covered
    const THETA_DEL int = 20 //deletion threshold
    const SIGMA float64 = 0.1
    const THETA_SUB int = 20 //subsumption threshold
    const P_SHARP float64  = 0.33 //probability of flipping a condition value to #
    const P_INITIAL float64 = .00001;
    const EPSILON_INITIAL float64 = .00000 //prediction error
    const FITNESS_INITIAL float64 = .5
    const P_EXP float64 = .5 //probability exploration/exploitation

func main() {
        make_action_angles() 
        start_time = int32(time.Now().Unix())
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/talk", talk)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("listening on 8081")
	log.Fatal(http.ListenAndServe(*addr, nil))

} //end of main

func make_action_angles() {
    angle := 0.0
    delta := 2.0*math.Pi/float64(len(actions))
    for i:=0;i<len(actions);i++ {
        action_angles[actions[i]] = angle
        angle+= delta
    }
}
