package main

import (
	"encoding/json"
	"fmt"
        "bytes"
        "os"
        "math/rand"
)

type State_record struct {
	State  []int
        Reward float64
}

type Angle_record struct {
    Status  string
    Angle int
}

var state_record State_record
 
var angle_record Angle_record

type Classifier struct {
    Condition    string
    Action       int
    p            float64      //predicted payoff
    Epsilon float64
    F            float64      //fitness
    Exp   float64 
    ts    float64
    as float64                   //average selection set size
    n   float64                  //numerosity
}

var p []Classifier
var m []Classifier
var A []Classifier
var A1 []Classifier

var my_time float64
var action int

var state_string string
var pa  map[int]float64 //prediction array

var reward float64
var reward_minus_one float64

func get_action( message []byte) []byte {
      
  
	var err error
        var state_record State_record

	jerr := json.Unmarshal(message, &state_record)
	if jerr != nil {
		fmt.Println("error on state unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr
        state_string = ""
        state_string = convert_state_input(state_record)

        reward = state_record.Reward
        if reward > 1.0 {
           fmt.Printf("reward big : %f \n",reward) 
        }

        fmt.Printf("REWARD IS: %f\n",reward)

        if len(state_string) != 36 {
            fmt.Printf("length is : %4d\n",len(state_string))
            fmt.Printf("%v\n",state_record)
            fmt.Printf("%s\n",state_string)
            os.Exit(2)
        }
   
        my_time++ 
        fmt.Printf("calling match set with string %s\n",state_string)

        generate_match_set(state_string)
fmt.Printf("back from generate match set\n")

        generate_prediction_array()
        select_action()
        generate_action_set()

        angle_record.Angle = action
        angle_record.Status = "angle"
        fmt.Printf("returning angle_record %+v\n",angle_record)
	message,err  = json.Marshal(angle_record)
        if err != nil {
                fmt.Println("bad angles Marshal")
        }
	return message
} //end of do_update

func think(state_record State_record) float64 {
	return 1.0	
} //end of think

func convert_state_input(state_record State_record) string {
    var s bytes.Buffer
    
    si := ""
    i := 0
    si = fmt.Sprintf("%06b",state_record.State[0]%20)
    s.WriteString(si)

    si = fmt.Sprintf("%06b",state_record.State[1]%20)
    s.WriteString(si)

    for i=2;i<len(state_record.State);i++ {
        si = fmt.Sprintf("%04b",state_record.State[i])
        s.WriteString(si)
    }
    return s.String()
}

func dump_population() {
     i := 0
     kntr := make(map[string]int)

     for i=0;i<len(p);i++ {
       kntr[p[i].Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("%d %s\n",v,k)
     }

} //end of dump

func dump_match_set(state_string string) {
     fmt.Println("DUMP MATCH SET \n")
     fmt.Printf( "%s\n",state_string)
     i := 0
     kntr := make(map[string]int)
     kntr2 := make(map[int]int)

     for i=0;i<len(m);i++ {
       kntr[m[i].Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("Match set %d %s\n",v,k)
     }

     for i=0;i<len(m);i++ {
       kntr2[m[i].Action]++
     }
     fmt.Printf("Action Counts\n")

     for k,v := range kntr2 {
        fmt.Printf("ACTIONS %d %d\n",k,v)
     }

} //end of dump_match_set


func generate_prediction_array() {
    fmt.Printf("in prediction array")

    i := 0
    pa := make(map[int]float64) 
    fsa := make(map[int]float64) 

    for i=0;i<len(m);i++ {
        pa[m[i].Action] += m[i].p * m[i].F
        fsa[m[i].Action] += m[i].F
    } 
    for i:=0;i<len(possible_actions);i++ {
        if fsa[i] > 0.0 {
            pa[i] = pa[i]/fsa[i]
        }
    }
    fmt.Printf("PA %+v\n",pa)
}
func select_action() {
    action = 0
    if rand.Float64() < parameters.Prob_explor {
        //explore
        fmt.Println("explore")
        action = rand.Intn(len(possible_actions))
    } else {
        //exploit
        fmt.Println("exploit")
        get_max_pa()
    } //end of exploit
} //end of select_action

func get_max_pa() {
    pa_max := -999.9
    for k,v := range pa {
        fmt.Printf("k %3d v %f \n",k,v)
        if v > pa_max {
             action = k
             pa_max = v
        }
    } //end of loop
} //end of get_max

func generate_action_set() {
   A = nil
   i := 0
   for i=0;i<len(m);i++ {
      if m[i].Action == action {
         A = append(A,m[i])
      }
   }
} //end of generate action set
