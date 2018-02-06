package main
import (
    "fmt"
    "math"
    "bytes"
)

func convert_state_input(state_record State_record) {

    reward = 0.0
    var s bytes.Buffer
    si := ""
    //last angle
    si = fmt.Sprintf("%03b",state_record.Last_angle)
    //s.WriteString(si)

    xchunk := int((state_record.X/canvas_width)*100.0)
    si = fmt.Sprintf("%08b",xchunk)
    //s.WriteString(si)

    ychunk := int((state_record.Y/canvas_height)*100.0)
    si = fmt.Sprintf("%08b",ychunk)
    //s.WriteString(si)

    for _,v := range state_record.Sensors {
        //near or far
        x_dist :=  math.Abs(v.Xpos - state_record.X)
        y_dist :=  math.Abs(v.Ypos - state_record.Y)
        nf_x := int((x_dist/canvas_width)*100.0)
        nf_y := int((y_dist/canvas_height)*100.0)

        i_nf_x := 0
        if nf_x < 25 {
            i_nf_x = 1
        }
        if nf_x >= 25 && nf_x < 50 {
            i_nf_x = 2
        }
        if nf_x >= 50 && nf_x < 75 {
            i_nf_x = 4
        }
        if nf_x >=  75 {
            i_nf_x = 8
        }

        i_nf_y := 0
        if nf_y < 25 {
            i_nf_y = 1
        }
        if nf_y >= 25 && nf_y < 50 {
            i_nf_y = 2
        }
        if nf_y >= 50 && nf_y < 75 {
            i_nf_y = 4
        }
        if nf_y >=  75 {
            i_nf_y = 8
        }

        si = fmt.Sprintf("%04b",i_nf_x)
        s.WriteString(si)

        si = fmt.Sprintf("%04b",i_nf_y)
        s.WriteString(si)

        istatus := int(v.Status)
        if istatus == 5 {
           rel_distance := math.Hypot(v.Xpos,v.Ypos)/canvas_max_distance
           reward+= 20.0 * (1.0 - rel_distance)
        }
        if istatus < 5 {
           reward+= 0.01
        }

        if istatus == 5 {
          istatus = 1
        } else {
          istatus = 0
        }
         
        si = fmt.Sprintf("%02b",istatus)
        s.WriteString(si)

         
    }
    
    state_string = s.String()
}

