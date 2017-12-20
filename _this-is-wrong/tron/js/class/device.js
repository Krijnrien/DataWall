class Device {

	constructor(x0,y0) {
		this.x = x0;
		this.y = y0;

		this.diameter = 10;							  // diameter of the device
		this.speed = 0.7;								  // speed of movement of the device

		if(random(-1,1) >= 0) {
			this.trailColor = color(255,0,0); 		  // @trailColor: color variable
			this.trailColorString = "red";			  // @trailColorString: string indicator of color
		}
		else {
			this.trailColor = color(0,0,255); 
			this.trailColorString = "blue";
		}
		this.changeColorPeriod = random(100,200);	  // How often device should change color
		this.lastChangedColorTime = 0;				  // keeps track of last change of color
		this.colorChangeSpeed = 20;					  // speed of color change of the device (lower -> slower)

		this.deviceColor = color(8, 20, 33);		  // color of the device

		this.particles = [];						  // stores particles representing trail of the device
		this.particleFrequency = 7; 				  // per how many frames new particle appears
		this.particleCounter = 0;					  // helps to control particles

		this.changeDirectionPeriod = random(50, 150); // How often device should change direction
		this.lastChangedDirectionTime = 0; 			  // keeps track of last change of direction
		this.currentDirection = random(-1, 1); 		  // signs direction of device

		this.units = [];
	}

	show() {
		noStroke();
		fill(this.deviceColor);
		ellipse(this.x, this.y, this.getSize());
	}

	goTo(x1, y1) {

		this.simpleMovement(x1,y1);
		
		if(!this.arrivedTo(x1,y1)) this.growTrail();


		this.displayTrail();

		this.changeColor();
	}

	simpleMovement(x1, y1){

		this.lastChangedDirectionTime += 1;

		if(this.lastChangedDirectionTime >= this.changeDirectionPeriod)
		{
			this.currentDirection *= -1;
			this.lastChangedDirectionTime = 0;
		}

		if(this.currentDirection < 0){

			// simple movement (first x, then y)
			if(!this.inDistanceOf("x", x1)){
				if(this.x < x1) this.x += this.speed;
				else if(this.x > x1) this.x -= this.speed;
			}
			else if(!this.inDistanceOf("y", y1)){
				if(this.y < y1) this.y += this.speed;
				else if(this.y > y1) this.y -= this.speed;
			}
		} else {

			// first y, then x
			if(!this.inDistanceOf("y", y1)){
				if(this.y < y1) this.y += this.speed;
				else if(this.y > y1) this.y -= this.speed;
			}
			else if(!this.inDistanceOf("x", x1)){
				if(this.x < x1) this.x += this.speed;
				else if(this.x > x1) this.x -= this.speed;
			}
		}
	}
	
	growTrail(){

		// adding new particles
		if(this.particleCounter >= this.particleFrequency){

			let p = new Particle(this.x, this.y, this.speed, this.trailColor, this.getSize());
			this.particles.push(p);
			this.particleCounter = 0;
		}
		this.particleCounter += 1;
	}

	displayTrail(){

		// displaying particles
		for(let i = this.particles.length - 1; i >= 0; i--){

			this.particles[i].show();

			if(this.particles[i].finished()){
				this.particles.splice(i, 1);
			}
		}
	}

	changeColor(){

		// control of periodic color change
		this.lastChangedColorTime += 1;

		if(this.lastChangedColorTime >= this.changeColorPeriod)
		{
			if(this.trailColorString == "red") this.trailColorString = "blue";
			else this.trailColorString = "red";
			this.lastChangedColorTime = 0;
		}

		let r = red(this.trailColor);
		let b = blue(this.trailColor);

		if(this.trailColorString == "blue")
		{
			if(r > 0){
				r -= this.colorChangeSpeed;
				if(r < 122) b += this.colorChangeSpeed;
			}
			else if(r >= 0 && b < 255)
			{
				b += this.colorChangeSpeed;
			}
		}
		else {
			if(b > 0){
				b -= this.colorChangeSpeed;
				if(b < 122) r += this.colorChangeSpeed;
			}
			else if(b >= 0 && r < 255)
			{
				r += this.colorChangeSpeed;
			}
		}

		this.trailColor = color(r, 0, b);
	}

	inDistanceOf(axis, xy){
		let distance = 3;
		let toReturn = false;

		let check;

		if(axis == "x"){
			check = Math.abs(this.x - xy);
		}

		if(axis == "y"){
			check = Math.abs(this.y - xy);
		}
		
		if (check <= distance) return true;
		return false;
	}

	getSize(){
		if(this.units.length == 1){
			return this.diameter;
		}
		else return this.diameter * this.units.length * 0.5;
	}

	arrivedTo(x1, y1){
		return ((this.inDistanceOf("x", x1))&&(this.inDistanceOf("y",y1)));
	}
}