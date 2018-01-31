package main
import ( "fmt")

type Parameters struct {
    N   int
    Beta float64
    Alpha float64
    Gamma float64
    Epsilon_zero float64
    V  float64
    Exp float64
    Theta_ga float64
    Chi   float64
    Mu    float64
    Theta_del int
    Sigma float64
    Theta_sub float64
    Prob_sharp float64
    Prediction_initial float64
    Epsilon_initial float64
    Fitness_initial float64
    Prob_explor float64
    Theta_mna int
    Do_ga_subsumption bool
    Do_action_set_subsumption bool
}

var possible_actions map[int]int

var parameters Parameters

func get_parameters() {

        parameters.N = 100
        parameters.Prob_sharp = 0.20
        parameters.Prediction_initial = 10.0 //make this relative to reward
        parameters.Gamma = 0.71
        parameters.V = 5
        parameters.Theta_ga = 25.0 //adjust
        parameters.Chi = 0.5
        parameters.Mu = 0.025
        parameters.Theta_del = 20
        parameters.Sigma = 0.1
        parameters.Theta_sub = 20
        parameters.Prediction_initial = 0.00001
        parameters.Epsilon_initial = 0.00001
        parameters.Fitness_initial = 0.00001
        parameters.Theta_mna = 5  //equal to the number of actions..could be smaller
        parameters.Prob_explor = 0.5
        parameters.Do_ga_subsumption = true
        parameters.Do_action_set_subsumption = true

        fmt.Printf("PARAMETERS %v ",parameters)

        //this is for actions dont move , up,down,left,right
        //used in testing for cover
        possible_actions = make(map[int]int) 
        i := 0
        for i=0;i<5;i++ {
             possible_actions[i] = 0
        }
} //end of get_parameters


