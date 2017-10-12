import Device from '../class/device.js'

let lastTime = 0;

export default class TestAnimation {

    constructor(images) {
        this.devices = [];
        this.originalData = [];
        this.readyToShow = false;
        this.images = images;
        this.floor = -1;
    }

    update() {
        let currentTime = new Date().getTime();

        if (currentTime - lastTime > 2000) {

            this.floor = ++this.floor % this.images.length;

            this.devices = this.originalData.filter(d => d.floor === this.floor);

            lastTime = currentTime;
        }
    }

    show() {
        if (!this.readyToShow) return;

        p5.image(this.images[this.floor], 0, 0, window.innerWidth, window.innerHeight)

        this.devices.forEach(device => {
            p5.ellipse(device.x, device.y, 7)
        });
    }

    setData(data) {
        this.originalData = data;
        this.devices = this.originalData.filter(d => d.floor === this.floor);

        lastTime = new Date().getTime();
        this.readyToShow = true;
    }
}
