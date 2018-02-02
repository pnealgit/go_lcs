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

    if A1 != nil {
        fmt.Printf("A1 not nil LENGTH : %d\n",len(A1))
        big_p = reward_minus_one + parameters.Gamma * max_pa
        fmt.Printf("BIG_P: %f R1 : %f GAMMA: %f MAXPA %f \n",big_p,reward_minus_one,parameters.Gamma,max_pa)

        update_set()
        run_ga()
        //to from
     }
     //copy_action_sets(&A1,&A)
fmt.Printf("BEFORE ACTION SET COPY: LEN A1: %d LEN A %d \n",len(A1),len(A))
       A1 = make([]Classifier,len(A))
       copy(A1,A)
fmt.Printf("AFTER ACTION SET COPY: LEN A1: %d LEN A %d \n",len(A1),len(A))

     reward_minus_one = reward
     state_string_minus_one =  state_string
} //end of update_lcs

func update_set() {
    //reward is 'P'  ... big_p
    //reward from the environment is reward

    fmt.Printf("UPDATE A1 BIG_P IS : %f REWARD %f \n",big_p,reward_minus_one)
    i := 0

    //action set size estimate
    //'c.n' is numerosity
    j := 0
    numerosity_sum := 0.0
    for j=0;j<len(A1);j++ {
        numerosity_sum += A1[j].as - A1[i].n
    } 


    for i=0;i<len(A1);i++ {
        A1[i].Exp++
        fmt.Printf("BEFORE A1.P : %f \n",A1[i].p )
        //prediction
        delta := reward - A1[i].p 
        if A1[i].Exp < 1.0/parameters.Beta {
            A1[i].p += delta/A1[i].Exp
        } else {
            A1[i].p += parameters.Beta * delta
        }
        fmt.Printf("AFTER A1.P : %f \n",A1[i].p )
       
        //prediction_error
        ad := math.Abs(delta)
        if A1[i].Exp < 1.0/parameters.Beta {
            A1[i].Epsilon += 
                 (ad - A1[i].Epsilon)/A1[i].Exp
       } else {
            A1[i].Epsilon += parameters.Beta * (ad - A1[i].Epsilon)
       }

       // this might have to be done with numerosity_sum computed after n has been updated

       //action set size estimate
        if A1[i].Exp < 1.0/parameters.Beta {
            A1[i].as += numerosity_sum /A1[i].Exp
        } else {
            A1[i].as += parameters.Beta * numerosity_sum
        }
    } //end of loop on lassifiers in action set

    update_fitness()
}

func update_fitness() {
    accuracy_sum := 0.0
    var k []float64

    k = make([]float64,len(A1))
    fmt.Printf("IN UPDATE FITNESS \n")
    fmt.Printf("LEN A1 : %d\n",len(A1))

    i := 0
    for i=0;i<len(A1);i++ {
        fmt.Printf("I: %d Epsilon: %f Epsilon_zero %f \n",i,A1[i].Epsilon,parameters.Epsilon_zero)

        if A1[i].Epsilon < parameters.Epsilon_zero {
            k[i] = 1.0 
            
        } else {
            tspread := (A1[i].Epsilon/parameters.Epsilon_zero)
            spread := math.Pow(tspread,parameters.V)
            k[i] = parameters.Alpha * spread
        } 
        fmt.Printf("K[i] %f \n",k[i])
        accuracy_sum += k[i] * A1[i].n
    }
    
    fmt.Printf("K: %+v \n",k)
    fmt.Printf("ACCURACY SUM: %f \n",accuracy_sum)

    for i=0;i<len(A1);i++ {
        fmt.Printf("I: %d A1.F before : %f \n",i,A1[i].F)
        quant := (k[i] * A1[i].n/accuracy_sum) - A1[i].F
    
        A1[i].F = A1[i].F + parameters.Beta * quant
        fmt.Printf("I: %d A1.F after  : %f \n",i,A1[i].F)
    }
} //end of update fitness
  
func copy_action_sets(a_to *[]Classifier,a_from *[]Classifier) {
     *a_to = *a_from
}
 
