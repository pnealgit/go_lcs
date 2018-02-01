var ws
var pause = false

function get_new_radians(angle_rec) {
    var angle = 0.0;
    var delta = 2.0 * Math.PI /8.0;
    angle = angle_rec.Angle * delta;
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
console.log("message: ",e.data);

      n = e.data.indexOf("angles);
      if (n != -1 ) {
         var response = JSON.parse(e.data)
         angles = response.Angle_records
         for (var iang=0;iang < angles.length;iang++) {
            angle_rec = angles[iang] 
            rovers[angle_rec.Id].angle = get_new_radians(angle_rec);
          } //end of loop on iang
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
    ws.send(stuff);
} //end of function senddata

function setup() {
    make_foods(num_foods);
    reset_food_positions();
    team = new Team(num_rovers,num_inputs);
    console.log("TEAM: ",team)
    //senddata(team);
    rovers = make_rovers(team);
console.log("rovers: ",rovers)

    console.log('after making rovers');
    episode_knt = 0;
    num_episodes = 0;

} //end of setup
    
function updateGameArea() {
    if (pause) {
       return
    }
    if (episode_knt >= 580) {
       var mydata = {};
       reset_rover_positions(rovers);
       mydata['num_episodes'] =  num_episodes;
       senddata(mydata);
       episode_knt = 0;
       reset_food_positions();
       num_episodes++;

} //end of if on episode_knt

    myGameArea.clear();
    update_rovers(team,rovers);
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


