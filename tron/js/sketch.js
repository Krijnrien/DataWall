let x1 = [];
let y1 = [];
let Devices = [];

//let url = "http://localhost:3000/devices";
let url = "dummy-data.json";


function preload() {
	fetchJson(fetchToDevices);
}

function setup() {
	createCanvas(800, 600);

	//getDummyDevices(10);
}

function draw() {

	setBackground();
	for(let i = Devices.length - 1; i > 0; i--){
		Devices[i].goTo(x1[i], y1[i]);
		Devices[i].show();
	}
}

function fetchJson(func){
	fetch(url).
      then((res) => {
        return res.json();
      }).
      then((json) => {
      	console.log(json);
      	func(json);
      	return json;
      }).
      catch((err) => {
      	console.log(err);
      });
}

function fetchToDevices(json) {

	for(let i = Object.keys(json).length - 1; i > 0; i--){

		let times = 1.2;
		if(json[i].z == 0)
		{
			let temp = new Device(json[i].x * times, json[i].y * times);
			Devices.push(temp);
		}
		else if(json[i].z == 1)
		{
			x1.push(json[i].x * times);
			y1.push(json[i].y * times);
		}
	}
}

function getDummyDevices(number){

	for(let i = 0; i <= number; i++){

		let x = random(0, width);
		let y = random(0, height);

		temp = new Device(x,y);
		Devices.push(temp);

		let x1p = random(0, width);
		let y1p = random(0, height);
		x1.push(x1p);
		y1.push(y1p);
	}

}

function setBackground(){

	background(245, 245, 245);

	for(let i = 0; i < displayWidth; i=i+120)
	{	
		fill(140,140,140);
		rect(i, 0, 1, displayHeight);
	}

	for(let i = 0; i < displayHeight; i=i+90)
	{
		fill(140,140,140);
		rect(0, i, displayWidth, 1);
	}

	textSize(32);
	fill(70,250,5);
	text("FPS: " + frameRate().toFixed(2), 10, height - 10);
}




