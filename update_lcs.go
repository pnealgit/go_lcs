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
       
    big_p = 0.0 
    if len(A1) > 0 {
       
        big_p = reward_minus_one + parameters.Gamma * max_pa
        if max_pa <= 0.0 {
           fmt.Printf("max_pa <= zero %f\n",max_pa)
           fmt.Printf("PREDICTIVE ARRAY IS: %+v\n",prediction_array)
           fmt.Printf("Exiting from update_lcs\n")
           os.Exit(2)
        }
        update_set()
        //run_ga()
     }
       //before := len(A1)
       //al := len(A)
       for k,_ := range A1 {
           delete(A1,k)
       }

       for k,v := range A {
            A1[k] = v
       }

       //after := len(A1)
       //if before != after {
        //   fmt.Printf("AT UPDATE LEN(A1) before %d LEN(A) %d LEN(A1) after %d\n",before,al,after)
       //}

       reward_minus_one = reward
       state_string_minus_one =  state_string
} //end of update_lcs

func update_set() {

    numerosity_sum := 0.0
    for _,v := range A1 {
        numerosity_sum += v.n
    } 

    for k,v := range A1 {
        //experience
        v.Exp++
        //action set size estimate
        //'c.n' is numerosity

        if v.Exp < 1.0/parameters.Beta {
            v.p += (big_p -v.p)/v.Exp    //prediction
        } else {
            v.p += parameters.Beta * (big_p - v.p)
        }

        //prediction error
        if v.Exp < 1.0/parameters.Beta {
            v.Epsilon += (math.Abs((big_p-v.p)) - v.Epsilon)/v.Exp   //prediction_error
        } else {
            v.Epsilon += parameters.Beta * (math.Abs((big_p-v.p)) - v.Epsilon)
        }

        if v.Exp < 1.0/parameters.Beta {
            v.as += (numerosity_sum - v.as)/v.Exp      //asction set size
        } else {
            v.as += parameters.Beta * (numerosity_sum - v.as)
        }

        if v.as <= 0.0 {
            fmt.Printf("Action set size < zero %f numerosity_sum %f \n",v.as,numerosity_sum)
            fmt.Printf("Exiting from update_set \n")
            os.Exit(7)
        }
        if v.Epsilon < 0.0 {
            fmt.Printf("Epsilon < zero %f\n",v.Epsilon)
            fmt.Printf("BIG_P %f p %f \n",big_p,v.p)
            fmt.Printf("reward1 %f Gamma %f max_pa %f \n",reward_minus_one, parameters.Gamma, max_pa)
            fmt.Printf("Exp %f Beta %f \n",v.Exp,parameters.Beta)
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
 
