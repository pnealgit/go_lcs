var Team = function(num_rovers,num_inputs) {
    this.team_name = "make_team";
    this.num_rovers = num_rovers;
    this.num_inputs = num_inputs;
}
//end of function

function make_rovers(team) {
    rovers = [];
    for (var ri = 0; ri < team.num_rovers; ri++) {
        rovers[ri] = new Rover(ri);
    }
    reset_rover_positions(rovers)
    return rovers;
}
//end of function make_rovers

function Rover(id) {
    this.id = id;
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

function update_rovers(team, rovers) {

    all_rovers = {};
    all_rovers['status'] = "state";

    best_score = -9999.9
    worst_score = 99999.9

    for (var i = 0; i < team.num_rovers; i++) {
        this.state = [];
        my_data = {};
        my_data['id'] = i;
        my_data['reward'] = 0;

        rovers[i].get_sensor_data(i);
        rrr = rovers[i].get_reward();
       
        my_data['state'] = rovers[i].state;
        my_data['reward'] = rrr

        rovers[i].move();
        rovers[i].draw();
    }
    //end of loop on rovers

    //all_rovers['all_recs'] = all_recs;
    //senddata(all_rovers);
    senddata(my_data);
}
//end of function 

Rover.prototype.get_reward = function() {
    //1-4 wall
    //5 is food
    //6 is other

    this.state = [];

    //this.state.push(this.x/width)
    //this.state.push(this.y/height)
    this.state.push(Math.round(this.x));
    this.state.push(Math.round(this.y));
    //var sx = this.x.toString(2);
    //var sy = this.y.toString(2);
    //sx ="0000000000".substr(sx.length)+sx;
    //sy ="0000000000".substr(sy.length)+sy;

    //this.state.push(sx)
    //this.state.push(sy);
    var new_reward = 0;
    //sensors first
    var s = {} ;
    var tx = 0.0;

    for(var ss = 0;ss < this.num_sensors; ss++) {
        var final_status = 0;
        s = this.sensors[ss];
        var stemp = s.status.toString(2);
        stemp ="0000".substr(stemp.length)+stemp;
        //this.state.push(stemp); 
        this.state.push(s.status);
        //walls numbered 1-4
        if (s.status > 0 && s.status < 5) {
           new_reward += -9;
        }
        //food
        if (s.status == 5) {
           new_reward += 50;
        }

        //others
        if (s.status == 6) {
           new_reward += -1;
        }
     } //end of loop on sensors
    //now do checks on rover center
 
    var border_status = 0;
    border_status = check_borders(this.x,this.y,this.r);
    if (border_status > 0 && border_status < 5 ) {
       new_reward += -1;
    }

    //food
    var food_status = 0;
    food_status = check_food(this.x,this.y);
    if (food_status > 0) {
       new_reward += 100;
    }

    //if hit another rover
    //var rover_status = 0;
    //rover_status = check_other_rovers(this.id,this.x,this.y);
    //if (rover_status > 0) {
     //  new_reward += -5;
    //}

    //just for breathing 
    new_reward += 1;
    return new_reward;
}
//end of reward

function reset_rover_positions(rovers) {
    sum = 0;
    best = -9999;
    worst = 9999;

    for (var nn = 0; nn < num_rovers; nn++) {
        rovers[nn].reset_position();

        r = rovers[nn].reward;
        if (r > best) {
            best = r;
        }
        if (r < worst) {
            worst = r;
        }
        sum += r
        rovers[nn].reward = 0;
    }
    //end of loop
    console.log("SUM: \t", sum, "\tBEST:\t", best, "\tWORST:\t", worst);
}

Rover.prototype.reset_position = function() {
    this.x = getRandomInt(50,width-50);
    this.y = getRandomInt(50,height-50);
    junk = getRandomInt(0, 8)
    this.angle = junk *  Math.PI / 4;
    this.velocity = 2.0;
}

