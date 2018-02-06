function make_rover() {
    rover = {}
    rover = new Rover();
    rover.make_sensors();
    return rover;
}
//end of function make_rover

function Rover() {
    this.x = getRandomInt(50,width-50);
    this.y = getRandomInt(50,height-50);
    this.r = 10;
    this.sensors =  [];
    this.num_antennae = 3;
    this.velocity = 2.0;
    this.delta_radians =  Math.PI/4.0
    this.state = [];
    this.received_angle = getRandomInt(0,4)
    this.last_received_angle = getRandomInt(0,4)
    this.angle = 2.0*Math.PI * Math.random();
    this.dx = this.velocity * Math.cos(this.angle);
    this.dy = this.velocity * Math.sin(this.angle);

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
    //rover.make_sensors()
    rover.get_sensor_data();
    //rover.get_reward();
    my_data['status'] = "State";
    my_data['sensors'] = rover.sensors;
    my_data['last_angle'] = rover.last_received_angle;
    my_data['x'] = rover.x;
    my_data['y'] = rover.y;

    rover.move();
    rover.draw();
    senddata(my_data);
} //end of function 

Rover.prototype.get_reward = function() {
    //1-4 wall
    //5 is food
    //6 is other

    //calculate reward and distance in brain
    this.state = {};
    this.state.last_angle = this.last_received_angle
    this.state.X = Math.round(this.x);
    this.state.Y = Math.round(this.y);
    this.state.Sensors = this.sensors
}


