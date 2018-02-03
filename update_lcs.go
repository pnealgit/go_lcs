package main

import (
	"fmt"
        "math"
        "os"
)

var big_p float64
var max_pa float64
var state_string_minus_one string
 
func update_lcs() {
        
    fmt.Printf("A1 LENGTH : %d\n",len(A1))
    if len(A1) > 0 {
        big_p = reward_minus_one + parameters.Gamma * max_pa
        fmt.Printf("BIG_P: %f R1 : %f GAMMA: %f MAXPA %f \n",big_p,reward_minus_one,parameters.Gamma,max_pa)
        if max_pa <= 0.0 {
           fmt.Printf("max_pa <= zero %f\n",max_pa)
           fmt.Printf("PREDICTIVE ARRAY IS: %+v\n",prediction_array)
           fmt.Printf("Exiting from update_lcs\n")
           os.Exit(2)
        }
        update_set()
        run_ga()
     }

       for k,_ := range A1 {
           delete(A1,k)
       }

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

    for k,v := range A1 {
        //experience
        v.Exp++
        A1[k] = v
    }


    //action set size estimate
    //'c.n' is numerosity
    numerosity_sum := 0.0
    for _,v := range A1 {
        numerosity_sum += v.n
    } 

    for k,v := range A1 {
        //experience
        v.Exp++

        delta := big_p - v.p 
        ad := math.Abs(delta)
        if v.Exp < 1.0/parameters.Beta {
            v.p += delta/v.Exp    //prediction
            v.Epsilon += (ad - v.Epsilon)/v.Exp   //prediction_error
            v.as += (numerosity_sum - v.as)/v.Exp      //asction set size
        } else {
            v.p += parameters.Beta * delta
            v.Epsilon += parameters.Beta * (ad - v.Epsilon)
            v.as += parameters.Beta * (numerosity_sum - v.as)
        }
        if v.as <= 0.0 {
            fmt.Printf("Action set size < zero %f numerosity_sum %f \n",v.as,numerosity_sum)
            fmt.Printf("Exiting from update_set \n")
            os.Exit(4)
        }
        if v.Epsilon <= 0.0 {
            fmt.Printf("Epsilon < zero %f\n",v.Epsilon)
            fmt.Printf("Exiting from update_set \n")
            os.Exit(4)
        }

        A1[k] = v

    } //end of loop on lassifiers in action set

    update_fitness()
    if parameters.Do_action_set_subsumption {
       do_action_set_subsumption()
    }
}

func update_fitness() {
    accuracy_sum := 0.0
    accuracy := make(map[string]float64)

    for k,v := range A1 {
        if v.Epsilon < parameters.Epsilon_zero {
            accuracy[k] = 1.0 
        } else {
            tspread := (v.Epsilon/parameters.Epsilon_zero)
            spread := math.Pow(tspread,parameters.V)
            accuracy[k] = parameters.Alpha * spread
            if accuracy[k] <= 0.0 {
                fmt.Printf("BAD ACCURACY\n")
                fmt.Printf("tspread %f spread %f accuracy[k] %f \n",tspread,spread,accuracy[k])
                fmt.Printf("EXITING FROM UPDATE FITNESS \n")
                os.Exit(8)
            }
        } 
        accuracy_sum += accuracy[k] * v.n
    } //end of loop
    if accuracy_sum <= 0.0 {
        fmt.Printf("Bad accuracy_sum %f\n",accuracy_sum)
        for k,v := range accuracy {
            fmt.Printf("key: %s accuracy %f\n",k,v)
        }
        os.Exit(9)
    }

    for k,v := range A1 {
        v.F += parameters.Beta * (accuracy[k] * v.n/accuracy_sum - v.F)
        if math.IsNaN(v.F) {
             fmt.Printf("F not a number in update fitness \n")
             os.Exit(7)
        } 
        A1[k] = v
    }
    
} //end of update fitness
  
func copy_action_sets(a_to *[]Classifier,a_from *[]Classifier) {
     *a_to = *a_from
}
 
