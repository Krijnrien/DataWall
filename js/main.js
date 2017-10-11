const url = "../json/dummy-data.json"
const magnification = 2;

let service
let devices = []
let data = [];
let readyToDraw = false
let floorSlider
let lastTime;
let level = 0;
let levelP;

function setup() {
    createCanvas(650 * magnification, 450 * magnification)
    background(255)

    this.lastTime = new Date().getTime();

    this.levelP = createP(level);

    this.service = new MqttService(
        data => this.onDataReceived(data),
        error => {
            console.log(error)
            this.readyToDraw = false
        }
    );

    this.service.subscribe(url)
}

function draw() {
    if (!this.readyToDraw) return

    this.updateLevel()

    background(255)

    this.devices.forEach(device => {
        if (device.userType === 1) {
            fill(255, 204, 0)
        } else {
            fill(255, 255, 255, 50)
        }

        ellipse(device.x * magnification, device.y * magnification, 5 * magnification)
    })
}

function onDataReceived(data) {
    this.data = data;

    this.devices = this.data.filter(d => d.z == level)
    this.readyToDraw = true
    this.draw()
}


function updateLevel() {
    let currentTime = new Date().getTime();

    if (currentTime - this.lastTime > 2000) {
        this.devices = this.data.filter(d => d.z == level)
        this.lastTime = currentTime;
        this.levelP.html(level)
        level++;
        
       

        if (level > 4) {
            level = 0;
        }
    }
}