var ws
var pause = false

function get_new_radians(angle_rec) {
    if (angle_rec == 0 ) {
        return 0;
    }

    var angle = 0.0;
    var delta = 2.0 * Math.PI /4.0;
    angle = angle_rec * delta;
    return angle;
}

function WebsocketStart() {

    ws = new WebSocket("ws://localhost:8081/talk")

    ws.onopen = function(evt) {
      senddata('CONNECTION MADE');
      setup();
      myGameArea.start(); 
    }
    ws.onclose = function(evt) {
      console.log('WEBSOCKET CLOSE');
      myGameArea.stop();
      //ws = null;
    }

    ws.onmessage = function(e) {
      n = e.data.indexOf("Angle");
      if (n != -1 ) {
         var response = JSON.parse(e.data)
         rover.angle = get_new_radians(response.Angle);
      } //end of found 'angle'
    } //endo of onmessage

    ws.onerror = function(evt) {
        console.log('onerror ',evt.data);
    }

} //end of WebsocketStart

senddata = function(data) {
    if (pause) {
      return;
    }
    if (!ws) {
        console.log('cannot send data -- no ws');
        return false;
    }
    stuff = JSON.stringify(data);
    console.log("SENDING : ",stuff)
    ws.send(stuff);
} //end of function senddata

function setup() {
    make_foods(num_foods);
    reset_food_positions();
    rover = make_rover();
    console.log("rover: ",rover)
} //end of setup
    
function updateGameArea() {
    if (pause) {
       return
    }

    myGameArea.clear();
    update_rover(rover);
    update_foods();
    episode_knt+= 1;
} //end of updateGameArea

myGameArea = {
    canvas : document.createElement("canvas"),
    start : function() {
        this.millis = 75;  //game intervale milliseconds
        this.canvas.width = width;
        this.canvas.height = height;
        this.context = this.canvas.getContext("2d");
        document.body.insertBefore(this.canvas, document.body.childNodes[0]);
        pause = false;
        this.interval = setInterval(updateGameArea,this.millis);
    },  
    stop : function() {
        pause = true; 
        console.log("STOP !!! ");
        clearInterval(this.interval);
        //ws.close();
    },  
    clear : function() {
        this.context.clearRect(0, 0, this.canvas.width, this.canvas.height);
        this.context.fillStyle = "rgba(255,255,255,255)";
        this.context.fillRect(0,0,this.canvas.width,this.canvas.height);
    } 
}    //end of gamearea


