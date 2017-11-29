class Particle {

	constructor(x0, y0, speed, color) {
		this.x = x0;
		this.y = y0;
		this.diameter = 5;
		this.alpha = 255;
		this.speed = speed;

		this.r = red(color);
		this.b = blue(color);
	}

	show() {

		noStroke();
		fill(this.r, 0, this.b, this.alpha);
		ellipse(this.x, this.y, this.diameter);

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
	}

	finished() {
		return this.alpha <= 0;
	}

}