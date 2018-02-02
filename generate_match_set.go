package main

import ( 
    "math/rand"
    "fmt"
)

const dont_care = '#'

func generate_match_set(state_string string) {
    fmt.Printf("in match set string is: %s \n",state_string)

    m = nil
    i := 0
    for i = 0; i<len(p);i++ {
        if len(p[i].Condition) < 5 {
           fmt.Printf("P - I: %4d %s %s \n",i,state_string,p[i].Condition)
           fmt.Printf("cl %+v\n",p[i])
        }

        if is_match(state_string,p[i].Condition) {
            m = append(m,p[i])
        }
    }
    fmt.Printf("after appending to me len m is %d \n",len(m))

    if len(m) <= parameters.Theta_mna {
        fmt.Println("going to cover")
        cover(state_string)
        fmt.Println("back from cover")
    }
} //end of generate_match_set

func is_match(state_string string,cc string) bool {
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


func cover(state_string string) {

    //setup temp map for counting what actions are used in covering
    var c Classifier

    i := 0
    for i=0;i<len(possible_actions);i++ {
        possible_actions[i] = 0
    }
    fmt.Printf("possibles for cover: %+v\n",possible_actions)
    for i=0;i<len(m);i++ {
       possible_actions[m[i].Action]++
    }
    fmt.Printf("possible_actions after m: %+v\n",possible_actions)
 
    for k,v := range possible_actions {
       if v <= 0 {
           c = make_classifier(state_string,k)
           fmt.Printf("MAKE C: %+v\n",c) 
           p = append(p,c)
           m = append(m,c)
           fmt.Printf("After cover append len P %d\n",len(p))
       }
    }
} //end of cover

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


