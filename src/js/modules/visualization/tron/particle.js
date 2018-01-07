export default class Particle {

	constructor(x0, y0, speed, color, diameter) {
		this.x = x0;
		this.y = y0;
		this.diameter = diameter/2;
		this.alpha = 255;
		this.speed = speed;

		this.r = p5.red(color);
		this.b = p5.blue(color);
	}

	show() {

		p5.noStroke();
		p5.fill(this.r, 0, this.b, this.alpha);
		p5.ellipse(this.x, this.y, this.diameter);
		this.update();
	}

	update() {
		if(this.alpha > 200){
			this.alpha -= 1.2*this.speed;
		}
		else if(this.alpha > 150){
			this.alpha -= 1*this.speed;
		}
		else if(this.alpha > 100){
			this.alpha -= 0.8*this.speed;
		}
		else if(this.alpha > 50){
			this.alpha -= 0.5*this.speed;
		}
		else{
			this.alpha -= 0.3*this.speed;
		}
		//this.alpha -= this.speed;
	}

	finished() {
		return this.alpha <= 0;
	}

}

module.exports = Particle;