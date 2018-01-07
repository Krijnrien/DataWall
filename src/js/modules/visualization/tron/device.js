import Particle from './particle';

export default class Device {

	constructor(x0,y0,hash) {
		this.x = x0;
		this.y = y0;
		this.hash = hash;

		this.diameter = 20;							  // diameter of the device
		this.speed = 3;//0.7;								  // speed of movement of the device

		if(p5.random(-1,1) >= 0) {
			this.trailColor = p5.color(255,0,0); 		  // @trailColor: color variable
			this.trailColorString = "red";			  // @trailColorString: string indicator of color
		}
		else {
			this.trailColor = p5.color(0,0,255); 
			this.trailColorString = "blue";
		}
		this.changeColorPeriod = p5.random(100,200);	  // How often device should change color
		this.lastChangedColorTime = 0;				  // keeps track of last change of color
		this.colorChangeSpeed = 20;					  // speed of color change of the device (lower -> slower)

		this.deviceColor = p5.color(255, 255, 255);		  // color of the device

		this.particles = [];						  // stores particles representing trail of the device
		this.particleFrequency = 7; 				  // per how many frames new particle appears
		this.particleCounter = 0;					  // helps to control particles

		this.changeDirectionPeriod = p5.random(50, 150); // How often device should change direction
		this.lastChangedDirectionTime = 0; 			  // keeps track of last change of direction
		this.currentDirection = p5.random(-1, 1); 		  // signs direction of device

		this.units = [];
		this.x1 = null;
		this.y1 = null;
	}

	setDestination(x,y){
		this.x1 = x;
		this.y1 = y;
	}

	show() {

		// trail settings
		if(!this.arrived()) this.growTrail();
		this.displayTrail();
		this.changeColor();

		// device
		p5.noStroke();
		p5.fill(this.deviceColor);
		p5.ellipse(this.x, this.y, this.getSize());
	}

	go() {
		this.simpleMovement(this.units[0].x, this.units[0].y);
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

		let r = p5.red(this.trailColor);
		let b = p5.blue(this.trailColor);

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

		this.trailColor = p5.color(r, 0, b);
	}

	inDistanceOf(axis, xy){
		let distance = 8;

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

		//return 20;
		if(this.units.length == 1){
			return this.diameter;
		}
		else {
			let multiplier = 1;
			this.units.forEach(unit => {
				multiplier += 0.2;
			});
			return this.diameter * multiplier;
		}
	}

	arrived(){
		let x1 = this.units[0].x;
		let y1 = this.units[0].y;
		return ((this.inDistanceOf("x", x1))&&(this.inDistanceOf("y",y1)));
	}

	finished(){
		return (this.particles == 0)
	}
}

module.exports = Device;