package main

import ( 
    "fmt"
    "math/rand"
)

const dont_care = '#'
var match_set = map[string]Classifier{}


func generate_match_set() {
    get_match_set()
    for covering_is_required() {
        new_rule := cover()
        insert_into_population(new_rule)
        get_match_set()
    }
}

func covering_is_required() bool {
    return len(match_set) < p.Minimum_actions
}

func get_match_set() {

    for k,_ := range match_set {
        delete(match_set,k)
    }
    
    for k,_ := range possible_actions {
        possible_actions[k] = 0
    }

    for k,v := range population {
        if does_match(state_string,v.Condition) {
            match_set[k] = v
            possible_actions[v.Action]++
        }
    }
} //end of get_match_set

func cover() Classifier {
    var c Classifier
    action_candidates := get_action_candidates()
    index := rand.Intn(len(action_candidates))
    action = action_candidates[index]
    c = make_classifier(state_string,action)
    return c 
}

func get_action_candidates() []int {
    var action_candidates []int

    for k,_ := range possible_actions {
        possible_actions[k] = 0
    }
    for _,v := range match_set {
            possible_actions[v.Action]++
    }
   
 
    for k,_ := range possible_actions {
            if possible_actions[k]== 0 {
               action_candidates = append(action_candidates,k)
            }
    }
    if len(action_candidates) == 0 {
         for k,_ := range possible_actions {
             action_candidates = append(action_candidates,k)
         }
    }
    return action_candidates
} //end of get_action_candidates


func does_match(state_string string,cc string) bool {

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

func dump_match_set() {
     fmt.Println("DUMP MATCH SET \n")
     fmt.Printf( "%s -- \n",state_string)
     kntr := make(map[string]int)
     kntr2 := make(map[int]int)

     for _,v := range match_set {
       kntr[v.Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("%s %d\n",k,v)
     }

     for _,v := range match_set {
       kntr2[v.Action]++
     }
     fmt.Printf("Action Counts\n")

     for k,v := range kntr2 {
        fmt.Printf("ACTIONS %d %d\n",k,v)
     }

} //end of dump_match_set
