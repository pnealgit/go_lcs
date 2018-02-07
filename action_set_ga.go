package main

import ( 
    "fmt"
    "os"
    "math/rand"
)
 
func update() {

    fmt.Printf("IN UPDATE/GA\n")
    sum := 0.0
    nsum:= 0.0
    for _,v := range A1 {
        sum += float64(v.Time_stamp) * v.Numerosity
        nsum += v.Numerosity
    }
    avg := sum/nsum;
    if my_time - avg < float64(p.Ga_threshold) {
       return
    }
    var parent0 Classifier
    var parent1 Classifier
    var child0 Classifier
    var child1 Classifier

    for k,v := range A1 {
          v.Time_stamp = my_time
          A1[k] = v
    }
    parent0 = select_parent()
    parent1 = select_parent()

    var new_children map[int]Classifier

    my_copy(&parent0,&child0)
    my_copy(&parent1,&child1)
    child0.Numerosity = 1
    child1.Numerosity = 1
    child0.Experience = 0
    child1.Experience = 0
    var condition0 string
    var condition1 string

    if rand.Float64() < p.Crossover_probability {
        condition0,condition1 = crossover(parent0.Condition,parent1.Condition)
        child0.Condition = condition0
        child1.Condition = condition1
        child0.Average_reward  = 
             (parent0.Average_reward + parent1.Average_reward)/2.0
        child0.Error = (parent0.Error + parent1.Error)/2.0
        child0.Fitness = (parent0.Fitness + parent1.Fitness)/2.0
        //dont know for sure
        child0.Action_set_size = 
              (parent0.Action_set_size + parent1.Action_set_size)/2.0
        child1.Average_reward = child0.Average_reward
        child1.Error = child0.Error
        child1.Fitness = child0.Fitness
        child1.Action_set_size = child0.Action_set_size
     }

      

    child0.Condition = mutate(child0.Condition)
    child1.Condition = mutate(child1.Condition)
    child0.Fitness = child0.Fitness * 0.1
    child1.Fitness = child1.Fitness * 0.1

    new_children[0] = child0
    new_children[1] = child1

    subsume := false
    for k,v := range new_children {
        if p.Do_ga_subsumption {
     
            if does_subsume(parent0,v) {
                parent0.Numerosity++
                insert_into_population(parent0)
                subsume = true
            } 
            if does_subsume(parent1,v) {
               parent1.Numerosity++
               insert_into_population(parent1)
               subsume = true
            }
            if subsume {
               delete(new_children,k)
            } 
        }
    }

   //any children left should be inserted
   for _,v := range new_children {
        insert_into_population(v)
        delete_from_population()
    }

} //end of run_ga

func select_parent()  Classifier {

    total_fitness := 0.0
    for _,v := range A1 {
         total_fitness += v.Fitness
    } 

    selector := rand.Float64() * total_fitness
    total_fitness = 0.0
    var selected_k string
    selected_k = ""
    for k,v := range A1 {
         selector -= v.Fitness
         if selector <= 0.0 {
           selected_k = k
           break
         } //end of if on fitness
    } //end of k,v loop 
    if selected_k == "" {
       fmt.Printf("BAD PARENT SELECTION \n")
       fmt.Printf("Exiting\n")
       os.Exit(7)
    }
    return A1[selected_k]

} //end of select_offspring

func crossover(c0 string ,c1 string)  (string,string){
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

func mutate(condition string) string {
    i := 0
    temp := []byte(condition)

    for i=0;i<len(temp);i++ {
         if rand.Float64() < p.Mutation_probability {
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
