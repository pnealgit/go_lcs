package main

import (
	"encoding/json"
	"fmt"
        "bytes"
        "os"
)

type State_record struct {
	State  []int
        Reward float64
}

type Angle_record struct {
    Angle float64
}

var state_record State_record
 
var angle_record Angle_record

type Classifier struct {
    Condition    string
    Action       int
    Prediction   float64
    Prediction_error float64
    Fitness      float64
    Experience   float64
    Last_ga_time_stamp   int
    Average_action_set_size int
    Numerosity   int
}

type Population struct {
    Classifiers []Classifier
}

type Match_set struct {
    Classifiers []Classifier
}

type Action_set struct {
    Classifiers []Classifier
}

var p Population
var m Match_set
var a Action_set
var a_minus Action_set

func do_lcs( message []byte) []byte {
       
	var err error
        var state_record State_record

	jerr := json.Unmarshal(message, &state_record)
	if jerr != nil {
		fmt.Println("error on state unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr
        state_string := ""
        state_string = convert_input(state_record)
        if len(state_string) != 36 {
            fmt.Printf("length is : %4d\n",len(state_string))
            fmt.Printf("%v\n",state_record)
            fmt.Printf("%s\n",state_string)
            os.Exit(2)
        }
    
        generate_match_set(state_string,state_record.Reward)
        generate_prediction_array()

        if len(p.Classifiers) >= parameters.N {
             dump_population()
             os.Exit(3)
        }
        get_action_set()
        

        angle_record.Angle = 1.0
	message,err  = json.Marshal(angle_record)
        if err != nil {
                fmt.Println("bad angles Marshal")
        }
	return message
} //end of do_update

func think(state_record State_record) float64 {
	return 1.0	
} //end of think

func convert_input(state_record State_record) string {
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
     for i=0;i<len(p.Classifiers);i++ {
       fmt.Printf("%4d %4d %v\n",i,p.Classifiers[i].Action,p.Classifiers[i].Condition)
     }
} //end of dump

func generate_prediction_array() {
    fmt.Println("generating prediction_array")
}
func get_action_set() {
}

