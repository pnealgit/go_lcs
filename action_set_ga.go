package main

import ( 
    "fmt"
    "math/rand"
    "os"
)
 
func run_ga() {
    sum := 0.0
    nsum:= 0.0
    for _,v := range A1 {
        sum += float64(v.ts) * v.n
        nsum += v.n
    }
    avg := sum/nsum;
    if my_time - avg < parameters.Theta_ga || len(A1) < 2 {
       return
    }

    var parent0 Classifier
    var parent1 Classifier
    var child  []Classifier

    for k,v := range A1 {
          v.ts = my_time
          A1[k] = v
    }
    p0x := select_offspring()
    p1x := select_offspring()

    parent0 = A1[p0x]
    parent1 = A1[p1x]

    child = append(child,Classifier{})
    child = append(child,Classifier{})
    my_copy(&parent0,&child[0])
    my_copy(&parent1,&child[1])

    child[0].n = 1
    child[1].n = 1
    child[0].Exp = 0
    child[1].Exp = 0
    fmt.Printf("c0 %+v\n",child[0]);
    fmt.Printf("p0 %+v\n",parent0);
    fmt.Printf("c1 %+v\n",child[1]);
    fmt.Printf("p1 %+v\n",parent1);

    if rand.Float64() < parameters.Chi {
          apply_crossover(&child[0],&child[1])
     
          child[0].p = (parent0.p + parent1.p)/2
          child[0].Epsilon = (parent0.Epsilon + parent1.Epsilon)/2
          child[0].F = (parent0.F + parent1.F)/2
          child[1].p = child[0].p
          child[1].Epsilon = child[0].Epsilon
          child[1].F = child[0].F
          fmt.Printf("after crossover\n")
          fmt.Printf("c0 %+v\n",child[0]);
          fmt.Printf("p0 %+v\n",parent0);
          fmt.Printf("c1 %+v\n",child[1]);
          fmt.Printf("p1 %+v\n",parent1);
    }
    child[0].F = child[0].F * 0.1
    child[1].F = child[1].F * 0.1
     
    i := 0 
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

          //this means clean up population
          delete_from_population()
      } //end of loop on child 
} //end of run_ga

func select_offspring()  string {

    fitness_sum := 0.0
    for _,v := range A1 {
         fitness_sum += v.F
    } 
    if fitness_sum <= 0.0 {
        fmt.Printf("FITNESS SUM %f\n",fitness_sum)
        fmt.Printf("Exiting from select_offspring\n")
        os.Exit(3)
    }

    choice_point := rand.Float64() * fitness_sum
    fitness_sum = 0.0
    pindex := ""
    for k,v := range A1 {
         fitness_sum += v.F
         if fitness_sum > choice_point {
           pindex = k
           break
         } //end of if on fitness
    } //end of k,v loop 
        
    return pindex  
} //end of select_offspring

func apply_crossover(c0,c1 *Classifier)  {

    cut_point := rand.Intn(len(c0.Condition))
    var tmp1a,tmp1b string
    var tmp2a,tmp2b string
    
    tmp1a = c0.Condition[cut_point:]
    tmp1b = c0.Condition[:cut_point]
    tmp2a = c1.Condition[cut_point:]
    tmp2b = c1.Condition[:cut_point]
    c0.Condition = tmp1a + tmp2b
    c1.Condition = tmp2a + tmp1b
}

func apply_mutation(condition string) string {
    i := 0
    temp := []byte(condition)

    for i=0;i<len(temp);i++ {
         if rand.Float64() < parameters.Mu {
             if temp[i] == '#' {
                temp[i] = state_string[i]
             } else {
                temp[i] = '#'
             }
         }
    } //end of loop
    condition = string(temp)
    return condition 
} //end of apply mutation

func remove_classifier_from_A1(cl Classifier) {
     fmt.Println("NOT FINISHED !!!!!! removing from action set ")
}
