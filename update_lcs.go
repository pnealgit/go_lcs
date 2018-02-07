package main

import (
        "math"
)

var big_p float64
var max_pa float64
var state_string_minus_one string
 
func update_lcs() {
       
    if len(A1) > 0 {
        distribute_payoff()
        update()
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

func distribute_payoff() {

    big_p = 0.0 
    payoff := reward_minus_one + p.Discount_factor * max_pa
    action_set_size := 0.0
    for _,v := range A1 {
        action_set_size += v.Numerosity
    } 

    for k,rule := range A1 {
        //experience
        rule.Experience++
        update_rate := math.Max(p.Learning_rate, 1.0/rule.Experience)

        rule.Average_reward += (payoff-rule.Average_reward) * update_rate

        rule.Error+= (math.Abs(payoff-rule.Average_reward) - rule.Error) * update_rate

        rule.Action_set_size += (action_set_size - rule.Action_set_size) * update_rate

        A1[k] = rule

    } //end of loop on lassifiers in action set

    update_fitness()
    if p.Do_action_set_subsumption {
       do_action_set_subsumption()
    }
}

func update_fitness() {
    total_accuracy := 0.0
    accuracies := make(map[string]float64)
    accuracy := 0.0

    for k,v := range A1 {
        if v.Error < p.Error_threshold {
            accuracy = 1.0 
        } else {
            ptmp := 0.0
            ptmp = math.Pow((v.Error / p.Error_threshold), -p.Accuracy_power)
            accuracy = p.Accuracy_coefficient * ptmp
        } 
        accuracies[k] = accuracy
        total_accuracy += accuracies[k] * v.Numerosity
    } //end of loop


    for k,v := range A1 {
        accuracy = 0.0
        accuracy = accuracies[k]
        v.Fitness += (p.Learning_rate * (accuracy * v.Numerosity/total_accuracy - v.Fitness))

        A1[k] = v
    }
    
} //end of update fitness
  
func copy_action_sets(a_to *[]Classifier,a_from *[]Classifier) {
     *a_to = *a_from
}
 
