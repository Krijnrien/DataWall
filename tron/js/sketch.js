let x1 = [];
let y1 = [];
let Devices = [];
let Units = [];

let time;
let wait = 20000;
let tick = false;

//let url = "http://145.93.119.178:3000/devices";
let url = "http://localhost:3000/devices";
//let url = "dummy-data.json";


function preload() {
	fetchJson(fetchToUnits);
	breakDownUnitsToDevices(Units);
}

function setup() {
	createCanvas(800, 600);
	//getDummyDevices(10);
	
	time = millis();
	console.log("Devices length: " + Devices.length + " Max:  " + findMax(Devices) + " Units: " + getTotalUnits(Devices));
}

function draw() {

	setBackground();
	for(let i = Devices.length - 1; i >= 0; i--){
		let destination = Devices[i].units[0];
		Devices[i].goTo(destination.x, destination.y);
		Devices[i].show();
	}

	if(millis() - time >= wait){
		tick = !tick;
		time = millis();

		fetchJson(fetchToUnits);
	}

	//console.log(findMax());
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

function fetchToUnits(json) {

	Units = [];
	for(let i = Object.keys(json).length - 1; i >= 0; i--){

		if(json[i].z == 0){

			let currentDevice = json[i];
			let temp = new Unit(currentDevice.x, currentDevice.y, currentDevice.hash);
			Units.push(temp);
		}
	}

	console.log("Expected number of units: " + Units.length);
	if(Devices.length == 0) breakDownUnitsToDevices(Units);
	else updateDestinations(Units);
}


function findMax(array){
	let toReturn = 0;
	let temp;
	for(let i = array.length - 1; i >= 0; i--){
		let size = array[i].units.length;
		if(size > toReturn) {
			toReturn = size;
			temp = "i: " + i + " size: " + size;
		}
	}
	return temp;
}

function getTotalUnits(array){
	let counter = 0;
	for(let j = array.length-1; j >= 0; j--){
		for(let i = array[j].units.length-1; i>=0; i--){
			counter++;
		}
	}
	return counter;
}

function updateDestinations(newunits){

	let newDevices = [];
	for(let i = Devices.length - 1; i >= 0; i--){
			
		let currentDevice = Devices[i];
		let newDevice = new Device(currentDevice.x, currentDevice.y);
		for(let i = currentDevice.units.length - 1; i >= 0; i--){
						
			let currentUnit = currentDevice.units[i];
			for(let i = newunits.length - 1; i >= 0; i--){
						
				let currentNewUnit = newunits[i];
				if(currentUnit.hash == currentNewUnit.hash){
					newDevice.units.push(currentNewUnit);	 // created new Device with latest units
					newunits.splice(i,1); 				  	 // delete used units
				}
			}
		}
		if(newDevice.units.length > 0) newDevices.push(newDevice);
	}

	let organizedDevices = [];
	for(let j = newDevices.length - 1; j >= 0; j--){
		
		//let currentDevice = newDevices[j];
		for(let i = newDevices[j].units.length - 1; i >= 0; i--){

			let organizedDevice;
			let deviceWasCreated = false;
			let currentUnit = newDevices[j].units[i];
			if(!currentUnit.processed){
				organizedDevice = new Device(newDevices[j].x,  newDevices[j].y);
				organizedDevice.units.push(currentUnit);
				newDevices[j].units[i].processed = true;
				deviceWasCreated = true;
				for(let i = newDevices[j].units.length - 1; i >= 0; i--){
	
					//console.log("j = " + j + "; i = " + i + "; newDevices[j].units.length: " + newDevices[j].units.length);
					let currentUnitNeighboor = newDevices[j].units[i];
					if(!currentUnitNeighboor.processed){

						if(onSameLocation(currentUnit, currentUnitNeighboor)){
							organizedDevice.units.push(currentUnitNeighboor);
							newDevices[j].units[i].processed = true;
						}
					}
				}
			}

			if(deviceWasCreated) organizedDevices.push(organizedDevice);
		}
	}
	//console.log(newDevices[0].units.length);
	console.log("UPD- Devices length: " + Devices.length + " Max:  " + findMax(Devices) + " Units: " + getTotalUnits(Devices) + "\n newDevices length: " + newDevices.length + " Max: " + findMax(newDevices) + " Units: " + getTotalUnits(newDevices) + "\n organizedDevices length: " + organizedDevices.length + " Max: " + findMax(organizedDevices) + " Units: " + getTotalUnits(organizedDevices));
	//Devices = organizedDevices;
	Devices = organizedDevices;
	breakDownUnitsToDevices(newunits);
}

function breakDownUnitsToDevices(units){

	for(let i = units.length - 1; i >= 0; i--){

		let newDevice;
		let deviceWasCreated = false;
		if(!units[i].processed) {
			newDevice = new Device(units[i].x, units[i].y);
			newDevice.units.push(units[i]);
			units[i].processed = !units[i].processed;
			deviceWasCreated = true;

			for(let i = units.length - 1; i >= 0; i--){
				
				if(!units[i].processed) {
					if(onSameLocation(newDevice, units[i])){
						newDevice.units.push(units[i]);
						units[i].processed = !units[i].processed;
					}
				}
			}
		}
		
		if(deviceWasCreated) Devices.push(newDevice);
	}
	console.log("BRK- Devices length: " + Devices.length + " Max:  " + findMax(Devices) + " Units: " + getTotalUnits(Devices));
}

function onSameLocation(device, unit){

	let distance = 5;
	let xdif = Math.abs(device.x - unit.x);

	if(xdif <= distance){
		let ydif = Math.abs(device.y - unit.y);
		if (ydif <= distance) return true;
	}
	return false;
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

