import Visualization from './Visualization';
import Device from './tron/device';
import Unit from './tron/unit';

export default class Tron extends Visualization {
    

  constructor(manager) {
    super(manager);
    this.name = 'Tron';
    this.Devices = [];
    this.Units = [];
  }

  refresh() {

		if(this.Devices.length <= 0){
			if(manager.previousData != null){
				this.fetchToUnits(manager.previousData);
			}
		}

    this.fetchToUnits(manager.currentData);
  }

  update() {
    this.Devices.forEach((device) => {
        device.go();
    });
  }

  show() {
    p5.clear();
    this.Devices.forEach((device) => {
        device.show();
    });


    p5.text(`${p5.frameRate().toFixed(2)} FPS`, 40, 40);
  }

	// weird methods
  fetchToUnits(data) {

		this.Units = [];
		for(let i = Object.keys(data).length - 1; i >= 0; i--){
			console.log("Current id: " + data[i].floor);

			if(data[i].floor == 0){

				let currentDevice = data[i];
				let temp = new Unit(currentDevice.x, currentDevice.y, currentDevice.hash);
				this.Units.push(temp);
			}
		}

		console.log("Expected number of units: " + this.Units.length);
		
		if(this.Devices.length == 0) this.breakDownUnitsToDevices(this.Units);
		else this.updateDestinations(this.Units);
  }

  breakDownUnitsToDevices(units) {

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
						if(this.onSameLocation(newDevice, units[i])){
							newDevice.units.push(units[i]);
							units[i].processed = !units[i].processed;
						}
					}
				}
			}
			
			if(deviceWasCreated) this.Devices.push(newDevice);
		}
		//console.log("BRK- Devices length: " + Devices.length + " Max:  " + findMax(Devices) + " Units: " + getTotalUnits(Devices));
		console.log("Max: " + this.findMax(this.Devices));
  }

  onSameLocation(device, unit){

		let distance = 5;
		let xdif = Math.abs(device.x - unit.x);

		if(xdif <= distance){
			let ydif = Math.abs(device.y - unit.y);
			if (ydif <= distance) return true;
		}
		return false;
  }

  updateDestinations(newunits){

		let newDevices = [];
		for(let i = this.Devices.length - 1; i >= 0; i--){
				
			let currentDevice = this.Devices[i];
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

							if(this.onSameLocation(currentUnit, currentUnitNeighboor)){
								organizedDevice.units.push(currentUnitNeighboor);
								newDevices[j].units[i].processed = true;
							}
						}
					}
				}

				if(deviceWasCreated) organizedDevices.push(organizedDevice);
			}
		}

		//console.log("UPD- Devices length: " + this.Devices.length + " Max:  " + this.findMax(this.Devices) + " Units: " + this.getTotalUnits(this.Devices) + "\n newDevices length: " + newDevices.length + " Max: " + this.findMax(newDevices) + " Units: " + this.getTotalUnits(newDevices) + "\n organizedDevices length: " + organizedDevices.length + " Max: " + this.findMax(organizedDevices) + " Units: " + this.getTotalUnits(organizedDevices));

		this.Devices = organizedDevices;
		this.breakDownUnitsToDevices(newunits);
	}
	
	findMax(array){
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

	getTotalUnits(array){
		let counter = 0;
		for(let j = array.length-1; j >= 0; j--){
			for(let i = array[j].units.length-1; i>=0; i--){
				counter++;
			}
		}
		return counter;
	}
}

