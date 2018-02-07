package main
import(
    "math/rand"
)
func do_action_set_subsumption() {
    var cl Classifier

    c_num_dont_cares := 0
    cl_num_dont_cares := 0
    
    for _,v := range A1 {
 
       if could_subsume(v) {
           c_num_dont_cares = get_num_dont_care(v.Condition)
           cl_num_dont_cares = get_num_dont_care(cl.Condition)

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
                cl.Numerosity = cl.Numerosity + v.Numerosity
                delete(A1,k)
                delete(population,k)
          }
        }
    } //end of if on nil
} //end of do_A1_subsumption

func could_subsume(cl Classifier) bool {
    if cl.Experience > float64(p.Subsumption_threshold) {
        if cl.Error < p.Error_threshold {
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

