package main
import (
    "fmt"
    "os"
    "strconv"
    "math/rand"
)

type Classifier struct {
    Condition    string       //what was sent from sensors
    Action       int          //what was result of thinking about condition
    Average_reward  float64      //predicted payoff
    Error        float64
    Fitness      float64      //fitness
    Experience   float64
    Time_stamp    float64
    Action_set_size float64   //average selection set size
    Numerosity   float64               //numerosity
}

func make_classifier(state_string string,action int ) Classifier {
    var c Classifier

    //maybe add dont cares
    s := add_dont_cares(state_string)
    c.Condition = string(s)
    c.Action = action
    c.Average_reward = p.Initial_prediction
    c.Error = p.Initial_error
    c.Fitness = p.Initial_fitness
    c.Experience = 0.0     //number times in an action set
    c.Time_stamp = my_time
    c.Action_set_size = 1.0      //average action set size
    c.Numerosity = 1.0       //numerosity
    return c
}

func make_key(condition string, action int) string {
     if len(condition) == 0 {
        fmt.Printf("BAD CONDITION LENGTH %d IS %s\n",len(condition),condition)
        os.Exit(5)
     }
     new_k := condition + strconv.Itoa(action)
     return new_k
}
func add_dont_cares(state_string string) []byte {
    s := make([]byte,len(state_string))

    copy(s,state_string)
    i := 0
    for i=0;i<len(state_string);i++ {
         if rand.Float64() < p.Wildcard_probability {
             s[i] = dont_care
         } 
    } //end of loop
    return s
} //end of add_dont_cares


