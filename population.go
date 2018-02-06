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

    av_fitness_in_population := 0.0
    av_fitness_in_population = fsum/sum 
    vote_sum := 0.0
    for _,v := range p {
       vote_sum += deletion_vote(v,av_fitness_in_population)
    }

    //fmt.Printf("vote_sum: %f\n",vote_sum)
    dooby := 0.0
    dooby = rand.Float64()
    choice_point := 0.0
    choice_point = dooby * vote_sum
    //fmt.Printf("vote_sum: %f dooby %f choice_point: %f\n",vote_sum,dooby,choice_point)
    
    vote_sum = 0.0
    for k,v := range p {
       vote_sum += deletion_vote(v,av_fitness_in_population)
       if vote_sum > choice_point {
           if v.n > 1.0 {
              v.n--
              p[k] = v
           } else {
              delete(p,k)
              break
           }
       } //end of if on choice_point
    } //end of loop on k,v
} //end of delete_from_population

func deletion_vote(cl Classifier, av_fit float64) float64 {
    vote := cl.as * cl.n
    duh := cl.F/cl.n
    if cl.Exp > float64(parameters.Theta_del) && duh < parameters.Sigma * av_fit {
         vote = vote * av_fit/duh
    }
    return vote
} //end of deletion vote

