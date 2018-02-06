package main

import ( 
    "math/rand"
    "fmt"
    "strconv"
    "os"
)

const dont_care = '#'

func generate_match_set(state_string string) {
    //zero out m
    for k,_ := range m {
       delete(m,k)
    }

    for k,_ := range possible_actions {
        possible_actions[k] = 0
    }

    for k,v := range p {
        if does_match(state_string,v.Condition) {
            m[k] = v
            possible_actions[v.Action]++
        }
    }
    sum_pos := 0
    for _,v := range possible_actions {
       if v > 0 {
           sum_pos++
       }
    } 

    if sum_pos < parameters.Theta_mna {
        cover()
    }


} //end of generate_match_set

func does_match(state_string string,cc string) bool {

if len(state_string) != len(cc) {
    fmt.Printf("does_match Lengths do not match len ss : %d len cc %d \n",len(state_string),len(cc))
    fmt.Printf("ss: %s\n",state_string)
    fmt.Printf("cc: %s\n",cc)
    fmt.Printf("Exiting \n");
    os.Exit(3)
}
 
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



func cover() {

    var c Classifier
     
    for k,v := range possible_actions {
       if v <= 0 {
           c = make_classifier(state_string,k)
           ka := make_key(c.Condition,c.Action)
           p[ka] = c
           m[ka] = c
           delete_from_population()
       }
    }
} //end of cover

func make_key(condition string, action int) string {
     if len(condition) == 0 {
        fmt.Printf("BAD CONDITION LENGTH %d IS %s\n",len(condition),condition)
        os.Exit(5)
     }
     new_k := condition + strconv.Itoa(action)
     return new_k
}

func make_classifier(state_string string,action int ) Classifier {
    var c Classifier

    //maybe add dont cares
    s := add_dont_cares(state_string)
    c.Condition = string(s)
    c.Action = action
    c.p = parameters.Prediction_initial
    c.Epsilon = parameters.Epsilon_initial
    c.F = parameters.Fitness_initial
    c.Exp = 0.0     //number times in an action set
    c.ts = my_time
    c.as = 1.0      //average action set size
    c.n = 1.0       //numerosity
    return c
}

func add_dont_cares(state_string string) []byte {
    s := make([]byte,len(state_string))

    copy(s,state_string)
    i := 0
    for i=0;i<len(state_string);i++ {
         if rand.Float64() < parameters.Prob_sharp {
             s[i] = dont_care
         } 
    } //end of loop
    return s
} //end of add_dont_cares


