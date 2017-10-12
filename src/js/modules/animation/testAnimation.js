import Device from '../class/device.js'

const canvasWidth = p5.windowWidth;
const canvasHeight = p5.windowHeight;

export default class TestAnimation {

    constructor(image) {
        this.devices = [];
        this.originalData = [];
        this.readyToShow = false;
        this.image = image;
    }

    update() {

    }

    show() {
        if (!this.readyToShow) return;

        p5.image(this.image, 0, 0, window.innerWidth, window.innerHeight)

        this.devices.forEach(device => {
            p5.ellipse(device.x, device.y, 7)
        });
    }

    setData(data) {
        this.originalData = data;
        this.devices = this.originalData.filter(d => d.floor === 0); // Gefilterd op beganegrond

        this.readyToShow = true;
    }
}
