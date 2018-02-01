package main

import ( 
    "math/rand"
)


func generate_match_set(state_string string) {
    m.Classifiers = nil

    i := 0
    for i = 0; i<len(p.Classifiers);i++ {
        if is_match(state_string,p.Classifiers[i]) {
            m.Classifiers = append(m.Classifiers,p.Classifiers[i])
        }
    }

    if len(m.Classifiers) <= parameters.Theta_mna {
        cover(state_string)
    }
} //end of generate_match_set

func is_match(state_string string,classifier Classifier) bool {
    cc := classifier.Condition

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


func cover(state_string string) {

    //setup temp map for counting what actions are used in covering
    var possibles  []int
    possibles = get_possibles()

    //i := 0
    var c Classifier
    //for len(m.Classifiers) < parameters.Theta_mna {
    for k,_ := range possibles {
       c = make_classifier(state_string,k)
       p.Classifiers = append(p.Classifiers,c)
       m.Classifiers = append(m.Classifiers,c)
    }
} //end of cover

func make_classifier(state_string string,action int ) Classifier {
    var c Classifier

    //maybe add dont cares
    s := add_dont_cares(state_string)
 
    c.Condition = string(s)
    c.Action = action
    c.p = parameters.Prediction_initial
    c.Epsilon = parameters.Epsilon_initial
    //c.F = parameters.Fitness_initial
    c.F = .00001
    c.Exp = 0.0
    c.ts = my_time
    c.as = 0.0 //average action set size
    c.n = 1.0                //numerosity
    return c
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

func get_possibles() []int {
    //this is inefficent for large number possible actions
    var possibles []int
    for k,_ := range possible_actions {
        possibles = append(possibles,k)
    }
    return possibles
} //end of get_possibles


