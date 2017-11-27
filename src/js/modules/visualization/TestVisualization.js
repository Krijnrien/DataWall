/* global p5: true*/

import Visualization from './Visualization';

export default class TestVisualization extends Visualization {

  constructor(manager) {
    super(manager);
    this.data = [];
  }

  update() {
    this.manager.data.forEach((device, i) => {
      this.data[i] = device;
      this.data[i].x += p5.random(-5, 5);
      this.data[i].y += p5.random(-5, 5);
    });
  }

  show() {
    p5.clear();

    this.data.forEach((device) => {
      if (device.userType === 1) {
        p5.fill(255, 204, 0);
        p5.ellipse(device.x, device.y, 10);
      } else {
        p5.fill(64, 224, 208, 75);
        p5.ellipse(device.x, device.y, 7);
      }
    });

    p5.fill(186, 42, 42);

    p5.textSize(32);
    p5.text(`${p5.frameRate().toFixed(2)} FPS`, 40, 40);
  }
}
