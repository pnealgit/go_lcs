package main

import (
	"encoding/json"
	"fmt"
        "bytes"
        "os"
        "math/rand"
        "math"
)

type State_record struct {
	State  []int
        Reward float64
}

type Angle_record struct {
    Angle float64
}

var state_record State_record
 
var angle_record Angle_record

type Classifier struct {
    Condition    string
    Action       int
    p            float64      //predicted payoff
    Epsilon float64
    F            float64      //fitness
    Exp   float64 
    Last_ga_time_stamp   int
    as float64                   //average selection set size
    n   float64                  //numerosity
}

type Population struct {
    Classifiers []Classifier
}

type Match_set struct {
    Classifiers []Classifier
}

type Action_set struct {
    Classifiers []Classifier
}

var p Population
var m Match_set
var action_set Action_set
var action_set_minus Action_set

func do_lcs( message []byte) []byte {
       
	var err error
        var state_record State_record

	jerr := json.Unmarshal(message, &state_record)
	if jerr != nil {
		fmt.Println("error on state unmarshal")
		panic(fmt.Sprintf("%s", "ARRRGGGH"))
	} //end of if on jerr
        state_string := ""
        state_string = convert_input(state_record)
        if len(state_string) != 36 {
            fmt.Printf("length is : %4d\n",len(state_string))
            fmt.Printf("%v\n",state_record)
            fmt.Printf("%s\n",state_string)
            os.Exit(2)
        }
    
        generate_match_set(state_string)
        dump_match_set(state_string)

        action := 0
        action = generate_prediction_array()
        fmt.Printf("Selected Action %d \n",action)

        generate_action_set(action)

        update_action_set(state_record.Reward)

        if len(p.Classifiers) >= parameters.N {
             dump_population()
             os.Exit(3)
        }
        

        angle_record.Angle = 1.0
	message,err  = json.Marshal(angle_record)
        if err != nil {
                fmt.Println("bad angles Marshal")
        }
	return message
} //end of do_update

func think(state_record State_record) float64 {
	return 1.0	
} //end of think

func convert_input(state_record State_record) string {
    var s bytes.Buffer
    
    si := ""
    i := 0
    si = fmt.Sprintf("%06b",state_record.State[0]%20)
    s.WriteString(si)

    si = fmt.Sprintf("%06b",state_record.State[1]%20)
    s.WriteString(si)

    for i=2;i<len(state_record.State);i++ {
        si = fmt.Sprintf("%04b",state_record.State[i])
        s.WriteString(si)
    }
    return s.String()
}

func dump_population() {
     i := 0
     kntr := make(map[string]int)

     for i=0;i<len(p.Classifiers);i++ {
       kntr[p.Classifiers[i].Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("%d %s\n",v,k)
     }

} //end of dump

func dump_match_set(state_string string) {
     fmt.Println("DUMP MATCH SET \n")
     fmt.Printf( "%s\n",state_string)
     i := 0
     kntr := make(map[string]int)
     kntr2 := make(map[int]int)

     for i=0;i<len(m.Classifiers);i++ {
       kntr[m.Classifiers[i].Condition]++
     }
     for k,v := range kntr {
        fmt.Printf("Match set %d %s\n",v,k)
     }

     for i=0;i<len(m.Classifiers);i++ {
       kntr2[m.Classifiers[i].Action]++
     }
     fmt.Printf("Action Counts\n")

     for k,v := range kntr2 {
        fmt.Printf("ACTIONS %d %d\n",v,k)
     }

} //end of dump_match_set


func generate_prediction_array() int {
    fmt.Println("generating prediction_array")
    pa := make(map[int]float64) 
    fsa := make(map[int]float64) 
    i := 0
    for i=0;i<len(m.Classifiers);i++ {
        pa[m.Classifiers[i].Action] += m.Classifiers[i].p * m.Classifiers[i].F
        fsa[m.Classifiers[i].Action] += m.Classifiers[i].F
    } 
    for k,_ := range possible_actions {
        if fsa[k] > 0.0 {
            pa[k] = pa[k]/fsa[k]
        }
    }
    fmt.Printf("PA %+v\n",pa)

    action := 0               //do nothing for the moment
    action = select_action(pa)
    return action
}
func select_action(pa map[int]float64) int {
    action := 0
    if rand.Float64() < parameters.Prob_explor {
        //explore
        fmt.Println("explore")
        action = rand.Intn(len(possible_actions))
    } else {
        //exploit
        fmt.Println("exploit")
        //get max
        max := -999.9
        for k,v := range pa {
            fmt.Printf("k %3d v %f \n",k,v)
            if v > max {
                action = k
                max = v
            }
        } //end of get max
    } //end of exploit
    return action
} //end of select_action


func generate_action_set(action int) {
   action_set.Classifiers = nil
   i := 0
   for i=0;i<len(m.Classifiers);i++ {
      if m.Classifiers[i].Action == action {
         action_set.Classifiers = append(action_set.Classifiers,m.Classifiers[i])
      }
   }
} //end of generate action set

func update_action_set(reward float64) {
    //reward is 'P' 
    fmt.Printf("UPDATE ACTION SET REWARD IS : %f \n",reward)
    i := 0
    for i=0;i<len(action_set.Classifiers);i++ {
        action_set.Classifiers[i].Exp++

        //prediction
        delta := reward - action_set.Classifiers[i].p 
        if action_set.Classifiers[i].Exp < 1.0/parameters.Beta {
            action_set.Classifiers[i].p += delta/action_set.Classifiers[i].Exp
        } else {
            action_set.Classifiers[i].p += parameters.Beta * delta
        }

        //prediction_error
        ad := math.Abs(delta)
        if action_set.Classifiers[i].Exp < 1.0/parameters.Beta {
            action_set.Classifiers[i].Epsilon += 
                 (ad - action_set.Classifiers[i].Epsilon)/action_set.Classifiers[i].Exp
       } else {
            action_set.Classifiers[i].Epsilon += parameters.Beta * (ad - action_set.Classifiers[i].Epsilon)
       }

       //action set size estimate
       //'c.n' is numerosity
       j := 0
       sum := 0.0
       for j=0;j<len(action_set.Classifiers);j++ {
            sum += action_set.Classifiers[j].as - action_set.Classifiers[i].n
       } 

        if action_set.Classifiers[i].Exp < 1.0/parameters.Beta {
            action_set.Classifiers[i].as += sum /action_set.Classifiers[i].Exp
        } else {
            action_set.Classifiers[i].as += parameters.Beta * sum
        }
    } //end of loop on classifiers in action set

    update_fitness()
}

func update_fitness() {
    accuracy_sum := 0.0
    var k []float64

    k = make([]float64,len(action_set.Classifiers))

    i := 0
    for i=0;i<len(action_set.Classifiers);i++ {
        if action_set.Classifiers[i].Epsilon < parameters.Epsilon_zero {
            k[i] = 1.0
        } else {
            tspread := (action_set.Classifiers[i].Epsilon/parameters.Epsilon_zero)
            spread := math.Pow(tspread,parameters.V)
            k[i] = parameters.Alpha * spread
        } 
        accuracy_sum += k[i] * action_set.Classifiers[i].n
    }
    for i=0;i<len(action_set.Classifiers);i++ {
        quant := (k[i] * action_set.Classifiers[i].n/accuracy_sum) - action_set.Classifiers[i].F
        action_set.Classifiers[i].F = action_set.Classifiers[i].F + parameters.Beta * quant
    }
} //end of update fitness
     
