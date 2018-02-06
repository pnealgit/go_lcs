package main

import ( 
    "fmt"
    "math/rand"
    "os"
)
 
func run_ga() {
    fmt.Printf("IN GA\n")
    sum := 0.0
    nsum:= 0.0
    for _,v := range A1 {
        sum += float64(v.ts) * v.n
        nsum += v.n
    }
    avg := sum/nsum;
    if my_time - avg < parameters.Theta_ga {
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

    fmt.Printf("LEN A1 %d ACTION IN GA: %d\n",len(A1),parent0.Action)

    child = append(child,Classifier{})
    child = append(child,Classifier{})
    my_copy(&parent0,&child[0])
    my_copy(&parent1,&child[1])

    child[0].n = 1
    child[1].n = 1
    child[0].Exp = 0
    child[1].Exp = 0
    
    if rand.Float64() < parameters.Chi {
          child[0].Condition,child[1].Condition = apply_crossover(child[0].Condition,child[1].Condition)
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
          //delete_from_population()
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

func apply_crossover(c0 string ,c1 string)  (string,string){
    //two point crossover
    //swap values betweens 2 points

    if c0 == c1 {
        fmt.Printf("in crossover arguments are the same\n")
    }

    b0 := []byte(c0)
    b1 := []byte(c1)

    x := rand.Intn(len(b0) )
    y := rand.Intn(len(b0) )

    if x > y {
       x,y = y,x
    }
    
    i := 0
    //var tmp byte

    for i=0;i<len(b0);i++ {
        if i >= x && i <= y {
             b0[i],b1[i] = b1[i],b0[i]
        } 
    }
    if string(b0) == string(b1) {
       if b0[x] == dont_care {
           b0[x] = state_string_minus_one[x]
       } else {
           b0[x] = dont_care
       }
       if b1[y] == dont_care {
           b1[y] = state_string_minus_one[y]
       } else {
           b1[x] = dont_care
       }
    }
 
    return string(b0),string(b1)
}

func apply_mutation(condition string) string {
    i := 0
    temp := []byte(condition)

    for i=0;i<len(temp);i++ {
         if rand.Float64() < parameters.Mu {
             if temp[i] == '#' {
                temp[i] = state_string_minus_one[i]
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
