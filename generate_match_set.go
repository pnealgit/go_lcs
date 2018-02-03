package main

import ( 
    "math/rand"
    "fmt"
    "strconv"
    "os"
)

const dont_care = '#'

func generate_match_set(state_string string) {
    m = make(map[string]Classifier)
    fmt.Printf("M: %+v \n",m)
 
    for k,v := range p {
        if is_match(state_string,v.Condition) {
            m[k] = v
        }
    }
   
    fmt.Printf("after appending to m len m is %d \n",len(m))
    for k,v := range m {
       fmt.Printf("M - K %s V: %+v \n",k,v)
    }

    if len(m) <= parameters.Theta_mna {
        fmt.Println("going to cover")
        cover(state_string)
        fmt.Printf("back from cover len M : %d \n",len(m))
    }
} //end of generate_match_set

func is_match(state_string string,cc string) bool {
    i := 0
fmt.Printf("ss: %s\n",state_string)
fmt.Printf("cc: %s\n",cc)

    for i=0;i<len(state_string);i++ {
        if state_string[i] == cc[i] || 
           cc[i] == dont_care {
              //do nothing
        } else {
           return false
        }
    }
    fmt.Printf("MATCH \n")
    return true
} //end of is_match


func cover(state_string string) {

    //setup temp map for counting what actions are used in covering
    var c Classifier
    for k,_ := range possible_actions {
       possible_actions[k] = 0
    }
    for _,v1 := range m {
       possible_actions[v1.Action]++
    }
    fmt.Printf("possible_actions after m: %+v\n",possible_actions)
 
    for k1,v1 := range possible_actions {
       if v1 <= 0 {
           c = make_classifier(state_string,k1)
           fmt.Printf("MAKE C: %+v\n",c) 
           ka := make_key(c.Condition,c.Action)
           
           p[ka] = c
           m[ka] = c
           fmt.Printf("After cover append len P %d\n",len(p))
           fmt.Printf("After cover append len M %d\n",len(m))
       }
    }
} //end of cover

func make_key(condition string, action int) string {
     if len(condition) == 0 {
        fmt.Printf("CONDITION IS %s\n",condition)
        os.Exit(5)
     }
     new_k := condition + strconv.Itoa(action)
     return new_k
}

func make_classifier(state_string string,action int ) Classifier {
    var c Classifier

    //maybe add dont cares
    s := add_dont_cares(state_string)
fmt.Printf("back from add_dont_cares s: %s \n",s)
 
    c.Condition = string(s)
    c.Action = action
    c.p = parameters.Prediction_initial
    c.Epsilon = parameters.Epsilon_initial
    //c.F = parameters.Fitness_initial
    c.F = .00001
    c.Exp = 0.0
    c.ts = my_time
    c.as = 1.0 //average action set size
    c.n = 1.0                //numerosity
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


