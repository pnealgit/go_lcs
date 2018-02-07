package main

import (
	"encoding/json"
	"fmt"
        "os"
        "math/rand"
)

type Sensor struct {
    Xpos   float64
    Ypos   float64
    Status float64
}

type State_record struct {
        Last_angle int
        X    float64
        Y    float64
	Sensors []Sensor
}

type Angle_record struct {
    Status  string
    Angle int
}

var state_record State_record
 
var angle_record Angle_record

var population = map[string]Classifier{}
var A = map[string]Classifier{}
var A1 = map[string]Classifier{}

var my_time float64
var sum_reward float64
var action int

var state_string string
var prediction_array  map[int]float64 //prediction array

var reward float64
var reward_minus_one float64

var canvas_width float64
var canvas_height float64
var canvas_max_distance float64

func get_action( message []byte) []byte {
      
	var err error
        var state_record State_record
        prediction_array = make( map[int]float64)

	jerr := json.Unmarshal(message, &state_record)
	if jerr != nil {
		fmt.Println("error on state unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr
        state_string = ""
        convert_state_input(state_record)
        if len(state_string) < 5 {
           fmt.Printf("State string length < 5 %d \n",len(state_string))
           fmt.Printf("Exiting ...  ")
           os.Exit(6)
        }

        my_time++ 
        sum_reward+= reward
        if (int(my_time) % 100 ) == 0 {
           fmt.Printf("MY_TIME %f POP SIZE: %d SUM REW: %f \n",my_time,len(population),sum_reward)
           sum_reward = 0.0
           dump_match_set()
        }
        generate_match_set()
        generate_prediction_array()
        select_action()
        generate_action_set()
        if len(A) <= 0 {
           fmt.Printf("New action set for action %d is zero\n",action)
           fmt.Printf("Possible_actions is %+v\n",possible_actions)
           fmt.Printf("Prediction array is %+v\n",prediction_array)
           fmt.Printf("Match set size is %d \n",len(match_set))
           dump_match_set()
           os.Exit(2)
        }

        angle_record.Angle = action
        angle_record.Status = "angle"
	message,err  = json.Marshal(angle_record)
        if err != nil {
                fmt.Println("bad angles Marshal")
        }
	return message
} //end of do_update

func dump_population() {

     kntr := make(map[string]int)
     for _,v := range population {
       kntr[v.Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("%d %s\n",v,k)
     }

} //end of dump


func generate_prediction_array() {

    fsa := make(map[int]float64) 
    for k,_ := range prediction_array {
        delete(prediction_array,k)
    }

    for _,v := range match_set {
        prediction_array[v.Action] += v.Average_reward * v.Fitness
        fsa[v.Action] += v.Fitness
    } 
    for k,_ := range prediction_array {
        if fsa[k] > 0.0 {
            prediction_array[k] = prediction_array[k]/fsa[k]
        }
    }
} //end of generate_prediction_array

func select_action() {
    action = 0
    if rand.Float64() < p.Exploration_probability {
        //explore
        action = rand.Intn(len(possible_actions))
        get_max_pa()
    } else {
        //exploit
        get_max_pa()
        get_max_pa_action()
    } //end of exploit
} //end of select_action

func get_max_pa() {

    max_pa = -999.9
    for _,v := range prediction_array {
        if v > max_pa {
             max_pa = v
        }
    } //end of loop

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
   
   for k,v := range match_set {
      if v.Action == action {
         A[k] = v
      }
   }

} //end of generate action set

