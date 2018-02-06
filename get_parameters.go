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

        //comments on left are Butz and Wilson names

        parameters.Population_size = 400       //N
        parameters.Learning_rate =  0.15 	//beta
        parameters.Accuracy_coefficient = 0.1   //alpha
        parameters.Error_threshold = 0.01       //epsilon_zero .. for rewards between 0 and 1
        parameters.Accuracy_power  = 5.0        //nu
        parameters.Discount_factor = 0.71       //gamma
        parameters.Ga_threshold = 35            //theta_GA
        parameters.Crossover_probability = 0.5  //chi
        parameters.Mutation_probability = 0.03  //mu
        parameters.Deletion_threshold = 20	//theta_del
        parameters.Fitness_threshold = .1       //delta
        parameters.Subsumption_threshold = 20   //theta_sub 
        parameters.Wildcard_probability = .33   //P_#
        parameters.Initial_prediction = 0.00001 //P_i
        parameters.Initial_error = 0.00001      //epsilon_i
        parameters.Initial_fitness = 0.00001    //f_i
        parameters.Exploration_probability = 0.5 //p_exp
        parameters.Minimum_actions = 3           //equal to the number of actions..could be smaller
        parameters.Do_ga_subsumption = false 
        parameters.Do_action_set_subsumption = false 

        fmt.Printf("PARAMETERS %+v ",parameters)

        //this is for actions dont move , up,down,left,right
        //used in testing for cover
        possible_actions = make(map[int]int) 
        i := 0
        for i=0;i<3;i++ {
             possible_actions[i] = 0
        }
} //end of get_parameters


func (p Parameters) String() string {
     //this will be invoked when you do a fmt.Printf("%+v\n",parameters)
     kaboodle := fmt.Sprintf(
              "Population Size: \t%f\n
               Learning_rate  : \t%f\n
               Accuracy_Coefficient: \t%f\n
               Error_threshold: \t%f\n
               Accuracy_power :  \t%f\n
               Discount_factor: \t%f\n
               Ga_threshold   :  \t%f\n",
               p.Population_size,
               p.Learning_rate,
               p.Accuracy_coefficient,
               p.Error_threshold,
               p.Accuracy_power,
               p.Discount_factor,
               p.Ga_threshold
               )
    return kaboodle
}



