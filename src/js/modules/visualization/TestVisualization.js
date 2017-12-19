/* global p5: true*/

import Visualization from './Visualization';

export default class TestVisualization extends Visualization {

  constructor(manager) {
    super(manager);
    this.name = 'test';
  }

  update() {
    // Useless method to keep the compiler quiet.
    // Should be filled in with actually updating the data.
    this.data = this.data;
  }

  show() {
    p5.clear();

    this.manager.currentData.forEach((device) => {
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
