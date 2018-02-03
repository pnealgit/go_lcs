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
    Condition    string       //what was sent from sensors
    Action       int          //what was result of thinking about condition
    p            float64      //predicted payoff
    Epsilon float64
    F            float64      //fitness
    Exp   float64 
    ts    float64
    as float64                   //average selection set size
    n   float64                  //numerosity
}

var p = map[string]Classifier{}
var m = map[string]Classifier{}
var A = map[string]Classifier{}
var A1 = map[string]Classifier{}

var my_time float64
var action int

var state_string string
var prediction_array  map[int]float64 //prediction array

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
        if len(state_string) < 5 {
           fmt.Printf("State string length < 5 %d \n",len(state_string))
           fmt.Printf("Exiting ...  ")
           os.Exit(6)
        }
        reward = state_record.Reward

        my_time++ 
        generate_match_set(state_string)
        generate_prediction_array()
        select_action()
        generate_action_set()

        angle_record.Angle = action
        angle_record.Status = "angle"
	message,err  = json.Marshal(angle_record)
        if err != nil {
                fmt.Println("bad angles Marshal")
        }
	return message
} //end of do_update

func convert_state_input(state_record State_record) string {
    var s bytes.Buffer
    si := ""
    i := 0
    si = fmt.Sprintf("%03b",state_record.State[0])
    s.WriteString(si)

    si = fmt.Sprintf("%06b",state_record.State[1]%20)
    s.WriteString(si)

    si = fmt.Sprintf("%06b",state_record.State[2]%20)
    s.WriteString(si)

    for i=3;i<len(state_record.State);i++ {
        si = fmt.Sprintf("%04b",state_record.State[i])
        s.WriteString(si)
    }
    return s.String()
}

func dump_population() {

     kntr := make(map[string]int)
     for _,v := range p {
       kntr[v.Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("%d %s\n",v,k)
     }

} //end of dump

func dump_match_set(state_string string) {
     fmt.Println("DUMP MATCH SET \n")
     fmt.Printf( "%s\n",state_string)
     kntr := make(map[string]int)
     kntr2 := make(map[int]int)

     for _,v := range m {
       kntr[v.Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("Match set %d %s\n",v,k)
     }

     for _,v := range m {
       kntr2[v.Action]++
     }
     fmt.Printf("Action Counts\n")

     for k,v := range kntr2 {
        fmt.Printf("ACTIONS %d %d\n",k,v)
     }

} //end of dump_match_set

func generate_prediction_array() {
    prediction_array = make( map[int]float64)
    fsa := make(map[int]float64) 
    for k,_ := range prediction_array {
        delete(prediction_array,k)
    }

    for _,v := range m {
        prediction_array[v.Action] += v.p * v.F
        fsa[v.Action] += v.F
    } 
    for k,_ := range possible_actions {
        if fsa[k] > 0.0 {
            prediction_array[k] = prediction_array[k]/fsa[k]
        }
    }
    fmt.Printf("PA %+v\n",prediction_array)
}
func select_action() {
    action = 0
    if rand.Float64() < parameters.Prob_explor {
        //explore
        action = rand.Intn(len(possible_actions))
        fmt.Printf("explore action is %d \n",action)
        get_max_pa()
    } else {
        //exploit
        get_max_pa()
        get_max_pa_action()
        fmt.Printf("exploit action is %d \n",action)
    } //end of exploit
} //end of select_action

func get_max_pa() {

    max_pa = -999.9
    for _,v := range prediction_array {
        if v > max_pa {
             max_pa = v
        }
    } //end of loop
    fmt.Printf("IN GET_MAX_PA pa is %+v MAX: %f ",prediction_array ,max_pa )

} //end of get_max

func get_max_pa_action() {
    zmax_pa := -999.9
    for k,v := range prediction_array {
        if v > zmax_pa {
             action = k
             zmax_pa = v
        }
    } //end of loop
} //end of get_max

func generate_action_set() {
   for k,_ := range A {
       delete(A,k)
   }
   
   for k,v := range m {
      if v.Action == action {
         A[k] = v
      }
   }
} //end of generate action set

