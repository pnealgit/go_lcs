function make_rover() {
    rover = {}
    rover = new Rover();
    reset_rover_positions()
    return rover;
}
//end of function make_rover

function Rover() {
    this.x = getRandomInt(50,width-50);
    this.y = getRandomInt(50,height-50);
    this.r = 10;
    this.sensors = [];
    this.num_antennae = 3;
    this.num_sensors_per_antenna = 2;
    this.num_sensors = this.num_antennae * this.num_sensors_per_antenna;
    this.velocity = 2.0;
    this.antenna_length =  40
    this.delta_radians =  Math.PI/4.0

    this.state = [];
    this.reward = 0.0;
    this.angle = 2.0*Math.PI * Math.random();
    this.last_food_x = 0.0;
    this.last_food_y = 0.0;

    this.dx = this.velocity * Math.cos(this.angle);
    this.dy = this.velocity * Math.sin(this.angle);
    this.make_sensors();

    this.move = function() {
        this.dx = this.velocity * Math.cos(this.angle)
        this.dy = this.velocity * Math.sin(this.angle)
        var status = -9;
        var tx = this.x + this.dx;
        var ty = this.y + this.dy;
        status = check_borders(tx,ty,this.r);
        if ( status < 1 || status > 5 )  {
            this.x += this.dx;
            this.y += this.dy;
        } 
    }

    this.draw = function() {
        ctx = myGameArea.context;
        ctx.beginPath();
        ctx.arc(this.x, this.y, this.r, 0, 2 * Math.PI);
        ctx.fillStyle = "red";
        ctx.fill();

        ctx.beginPath();
        ctx.strokeStyle = '#000000';
        tangle = this.angle- this.delta_radians;
       
        this.set_sensor_positions();
	for (var s=0;s<this.num_antennae;s++) {
            ctx.moveTo(this.x,this.y);
            ctx.lineTo(this.sensors[s].xpos, this.sensors[s].ypos)
        }
        //end of loop on sensors
        ctx.stroke();
        ctx.closePath();
    }
    //end of rover draw
}
//end of Rover function

function update_rover() {

    my_data = {};
    rover.get_sensor_data();
    rover.get_reward();
if (rover.reward > 1 ) {
    console.log("ROVER REWARD: ",rover.reward);
}
    my_data['status'] = "State";
 
    my_data['state'] = rover.state;
    my_data['reward'] = rover.reward
    rover.move();
    rover.draw();
    senddata(my_data);
} //end of function 

Rover.prototype.get_reward = function() {
    //1-4 wall
    //5 is food
    //6 is other

    this.state = [];

    this.state.push(Math.round(this.x));
    this.state.push(Math.round(this.y));
    this.reward = 0;
    //sensors first
    var s = {} ;
    var tx = 0.0;

    for(var ss = 0;ss < this.num_sensors; ss++) {
        var final_status = 0;
        s = this.sensors[ss];

if (s.status > 4) {
    console.log("sensor : ",s)
}

        var stemp = s.status.toString(2);
        stemp ="0000".substr(stemp.length)+stemp;
        //this.state.push(stemp); 
        this.state.push(s.status);
        //walls numbered 1-4
        if (s.status > 0 && s.status < 5) {
           this.reward += .01;
           continue
        }
        //food
        if (s.status == 5) {
           this.reward += 10.0;
        }
        if (s.status == 0) {
             this.reward+= 1.0
        }
     } //end of loop on sensors
    //now do checks on rover center
 
    var border_status = 0;
    border_status = check_borders(this.x,this.y,this.r);
    if (border_status > 0 && border_status < 5 ) {
       //this.reward += 0.01;
    }

    //food
    var food_status = 0;
    food_status = check_food(this.x,this.y);
    if (food_status > 0) {
       //this.reward += 20.0;
    }

    //just for breathing 
    this.reward += 0.01;
}
//end of reward

function reset_rover_positions(rovers) {
    sum = 0;
    best = -9999;
    worst = 9999;

        rover.reset_position();

        r = rover.reward;
        if (r > best) {
            best = r;
        }
        if (r < worst) {
            worst = r;
        }
        sum += r
        rover.reward = 0;
}

Rover.prototype.reset_position = function() {
    this.x = getRandomInt(50,width-50);
    this.y = getRandomInt(50,height-50);
    junk = getRandomInt(0, 8)
    this.angle = junk *  Math.PI / 4;
    this.velocity = 2.0;
}

