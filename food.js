function Food(num_foods) {
console.log("new Food ",num_foods);

    this.color = 'green';
    this.num_foods = num_foods;
    this.r = 15;
    this.morsels = [];
    this.make_morsels();
}

Food.prototype.make_morsels = function() {
    var start = (2*this.r)+10; //adjust for radius of food
      
    console.log("width,height",width,height); 
    var px = 0;
    var py = 0;
 
    for(var im= 0;im<this.num_foods;im++) {
        px = getRandomInt(start,width-start);
        py = getRandomInt(start,height-start);

        var junk = new Morsel(im,px,py);
        this.morsels.push(junk);
    }
} //end of make_morsels


Food.prototype.update_morsels = function() {
    for(var im=0;im<this.morsels.length;im++) {
        this.morsels[im].update();
    }
} //end of 
Food.prototype.draw_morsels = function() {

    for(var im=0;im<this.morsels.length;im++) {
        this.morsels[im].draw();
    }
} //end of draw_morsels

function Morsel(id,xpos,ypos) {
    this.id = id;
    this.xpos  = xpos;
    this.ypos  = ypos;
    this.r  = getRandomFloat(5,16);
}

Morsel.prototype.draw = function() {
    if (this.r <=0.0) {
        return;
    }

    ctx = myGameArea.context;
    ctx.beginPath();
    ctx.arc(this.xpos,this.ypos,this.r,0,2*Math.PI);
    ctx.fillStyle = "green";
    ctx.fill();
    ctx.strokeStyle = '#ff0000';
    ctx.stroke();
    ctx.closePath();
}
Morsel.prototype.update = function() {
}

