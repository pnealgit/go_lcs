package main

import(
    "math/rand"
)

func insert_into_population(cl Classifier) {
    k := ""
    k = make_key(cl.Condition,cl.Action)

    value, ok := population[k]
    if ok {
          value.Numerosity++
          population[k] = value
    } else {
          population[k] = cl
    }
} //end of insert_in_population

func delete_from_population() {
    sum := 0.0
    fsum := 0.0
    for _,v := range population {
        sum += v.Numerosity
        fsum += v.Fitness
    }

    if sum < float64(p.Population_size) {
        return
    }

    av_fitness_in_population := 0.0
    av_fitness_in_population = fsum/sum 
    vote_sum := 0.0
    for _,v := range population {
       vote_sum += deletion_vote(v,av_fitness_in_population)
    }

    //fmt.Printf("vote_sum: %f\n",vote_sum)
    dooby := 0.0
    dooby = rand.Float64()
    choice_point := 0.0
    choice_point = dooby * vote_sum
    
    vote_sum = 0.0
    for k,v := range population {
       vote_sum += deletion_vote(v,av_fitness_in_population)
       if vote_sum > choice_point {
           if v.Numerosity > 1.0 {
              v.Numerosity--
              population[k] = v
           } else {
              delete(population,k)
              break
           }
       } //end of if on choice_point
    } //end of loop on k,v
} //end of delete_from_population

func deletion_vote(cl Classifier, av_fit float64) float64 {
    vote := cl.Action_set_size * cl.Numerosity
    duh := cl.Fitness/cl.Numerosity
    if cl.Experience > float64(p.Deletion_threshold) && duh < p.Fitness_threshold * av_fit {
         vote = vote * av_fit/duh
    }
    return vote
} //end of deletion vote

