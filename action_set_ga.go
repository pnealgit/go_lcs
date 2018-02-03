package main

import ( 
    "fmt"
    "math/rand"
)
 
func run_ga() {
    fmt.Println("in run_ga")
    sum := 0.0
    nsum:= 0.0
    for _,v := range A1 {
        sum += float64(v.ts) * v.n
        nsum += v.n
    }
    avg := sum/nsum;
    if my_time - avg < parameters.Theta_ga {
       fmt.Println("NOT ENOUGH TO DO GA RETURNING ")
       return
    }

    fmt.Println("looks like we do ga")

    var parent0 Classifier
    var parent1 Classifier
    var child []Classifier

    l := len(A1)
    fmt.Printf("before loop length action set classifiers %d\n",l)
 
    for k,v := range A1 {
          v.ts = my_time
          A1[k] = v
          p0x := select_offspring()
          p1x := select_offspring()
fmt.Printf("p0x %+v\n",p0x);
fmt.Printf("p1x %+v\n",p1x);

          parent0 = A1[p0x]
          parent1 = A1[p1x]
fmt.Printf("p0 %+v\n",parent0);
fmt.Printf("p1 %+v\n",parent1);
          
          child = append(child,Classifier{})
          child = append(child,Classifier{})
          my_copy(&parent0,&child[0])
          my_copy(&parent1,&child[1])

          child[0].n = 1
          child[1].n = 1
          child[0].Exp = 0
          child[1].Exp = 0
fmt.Printf("c0 %+v\n",child[0]);
fmt.Printf("c1 %+v\n",child[1]);

          if rand.Float64() < parameters.Chi {
              apply_crossover(&child[0],&child[1])
              child[0].p = (parent0.p + parent1.p)/2
              child[0].Epsilon = (parent0.Epsilon + parent1.Epsilon)/2
              child[0].F = (parent0.F + parent1.F)/2
              child[1].p = child[0].p
              child[1].Epsilon = child[0].Epsilon
              child[1].F = child[0].F
          }

          child[0].F = child[0].F * 0.1
          child[1].F = child[1].F * 0.1
     
          i := 0 
          k := ""
          for i=0;i<len(child);i++ {
              child[i].Condition = apply_mutation(child[i].Condition)
              if parameters.Do_ga_subsumption {
                  if does_subsume(parent0,child[i]) {
                     parent0.n++
                  } else if does_subsume(parent1,child[i]) {
                     parent1.n++
                  } else { 
                     k = make_key(child[i].Condition,child[i].Action)
                     p[k] = child[i]
                  }
               } else {
                  k = make_key(child[i].Condition,child[i].Action)
                  p[k] = child[i]
               } //end of if on do_ga

               //this means clean up population
               delete_from_population()
           } //end of loop on child 
    } //end of loop on classifiers in A1
} //end of run_ga

func select_offspring()  string {

    fmt.Printf("A1: %+v\n",A1)

    fmt.Printf("\nIN SELECT len(A1) %d \n",len(A1))

    fitness_sum := 0.0
    for _,v := range A1 {
         fitness_sum += v.F
    }

    choice_point := rand.Float64() * fitness_sum
    fmt.Printf("select off  fitness_sum %f choice_point %f \n",fitness_sum,choice_point)
    fitness_sum = 0.0
    pindex := ""
    for k,v := range A1 {
         fitness_sum += v.F
         if fitness_sum > choice_point {
           pindex = k
           fmt.Printf("key for parent choicepoint : %s \n",pindex)
           break
         } //end of if on fitness
    } //end of k,v loop 
        
    return pindex  
} //end of select_offspring

func apply_crossover(c0,c1 *Classifier) {

    fmt.Printf("\n apply crossover c0 %+v ",c0)
    fmt.Printf("apply crossover c1 %+v ",c1)
    
    cut_point := rand.Intn(len(c0.Condition))
    //var tmp1 []byte
    var tmp1 string
    var tmp2 string
    //var tmp2 []byte
    
    tmp1 = c0.Condition[cut_point:]
    tmp2 = c1.Condition[cut_point:]

    c0.Condition = c0.Condition[0:cut_point] + tmp2
    c1.Condition = c1.Condition[0:cut_point] + tmp1
}

func apply_mutation(condition string) string {
    fmt.Println("IN MUTATION")
    fmt.Printf("before %s\n",condition)
    fmt.Printf("state  %s\n",state_string)
    i := 0
    temp := []byte(condition)
    fmt.Printf("after convert %s Mu: %f \n",temp,parameters.Mu)
    fmt.Printf("LENGTH OF TEMP: %d \n",len(temp))

    for i=0;i<len(temp);i++ {
         if rand.Float64() < parameters.Mu {
             if temp[i] == '#' {
                temp[i] = state_string[i]
             } else {
                temp[i] = '#'
             }
         }
    } //end of loop
    fmt.Printf("after  %s \n",string(temp))
    condition = string(temp)
    return condition 
} //end of apply mutation


func insert_in_population(cl Classifier) {
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

func do_action_set_subsumption() {
    var cl Classifier

    c_num_dont_cares := 0
    cl_num_dont_cares := 0
    
    for _,v := range A1 {
       c_num_dont_cares = get_num_dont_care(v.Condition)
       cl_num_dont_cares = get_num_dont_care(cl.Condition)
 
       if could_subsume(v) {

          //line break on operand for continuation
          if (cl.Condition == "" || c_num_dont_cares > cl_num_dont_cares ) ||
             (c_num_dont_cares == cl_num_dont_cares && rand.Float64() < 0.5)  {
                  cl = v
          }
       } //end of could subsume if
    } //end of loop on A1

    if cl.Condition != "" {
        for k,v := range A1 {
          if is_more_general(cl.Condition,v.Condition) {
                cl.n = cl.n + v.n
                delete(A1,k)
                delete(p,k)
          }
        }
    } //end of if on nil
} //end of do_A1_subsumption

func could_subsume(cl Classifier) bool {
    if cl.Exp > parameters.Theta_sub {
        if cl.Epsilon < parameters.Epsilon_zero {
           return true
        }
    }
    return false
}

func is_more_general(clgen string,clspec string) bool {
    
    if get_num_dont_care(clgen) <= get_num_dont_care(clspec) {
         return false
    }

    i := 0
    for i=0;i<len(clgen);i++ {
         if clgen[i] != dont_care && clgen[i] != clspec[i] {
            return false
         }
    }
    return true
} //end of is_more_general

func does_subsume(clsub Classifier,cltos Classifier) bool {
    if clsub.Action == cltos.Action {
       if could_subsume(clsub) {
           if is_more_general(clsub.Condition,cltos.Condition) {
              return true
           }
       }
    }
    return false
}

func my_copy(from *Classifier, to *Classifier) {
    *to = *from
}

func get_num_dont_care(condition string) int {
     i:= 0
     knt:= 0
     for i=0;i<len(condition);i++ {
        if condition[i] == '#' {
           knt++
        }
     }
     return knt
}

func remove_classifier_from_A1(cl Classifier) {
     fmt.Println("NOT FINISHED !!!!!! removing from action set ")
}
