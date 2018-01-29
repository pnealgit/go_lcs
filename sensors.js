
Rover.prototype.make_sensors = function() {
console.log("num sensors : ",this.num_sensors);

    for (var ns = 0; ns < this.num_sensors; ns++) {
        var s = new Sensor();
        this.sensors.push(s);
    }
} //end of make_sensors

function Sensor() {
      this.xpos = 0;
      this.ypos = 0;
      this.status = 0;
} //end of sensor 

Rover.prototype.set_sensor_positions = function() {

    var tangle = this.angle- this.delta_radians;
    var sensor_spacing = this.antenna_length / this.num_sensors_per_antenna;
    var knt = 0;
    var sr = this.r+this.antenna_length;
    //going from the end in
    for (var is=0;is<this.num_sensors_per_antenna;is++) {
      sr -= sensor_spacing * is;
      for (var ia =0;ia<this.num_antennae;ia++) {
        //going towards center so first num_antennas are the ends of antennas
        this.sensors[knt].ypos = Math.round(this.y + (sr * Math.sin(tangle)));
        this.sensors[knt].xpos = Math.round(this.x + (sr * Math.cos(tangle)));
        tangle += this.delta_radians;
        knt++;
      } 
    }
} 

Rover.prototype.get_sensor_data = function(id) {
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
            this.sensors[ss].status = status; //greater than wall index
            continue;
        }
        //status = check_other_rovers(id,s.xpos,s.ypos);
        //if (status == 6 ) {
         //   this.sensors[ss].status = status;
          //  continue;
        //}
    } //end of loop on sensors
}
//end of get_sensor_data function

function check_food(xp,yp) {
        var status = 0;
        for (var i = 0; i < num_foods; i++) {
            f = foods[i];
            test = f.r;
            dist = Math.hypot((f.x - xp), (f.y - yp));
            if (dist <= test) {
                status = 1;
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

