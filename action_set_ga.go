package main

import ( 
    "fmt"
    "math/rand"
)
 
func run_ga() {
    fmt.Println("in run_ga")
    sum := 0.0
    nsum:= 0.0
    i := 0
    for i=0;i<len(A1);i++ {
        sum += float64(A1[i].ts) * A1[i].n
        nsum += A1[i].n
    }
    avg := sum/nsum;
    fmt.Printf("my_time: %f sum %f nsum %f Avg %f  Theta_ga %f \n",my_time,sum,nsum,avg,parameters.Theta_ga)
    if my_time - avg < parameters.Theta_ga {
       fmt.Println("NOT ENOUGH TO DO GA RETURNING ")
       return
    }

    //var child0 Classifier
    //var child1 Classifier
    fmt.Println("looks like we do ga")

    var parent0 Classifier
    var parent1 Classifier
    var child []Classifier

    l := len(A1)
    fmt.Printf("before loop length action set classifiers %d\n",l)
 
    for i=0;i<len(A1);i++ {
fmt.Printf("action set classifier %d \n %+v \n",i,A1[i])
          A1[i].ts = my_time
          parent0 = select_offspring()
          parent1 = select_offspring()
          
fmt.Println("building children") 
          child = append(child,Classifier{})
          child = append(child,Classifier{})
          //child[1] = Classifier{}
          my_copy(&parent0,&child[0])
          my_copy(&parent1,&child[1])

          child[0].n = 1
          child[1].n = 1
          child[0].Exp = 0
          child[1].Exp = 0
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
fmt.Println("DOING CHILD STUFF")
          for i=0;i<len(child);i++ {
              child[i].Condition = apply_mutation(child[i].Condition)
              if parameters.Do_ga_subsumption {
                  if does_subsume(parent0,child[i]) {
                     parent0.n++
                  } else if does_subsume(parent1,child[i]) {
                     parent1.n++
                  } else { 
                     insert_in_population(child[i])
                  }
               } else {
                   insert_in_population(child[i])
               } //end of if on do_ga

               delete_from_population()
           } //end of loop on child 
    } //end of loop on classifiers in A1
fmt.Println("leaving run ga after doing ga")
} //end of run_ga

func select_offspring() Classifier {

    fitness_sum := 0.0
    i := 0
    for i=0;i<len(A1);i++ {
         fitness_sum += A1[i].F
    }
    choice_point := rand.Float64() * fitness_sum
    fitness_sum = 0.0
    pindex := 0
    for i=0;i<len(A1);i++ {
         fitness_sum += A1[i].F
         if fitness_sum > choice_point {
           pindex = i 
           break
        } //end of if on fitness
    } //end of loop on i
           
    return A1[pindex]

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
fmt.Printf("in insert_in_population cl: %+v\n",cl)
    i := 0
    for i=0;i<len(p);i++ {
         if p[i].Condition == cl.Condition &&
            p[i].Action == cl.Action {
                p[i].n++
                return
         }
    }
    p = append(p,cl)
} //end of insert_in_population

func delete_from_population() {
    sum := 0.0
    fsum := 0.0
    i := 0
    for i=0;i<len(p);i++ {
        sum += p[i].n
        fsum += p[i].F
    }

    if sum < float64(parameters.N) {
        return
    }

    av_fitness_in_population := sum/fsum 
    vote_sum := 0.0
    for i=0;i<len(p);i++ {
       vote_sum += deletion_vote(p[i],av_fitness_in_population)
    }

    choice_point := rand.Float64() * vote_sum
    vote_sum = 0.0
    for i=0;i<len(p);i++ {
       vote_sum += deletion_vote(p[i],av_fitness_in_population)
       if vote_sum > choice_point {
           if p[i].n > 1 {
              p[i].n--
           } else {
             remove_classifier_from_population(i)
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

func do_A1_subsumption() {
    var cl *Classifier

//    cl = make_classifier()
    c_num_dont_cares := 0
    cl_num_dont_cares := 0
    i := 0
    for i=0;i<len(A1);i++ {
       c_num_dont_cares = get_num_dont_care(A1[i].Condition)
       cl_num_dont_cares = get_num_dont_care(cl.Condition)
 
       if could_subsume(A1[i]) {
          if cl == nil || c_num_dont_cares > cl_num_dont_cares  {
              //from to
              //cl is already a pointer
              my_copy(cl,&A1[i])
          }

          if c_num_dont_cares == cl_num_dont_cares && rand.Float64() < 0.5 {
              //from to
              my_copy(cl,&A1[i])
          }
       }
    } //end of loop on c

    if cl != nil {
        i := 0
        for i=0;i<len(A1);i++ {
          if is_more_general(cl,A1[i]) {
                cl.n = cl.n + A1[i].n
                remove_classifier(A1,i)[_from_A1(A1[i])
                remove_classifier_from_population(A1[i])
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

func is_more_general(clgen *Classifier,clspec Classifier) bool {
    
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

func does_subsume(clsub Classifier,cltos Classifier) bool {
    if clsub.Action == cltos.Action {
       if could_subsume(clsub) {
           if is_more_general(&clsub,cltos) {
              return true
           }
       }
    }
    return false
}

func my_copy(from *Classifier, to *Classifier) {
    *to = *from
}

func remove_classifier_from_population(index int) {
    fmt.Printf("index %d removing %v\n",index,p[i])
    fmt.Printf("before length of p %4d \n",len(p))
    p[len(p)-1],p[index] = p[index],p[len(p) -1]
    p = p[:len(p) -1]
    fmt.Printf("after length of p %4d \n",len(p))
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
