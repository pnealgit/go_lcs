package main

import(
    "math/rand"
    "os"
    "fmt"
)

func insert_in_population(cl Classifier) {
    if len(cl.Condition) != len(state_string) {
       fmt.Printf("Length of condition %d does not equal length of ss %d\n",len(cl.Condition),len(state_string))
       fmt.Printf("Exiting\n")
       os.Exit(1)
    }

    k := ""
    k = make_key(cl.Condition,cl.Action)

    value, ok := p[k]
    if ok {
          value.n++
          p[k] = value
    } else {
          p[k] = cl
    }
} //end of insert_in_population

func delete_from_population() {
    sum := 0.0
    fsum := 0.0
    for _,v := range p {
        sum += v.n
        fsum += v.F
    }

    if sum < float64(parameters.N) {
        return
    }

    av_fitness_in_population := sum/fsum 
    vote_sum := 0.0
    for _,v := range p {
       vote_sum += deletion_vote(v,av_fitness_in_population)
    }

    choice_point := rand.Float64() * vote_sum
    vote_sum = 0.0
    for k,v := range p {
       vote_sum += deletion_vote(v,av_fitness_in_population)
       if vote_sum > choice_point {
           if v.n > 1 {
              v.n--
              p[k] = v
           } else {
             delete(p,k)
           }
       } //end of if on choice_point
    } //end of loop on k,v
} //end of delete_from_population

func deletion_vote(cl Classifier, av_fit float64) float64 {
    vote := cl.as * cl.n
    if cl.Exp > float64(parameters.Theta_del) && cl.F/cl.n < parameters.Sigma * av_fit {
         vote = vote * av_fit/(cl.F/cl.n)
    }
    return vote
} //end of deletion vote

