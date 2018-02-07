package main
import ( "fmt")

type Parameters struct {
    Population_size   int
    Learning_rate     float64
    Accuracy_coefficient float64
    Error_threshold  float64
    Accuracy_power  float64
    Discount_factor float64
    Ga_threshold    int
    Crossover_probability float64
    Mutation_probability float64
    Deletion_threshold  int
    Fitness_threshold    float64
    Subsumption_threshold  int
    Wildcard_probability  float64
    Initial_prediction  float64
    Initial_error       float64
    Initial_fitness     float64
    Exploration_probability  float64
    Minimum_actions     int
    Do_ga_subsumption   bool
    Do_action_set_subsumption bool
}

var possible_actions map[int]int

var p Parameters

func get_parameters() {

        //comments on left are Butz and Wilson names

        p.Population_size = 400       //N
        p.Learning_rate =  0.15 	//beta
        p.Accuracy_coefficient = 0.1   //alpha
        p.Error_threshold = 0.01       //epsilon_zero .. for rewards between 0 and 1
        p.Accuracy_power  = 5.0        //nu
        p.Discount_factor = 0.71       //gamma
        p.Ga_threshold = 35            //theta_GA
        p.Crossover_probability = 0.5  //chi
        p.Mutation_probability = 0.03  //mu
        p.Deletion_threshold = 20	//theta_del
        p.Fitness_threshold = .1       //delta
        p.Subsumption_threshold = 20   //theta_sub 
        p.Wildcard_probability = .33   //P_#
        p.Initial_prediction = 0.00001 //P_i
        p.Initial_error = 0.00001      //epsilon_i
        p.Initial_fitness = 0.00001    //f_i
        p.Exploration_probability = 0.5 //p_exp
        p.Minimum_actions = 3           //equal to the number of actions..could be smaller
        p.Do_ga_subsumption = false 
        p.Do_action_set_subsumption = false 

        dump_parameters()

        //this is for actions dont move , up,down,left,right
        //used in testing for cover
        possible_actions = make(map[int]int) 
        i := 0
        for i=0;i<3;i++ {
             possible_actions[i] = 0
        }
} //end of get_parameters


func dump_parameters() {
      fmt.Printf("Population Size: \t%f\n", p.Population_size)
      fmt.Printf("Learning_rate  : \t%f\n", p.Learning_rate)
      fmt.Printf("Accuracy_Coefficient: \t%f\n", p.Accuracy_coefficient)
      fmt.Printf("Error_threshold: \t%f\n", p.Error_threshold)
      fmt.Printf("Accuracy_power :  \t%f\n", p.Accuracy_power)
      fmt.Printf("Discount_factor: \t%f\n", p.Discount_factor)
      fmt.Printf("Ga_threshold   :  \t%f\n", p.Ga_threshold)
      fmt.Printf("Crossover_probability : \t%f\n",p.Crossover_probability)
      fmt.Printf("Mutation_probability : \t%f\n",p.Mutation_probability)
}



