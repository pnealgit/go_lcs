package main

import (
	"encoding/json"
        "math/rand"
        "math"
	"fmt"
        "time"
)


type Sensor_record struct {
	Id     int
	Type   string
	Data   string
}

var sensor_record Sensor_record

var action_set = []int{}
var previous_action_set = []int{}
var bigP  float64
var maxPA float64

var previous_reward float64
var sensor_data string
var previous_sensor_data string
 
 
func do_generates(message []byte) []byte {
//this will receive a sensor record and return an
//angle record


	type Angle_record struct {
		Id    int
		Angle float64
	}

        var angle_record Angle_record;
	var err error

	jerr := json.Unmarshal(message, &sensor_record)
	if jerr != nil {
		fmt.Println("error on update unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr

        var match_set  []int

        match_set = generate_match_set();
     
        var pa map[string]float64
        pa = generate_prediction_array(match_set);

        //var best_action string
        //var maxPA float64

        best_action,maxPA := select_action(pa);
fmt.Println("BEST ACTION,maxPA: ",best_action,maxPA)

        action_set = generate_action_set(match_set,best_action)
fmt.Println("ACTION SET ",action_set)

        angle_record.Angle = action_angles[best_action]
	angle_record.Id = sensor_record.Id

	message, err = json.Marshal(angle_record)
	if err != nil {
		fmt.Println("bad angle Marshal")
	}
	return message
} //end of do_generate

func do_updates(message []byte) {
//this will receive a reward record and return nothing;

type Reward_record struct {
        Id     int
        Type   string 
        Data   float64
}

	var reward_record Reward_record
	//var err error

	jerr := json.Unmarshal(message, &reward_record)
	if jerr != nil {
		fmt.Println("error on reward unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr
        reward :=  reward_record.Data
if reward > 0.0 {
    fmt.Println("\nYAAAYYYYYYY REWARD: ",reward,"\n");
}

    if len(previous_action_set) > 0 {
         bigP = previous_reward + GAMMA * maxPA
         update_action_set()
         run_ga()
    } else {
         previous_action_set = action_set
         previous_reward = reward
         previous_sensor_data = sensor_data
    }

    
} //end of do_update

func generate_match_set() []int {

    var match_set = []int{}
    covered_actions := map[string]int{}
    for _,ac := range actions {
        covered_actions[ac] = 0
    }

    for ic,cl := range cset {
        //use click count for GA later on
        cset[ic].Ga_clicks++
        if does_match(cl.Condition,sensor_record.Data) {
            match_set = append(match_set,ic)
            covered_actions[cl.Action]+= 1
        } //end of if  on match
    } //end of loop on cl set
   
    var cl1 int
    for ac,knt := range covered_actions {
       if knt < 1 {
           cl1 = cover(sensor_record.Data,ac)
           match_set = append(match_set,cl1)
       }
    } //end of loop on covered_actions

    return match_set
 
} //end of get_matched_set

func cover(sensor_data string,action string) int {
    cond := ""
    for spot:=0;spot<len(sensor_data);spot++ {
        if rand.Float64() < P_SHARP {
           cond+= "#"
        } else {
           cond+= string(sensor_data[spot])
        }
    } //end of loop on spot

    len_c := len(cset)
    cl := Classifier{Id:len_c}
    cl.Condition = cond 
    cl.Action = action
    cl.Prediction = P_INITIAL
    cl.Epsilon = EPSILON_0
    cl.Fitness = FITNESS_INITIAL
    cl.Numerosity = 1
    cl.Experience = 1.0
    //cl.Ga_ts = int32(time.Now().Unix()) - start_time
    cl.Ga_clicks = 0;
    cset = append(cset,cl);
    return len_c

} //end of generate_covering classifier


func does_match(cls string, sensed string) bool {
    for i := 0;i<len(sensed);i++ {
        if cls[i] != '#' && cls[i] != sensed[i] {
            return false
        }
    }
    return true
}//end of does match
   
func generate_prediction_array(m_set []int) map[string]float64 {
    prediction_array := make(map[string]float64)
    fitness_sum := make(map[string]float64)
    for _,ac := range actions {
        prediction_array[ac] = 0.00
        fitness_sum[ac] = 0.0
    }
    
    var cl Classifier
 
    for _,cix := range m_set {
       cl = cset[cix]
       prediction_array[cl.Action] += cl.Prediction * cl.Fitness
       fitness_sum[cl.Action] += cl.Fitness
    }

    for _,ac := range actions {
        if fitness_sum[ac] > 0.0 {
           prediction_array[ac] = prediction_array[ac]/fitness_sum[ac]
        }
    } //end of loop on actions
    return prediction_array
}
func select_action(pa map[string]float64) (string,float64) {
   var best_action string
   maxPA := -9999.0

   if rand.Float64() < P_EXP {
       //exploitation
       for ac,val := range pa {
          if val > maxPA {
             best_action = ac
             maxPA = val
          }
       } //end of loop through pa

       //end of exploitation
   } else {
       //exploration
       mlen := len(actions)
       ik := getRandomInt(0,mlen)
       best_action = actions[ik];
       maxPA = pa[best_action]
   }
   return best_action,maxPA

} //end of select action

func generate_action_set(match_set []int,best_action string) []int {
    var action_set []int
    for _,cix := range match_set {
        if cset[cix].Action == best_action {
           action_set = append(action_set,cix)
        }
    } //end of loop on match set
    return action_set
}//end of action set

func update_action_set() {
fmt.Println("previous action set bigP is ",previous_action_set,bigP)
  
    //get numerosity
    sum_numerosity := 0.0
    for _,cix :=range previous_action_set {
        sum_numerosity += cset[cix].Numerosity
    }

    for _,cix :=range previous_action_set {

        cset[cix].Experience++

        //prediction
        pp := bigP - cset[cix].Prediction
        if cset[cix].Experience < (1.0 /BETA) {
            cset[cix].Prediction += pp/cset[cix].Experience
        } else {
            cset[cix].Prediction += BETA * pp
        }
 
        //prediction error (epsilon)
        //might have to flip prediction and prediction error sequence

        pp = bigP - cset[cix].Prediction  //just updated it
        if cset[cix].Experience < 1.0 /BETA {
           cset[cix].Epsilon+= 
               (pp-cset[cix].Epsilon)/cset[cix].Experience
        } else {
           cset[cix].Epsilon+= BETA * (pp-cset[cix].Epsilon)
        }

        //action set size

        if cset[cix].Experience < 1.0 /BETA {
            cset[cix].As += (sum_numerosity-cset[cix].As)/cset[cix].Experience
        } else {
            cset[cix].As += BETA * (sum_numerosity-cset[cix].As)
        }
    } //end of loop on action set
 
} //end of update action set

func update_fitness() {
fmt.Println("update_fitness");
    accuracy_sum := 0.0
    kappa := []float64{}

    //action set is an array so entries are in order and stay that way
    for ia,cix := range previous_action_set {
       if cset[cix].Epsilon < EPSILON_0 {
           kappa[ia] = 1.0
       } else {
           q := cset[cix].Epsilon/EPSILON_0
           kpow := math.Pow(q,GAMMA)
           kappa[ia] = ALPHA * kpow
       }
       accuracy_sum+=kappa[ia]*cset[cix].Numerosity
    } //end of loop on previous action set

    for ia,cix := range previous_action_set {
         pq := kappa[ia] * cset[cix].Numerosity/accuracy_sum
         cset[cix].Fitness+= BETA*(pq-cset[cix].Fitness)
    }
} //end of update fitness

func run_ga() {
    fmt.Println("run ga")
    sum_clicks := 0.0
    sum_numerosity = 0.0
    for _,cix := range previous_action_set {
        sum_clicks+= cset[cix].Ga_clicks*cset[cix].Numerosity
        sum_numerosity += cset[cix].Numerosity
    }

    avg_clicks := math.Floor(sum_clicks/sum_numerosity)
    fmt.Println("AVG CLICKS,THETA_CLICKS ",avg_clicks,THETA_CLICKS);

    if avg_clicks < THETA_CLICKS {
        return
    }

    fmt.Println("DOING GA");
    for _,cix := range previous_action_set {
        cset[cix].Ga_clicks = 0.0
    }

    pix1 := select_offspring()
    pix2 := select_offspring()
    p1 = cset[pix1]
    p2 = cset[pix2]
  
    child1 := p1
    child2 := p2

    child1.Numerosity = 1.0
    child1.Experience = 0.0 
    child2.Numerosity = 1.0
    child2.Experience = 0.0 
    
    if rand.Float64() < CHI {
        cx1,cx2 := apply_crossover(child1.Condition,child2.Condition)
        child1.Condition = cx1
        child2.Condition = cx2
        child1.Prediction = (p1.Prediction + p2.Prediction)/2.0 
        child1.Prediction_error = 0.25 * (p1.Prediction_error + p2.Prediction_error)/2.0
        child1.fitness = 0.1 * (p1.Fitness + p2.Fitness)/2.0
        child2.Prediction = child1.Prediction
        child2.Prediction_error = child1.Prediction_error
        child2.Fitness = child1.Fitness

        //subsumption goes here --
     }
     child1.Condition = apply_condition_mutation(child1.Condition)
     child2.Condition = apply_condition_mutation(child2.Condition)
     child1.Action = apply_action_mutation(child1.Action)
     child2.Action = apply_action_mutation(child2.Action)
     insert_in_population(child1)
     insert_in_population(child2)
} //end of run_ga

func select_offspring() int32 {
     fitness_sum := 0.0
     for _,cix := range previous_action_set {
        fitness_sum += cset[cix].fitness
     }
     choice_point = rand.Float64() * fitness_sum

     fitness_sum = 0.0
     for _,cix := range previous_action_set {
        fitness_sum += cset[cix].fitness
        if fitness_sum > choice_point {
           return cix
        } //end of if on choice_point
     } //end of loop

} //end of select offsprint

func apply_crossover(c1s string,c2s string) string string {
    x := getRandomInt(0,len(c1s)+1)
    y := getRandomInt(0,len(c1s)+1)
    if x > y {
       temp := x
       x = y
       y = temp
    }
    i := 0
    for {
       if x <= i && i < y {
          ctemp := c1s[i]
          c1s[i] = c2s[i]
          c2s[i] = ctemp
       }
       if i >= y {
          break
       }
    } //end of for loop that looks like a do while
    return c1s,c2s
} //end of apply_crossover

func apply_condition_mutation( c string) string {
    fmt.Println("APPLY CONDITION  MUTATION")
    for i:=0;i<len(c); i++ {
        if rand.Float64() < MU {
            if c[i] == "#" {
                c[i] = previous_sensor_data[i]
            } else {
                c[i] = "#"
            } //end of if on hash
        } //end of if on mu
    } //end of loop on len
    return c
} //end of apply_condition_mutation

func apply_action_mutation(ac string) string {
    if randFloat64 >= MU {
        return ac
    }
    rx := 0
    tac := ""
    for {
        rx = getRandomInt(0,len(actions))
        if actions[rx] != ac {
            return actions[rx]
        } //end of if on not match
    } //end of fake do while
} //end of apply_action_mutation

func insert_in_population(cl Classifier) {
    for cix,cl := range cset {
        if cl.Condition == cond && cl.Action == ac {
            cset[cix].Numerosity++
            return
        }
    } //end of loop on cset

    cset = append(cset,cl)
}//end of insert in population
