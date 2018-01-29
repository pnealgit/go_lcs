package main

import (
	"encoding/json"
	"fmt"
        "bytes"
        "math/rand"
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

type Parameters struct {
    N	int
    Beta float64
    Alpha float64
    Gamma float64
    Epsilon_zero float64
    Nu  float64
    Theta_ga int
    Chi   float64
    Mu    float64
    Theta_del int
    Sigma float64
    Theta_sub float64
    Prob_sharp float64
    Prediction_initial float64
    Epsilon_initial float64
    Fitness_initial float64
    Prob_explor float64
    Theta_mna int
    Do_ga_subsumption bool
    Do_action_set_subsumption bool
}
 

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

var parameters Parameters

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
    
        get_match_set(state_string,state_record.Reward)
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

func get_match_set(state_string string,reward float64) {
    m.Classifiers = nil
    i := 0
    for i = 0; i<len(p.Classifiers);i++ {
        if is_match(state_string,p.Classifiers[i]) {
            m.Classifiers = append(m.Classifiers,p.Classifiers[i])
        }
    }
    if len(m.Classifiers) <= 0 {
        cover(state_string,reward)
    }

} //end of get_match_set

func is_match(state_string string,classifier Classifier) bool {
    cc := classifier.Condition
    //fmt.Printf("cc: %s\n",cc)
    //fmt.Printf("ss: %s\n",state_string)

    dont_care := byte('#')
    i := 0
    for i=0;i<len(state_string);i++ {
        if state_string[i] == cc[i] || 
           cc[i] == dont_care {
              //do nothing
        } else {
           return false
        }
    }
    return true
} //end of is_match

func get_action_set() {
}

func cover(state_string string,reward float64) {
    i := 0
    var c Classifier
    // 0 and 1 temporary
    for i=0;i<2;i++ {
       c = make_classifier(state_string,reward,i)
       if len(state_string) != len(c.Condition) {
           fmt.Printf("cc: %s",state_string)
           fmt.Printf("ss: %s",c.Condition)
       }


       p.Classifiers = append(p.Classifiers,c)
       m.Classifiers = append(m.Classifiers,c)
    }
} //end of cover

func make_classifier(state_string string, reward float64,action int ) Classifier {
    var c Classifier

    //maybe add dont cares
    s := add_dont_cares(state_string)
 
    c.Condition = string(s)
    c.Action = action
    c.Prediction = parameters.Prediction_initial
    c.Prediction_error = parameters.Epsilon_initial
    c.Fitness = parameters.Fitness_initial
    c.Experience = 0.0
    c.Last_ga_time_stamp = 0
    c.Average_action_set_size = 0
    c.Numerosity = 0
    return c
}


func get_parameters() {

        parameters.N = 100
        parameters.Prob_sharp = 0.20
        parameters.Prediction_initial = 10.0 //make this relative to reward
        parameters.Gamma = 0.71
        parameters.Theta_ga = 25 //adjust
        parameters.Chi = 0.5
        parameters.Mu = 0.025
        parameters.Theta_del = 20
        parameters.Sigma = 0.1
        parameters.Theta_sub = 20
        parameters.Prediction_initial = 0.00001
        parameters.Epsilon_initial = 0.00001
        parameters.Fitness_initial = 0.00001
        parameters.Theta_mna = 2  //equal to the number of actions..could be smaller
        parameters.Prob_explor = 0.5
        parameters.Do_ga_subsumption = true
        parameters.Do_action_set_subsumption = true

        fmt.Printf("PARAMETERS %v ",parameters)
}

func add_dont_cares(state_string string) []byte {
    s := make([]byte,len(state_string))

    copy(s,state_string)
    i := 0
    for i=0;i<len(state_string);i++ {
         if rand.Float64() < parameters.Prob_sharp {
             s[i] = '#'
         } 
    } //end of loop
    return s
} //end of add_dont_cares

func dump_population() {
     i := 0
     for i=0;i<len(p.Classifiers);i++ {
       fmt.Printf("%4d %4d %v\n",i,p.Classifiers[i].Action,p.Classifiers[i].Condition)
     }
} //end of dump


