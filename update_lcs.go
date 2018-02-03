package main

import (
	"fmt"
        "math"
)

var big_p float64
var max_pa float64
var state_string_minus_one string
 
func update_lcs() {
    fmt.Printf("in UPDATE \n")

    if len(A1) > 0 {
        fmt.Printf("A1 not nil LENGTH : %d\n",len(A1))
        big_p = reward_minus_one + parameters.Gamma * max_pa
        fmt.Printf("BIG_P: %f R1 : %f GAMMA: %f MAXPA %f \n",big_p,reward_minus_one,parameters.Gamma,max_pa)

        update_set()
        run_ga()
        //to from
     }
fmt.Printf("BEFORE ACTION SET COPY: LEN A1: %d LEN A %d \n",len(A1),len(A))
       A1 = make(map[string]Classifier)
       for k,v := range A {
            A1[k] = v
       }
fmt.Printf("AFTER ACTION SET COPY: LEN A1: %d LEN A %d \n",len(A1),len(A))

     reward_minus_one = reward
     state_string_minus_one =  state_string
} //end of update_lcs

func update_set() {
    //reward is 'P'  ... big_p
    //reward from the environment is reward

    fmt.Printf("UPDATE A1 BIG_P IS : %f REWARD %f \n",big_p,reward_minus_one)

    //action set size estimate
    //'c.n' is numerosity
    numerosity_sum := 0.0
    for _,v := range A1 {
        numerosity_sum += v.as - v.n
    } 

    for k,v := range A1 {
        //experience
        v.Exp++

        //prediction
        delta := reward - v.p 
        if v.Exp < 1.0/parameters.Beta {
            v.p += delta/v.Exp
        } else {
            v.p += parameters.Beta * delta
        }
        fmt.Printf("AFTER A1.P : %f \n",v.p )
       
        //prediction_error
        ad := math.Abs(delta)
        if v.Exp < 1.0/parameters.Beta {
            v.Epsilon += 
                 (ad - v.Epsilon)/v.Exp
       } else {
            v.Epsilon += parameters.Beta * (ad - v.Epsilon)
       }

       // this might have to be done with numerosity_sum computed after n has been updated

       //action set size estimate
        if v.Exp < 1.0/parameters.Beta {
            v.as += numerosity_sum /v.Exp
        } else {
            v.as += parameters.Beta * numerosity_sum
        }
        A1[k] = v

    } //end of loop on lassifiers in action set

    update_fitness()
}

func update_fitness() {
    accuracy_sum := 0.0
    var accuracy map[string]float64

    accuracy = make(map[string]float64)
    fmt.Printf("IN UPDATE FITNESS \n")
    fmt.Printf("LEN A1 : %d\n",len(A1))

    for k,v := range A1 {
        if v.Epsilon < parameters.Epsilon_zero {
            accuracy[k] = 1.0 
        } else {
            tspread := (v.Epsilon/parameters.Epsilon_zero)
            spread := math.Pow(tspread,parameters.V)
            accuracy[k] = parameters.Alpha * spread
        } 
        accuracy_sum += accuracy[k] * v.n
    } //end of loop
    
    for k,v := range A1 {
        v.F += parameters.Beta * (accuracy[k] * v.n/accuracy_sum) - v.F
        A1[k] = v
    }
    
} //end of update fitness
  
func copy_action_sets(a_to *[]Classifier,a_from *[]Classifier) {
     *a_to = *a_from
}
 
