package main

import ( 
    "fmt"
    "math/rand"
)
 
func run_ga() {
    fmt.Println("running ga")
    sum := 0.0
    nsum:= 0.0
    i := 0
    for i=0;i<len(action_set.Classifiers);i++ {
        sum += float64(action_set.Classifiers[i].ts) * action_set.Classifiers[i].n
        nsum += action_set.Classifiers[i].n
    }
    if my_time - (sum/nsum) < parameters.Theta_ga {
       return
    }

    //var child0 Classifier
    //var child1 Classifier
    var parent0 Classifier
    var parent1 Classifier

    for i=0;i<len(action_set.Classifiers);i++ {
          action_set.Classifiers[i].ts = my_time
          parent0 = select_offspring()
          parent1 = select_offspring()
           
          child0 := Classifier{}
          child1 := Classifier{}
          my_copy(&parent0,&child0)
          my_copy(&parent1,&child1)

          child0.n = 1
          child1.n = 1
          child0.Exp = 0
          child1.Exp = 0
          if rand.Float64() < parameters.Chi {
              //child0,child1 = apply_crossover(child0,child1)
              apply_crossover(&child0,&child1)
              child0.p = (parent0.p + parent1.p)/2
              child0.Epsilon = (parent0.Epsilon + parent1.Epsilon)/2
              child0.F = (parent0.F + parent1.F)/2
              child1.p = child0.p
              child1.Epsilon = child0.Epsilon
              child1.F = child0.F
          }

          child0.F = child0.F * 0.1
          child1.F = child1.F * 0.1
     
          i := 0 
          apply_mutation(&child0)
          apply_mutation(&child1)
          if parameters.Do_ga_subsumption {
             work_subsumption(child0,&parent0,&parent1)
             work_subsumption(child1,&parent0,&parent1)
          } else {                         
             insert_in_population(&child0)
             insert_in_population(&child1)
          }
          delete_from_population()
          delete_from_population()
    } //end of loop on classifiers in action_set
} //end of run_ga

func work_subsumption(kid Classifier,parent0 *Classifier,parent1 *Classifier) {
    subsumption := false
    if does_subsume(parent0,kid) {
          parent0.n++ 
          subsumption = true
    }
                
    if does_subsume(parent1,kid) {
          parent1.n++ 
          subsumption = true
    }
                 
    if subsumption == false {
         insert_in_population(&kid)
    }
} //end of work_subsumption

 
func select_offspring() Classifier {

    fitness_sum := 0.0
    i := 0
    for i=0;i<len(action_set.Classifiers);i++ {
         fitness_sum += action_set.Classifiers[i].F
    }
    choice_point := rand.Float64() * fitness_sum
    fitness_sum = 0.0
    pindex := 0
    for i=0;i<len(action_set.Classifiers);i++ {
         fitness_sum += action_set.Classifiers[i].F
         if fitness_sum > choice_point {
           pindex = i 
           break
        } //end of if on fitness
    } //end of loop on i
           
    return action_set.Classifiers[pindex]

} //end of select_offspring

func apply_crossover(cl1 *Classifier,cl2 *Classifier) {
    cut_point := rand.Intn(len(cl1.Condition))
    //var tmp1 []byte
    var tmp1 string
    var tmp2 string
    //var tmp2 []byte
    
    tmp1 = cl1.Condition[cut_point:]
    tmp2 = cl1.Condition[cut_point:]

    cl1.Condition = cl1.Condition[0:cut_point] + tmp2
    cl2.Condition = cl2.Condition[0:cut_point] + tmp1

}

func apply_mutation(cl *Classifier) {
    i := 0
    for i=0;i<len(cl.Condition);i++ {
         if rand.Float64() < parameters.Mu {
             if cl.Condition[i] == '#' {
                cl.Condition[i] = state_string[i]
             } else {
                cl.Condition[i] = '#'
             }
         }
    }
} //end of apply mutation


func insert_in_population(cl *Classifier) {
    i := 0
    for i=0;i<len(p.Classifiers);i++ {
         if p.Classifiers[i].Condition == cl.Condition &&
            p.Classifiers[i].Action == cl.Action {
                p.Classifiers[i].n++
                return
         }
    }
    p.Classifiers = append(p.Classifiers,cl)
} //end of insert_in_population

func delete_from_population() {
    sum := 0.0
    fsum := 0.0
    i := 0
    for i=0;i<len(p.Classifiers);i++ {
        sum += p.Classifiers[i].n
        fsum += p.Classifiers[i].F
    }

    if sum < float64(parameters.N) {
        return
    }

    av_fitness_in_population := sum/fsum 
    vote_sum := 0.0
    for i=0;i<len(p.Classifiers);i++ {
       vote_sum += deletion_vote(p.Classifiers[i],av_fitness_in_population)
    }

    choice_point := rand.Float64() * vote_sum
    vote_sum = 0.0
    for i=0;i<len(p.Classifiers);i++ {
       vote_sum += deletion_vote(p.Classifiers[i],av_fitness_in_population)
       if vote_sum > choice_point {
           if p.Classifiers[i].n > 1 {
              p.Classifiers[i].n--
           } else {
             remove_classifier_from_population(p.Classifiers[i])
           }
       return

       } //end of if on choice_point
    } //end of loop on rule
} //end of delete_from_population



func deletion_vote(cl Classifier, av_fit float64) float64 {
    vote := cl.as * cl.n
    if cl.Exp > float64(parameters.Theta_del) && cl.F/cl.n < parameters.Sigma * av_fit {
         vote = vote * av_fit/(cl.F/cl.n)
    }
    return vote
} //end of deletion vote

func do_action_set_subsumption() {
    cl = make_classifier()
    c_num_dont_care := 0
    cl_num_dont_care := 0
    i := 0
    for i=0;i<len(action_set.Classifiers);i++ {
       c_num_dont_care = get_num_dont_care(action_set.Classifiers[i].Condition)
       cl_num_dont_care = get_num_dont_care(cl.Condition)
 
       if could_subsume(c) {
          if (cl == nil || c_num_dont_care > cl_num_dont_care) {
              cl = c
          }

          if c_num_dont_care == cl_num_dont_care && randFloat64() < 0.5 {
              cl = c
          }
       }
    } //end of loop on c

    if cl != nil {
        for c,_ := range action_set {
            if is_more_general(cl,c) {
                cl.n = cl.n + c.n
                remove_classifier_from_action_set(c)
                remove_classifier_from_population(c)
            }
        }
    } //end of if on nil

} //end of do_action_set_subsumption

func could_subsume(cl Classifier) {
    if cl.exp > parameters.Theta_sub {
        if cl.Epsilon < parameters.Epsilon_zero {
           return true
        }
    }
    return false
}

func is_more_general(clgen Classifier,clspec Classifier) {
    
    if get_num_dont_care(clgen.Condition) <= get_num_dont_care(clspec.Condition) {
         return false
    }

    i := 0
    for i=0;i<len(clgen.Condition);i++ {
         if clgen.Condition[i] != '#' && clgen.Condition[i] != clspec.Condition[i] {
            return false
         }
    }
    return true
} //end of is_more_general

func does_subsume(clsub,cltos) bool {
    if clsub.A == cltos.A {
       if could_subsume(clsub) {
           if is_more_general(clsub,cltos) {
              return true
           }
       }
    }
    return false
}

func my_copy(from *Classifier, to *Classifier) {
    *to = *from
}

func remove_classifier_from_population(cl Classifier) {
    fmt.Printf("removing %v\n",cl)
}

    
func get_num_dont_care(cl Classifier) int {
     i:= 0
     knt:= 0
     for i=0;i<len(cl.Condition);i++ {
        if cl.Condition[i] == '#' {
           knt++
        }
     }
     return knt
}

