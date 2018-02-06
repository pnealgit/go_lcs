
Rover.prototype.make_sensors = function() {
    for (var ns = 0; ns < this.num_antennae; ns++) {
        var s = new Sensor();
        this.sensors.push(s);
    }
} //end of make_sensors

function Sensor() {
      this.xpos = 0;
      this.ypos = 0;
      this.angle = 0.0;
      this.status = 0;
} //end of sensor 

Rover.prototype.set_sensor_positions = function() {
    //right now only 3 so minus delta_radians 0 delta radians and + delta_radians
    var tangle = this.angle- this.delta_radians;
    var delta_x = 0;
    var delta_y = 0;
    var velocity = 2.0;

    for (var knt=0;knt<this.num_antennae;knt++ ){
        this.sensors[knt].status = 0;
        this.sensors[knt].xpos = this.x
        this.sensors[knt].ypos = this.y
        this.sensors[knt].angle = tangle

        if (this.sensors[knt].status <= 0 ) {
            //do nothing
        } else {
          console.log('something fishy')
        }
      
        while (this.sensors[knt].status <= 0) {
            delta_x = velocity * Math.cos(tangle)
            delta_y = velocity * Math.sin(tangle)
            this.sensors[knt].xpos += delta_x
            this.sensors[knt].ypos += delta_y
            this.sensors[knt].status = check_borders(this.sensors[knt].xpos,this.sensors[knt].ypos,1)
            if (this.sensors[knt].status > 0.0) {
               break;
            }
            this.sensors[knt].status = check_food(this.sensors[knt].xpos,this.sensors[knt].ypos,1)
            if (this.sensors[knt].status > 0.0) {
               break;
            }
        } //end of while
        tangle += this.delta_radians;
      } //end of loop on knt
} 

Rover.prototype.get_sensor_data = function() {
    var status = 0;
    var border_status = -9;
    //food 
    for (var ss=0;ss<this.num_sensors;ss++) {
        border_status = -9;
        this.sensors[ss].status = 0;
        s = this.sensors[ss];
        status = 0;

        if (s.xpos < 0 ) {
           this.sensors[ss].xpos = 0;
           s.xpos = 0;
        }

        if (s.xpos > width ) {
           this.sensors[ss].xpos = width;
           s.xpos = width;
        }

        if (s.ypos < 0 ) {
           this.sensors[ss].ypos = 0;
            s.ypos = 0;
        }

        if (s.ypos > height ) {
           this.sensors[ss].ypos = height;
           s.ypos = height;
        }

        border_status = check_borders(s.xpos,s.ypos,1); //1 is fake radius
        if (border_status > 0 && border_status < 5) {
            this.sensors[ss].status = border_status; //status is wall number
            continue;
        }

        status = check_food(s.xpos,s.ypos,1);
        if (status == 5 ) {
console.log("sensor ",ss, "found food status : ",status)
            this.sensors[ss].status = status; //greater than wall index
        }
    } //end of loop on sensors
}
//end of get_sensor_data function

function check_food(xp,yp,radius) {
        var status = 0;
        for (var i = 0; i < num_foods; i++) {
            f = foods[i];
            test = f.r;
            dist = Math.hypot((f.x - xp), (f.y - yp));
            if (dist <= test) {
                status = 5;
                break;
            } //end if food if
        } //end of loop on food
        return status;
}

function check_other_rovers(id,xp,yp) {
        var status = 0;
        for (var i = 0; i < num_rovers; i++) {
           if(i != id) { 
            rvr = rovers[i];
            test = rvr.r;
            dist = Math.hypot((rvr.x - xp), (rvr.y - yp));
            if (dist <= test) {
                status = 6;
                return status;
            } //end of if on distance
           } //end of if on not same rover
        } //end of loop on rovers
        return status;
}

function check_borders(xp,yp,rad) {
    //top
    if ((yp - rad) < 2) {
       return 1;
    }
    //bottom
    if((yp+rad) > height -2) {
       return 2;
    }
    //left
    if((xp-rad) < 2) {
       return 3;
    }
    //right
    if((xp+rad) > width -2 ) {
       return 4;
    }
    return 0;
}

