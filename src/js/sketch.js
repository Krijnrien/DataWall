import p5 from 'p5';
import 'p5/lib/addons/p5.sound';
import 'p5/lib/addons/p5.dom';

import MqttSerivce from './modules/service/mqttService.js';
import Device from './modules/class/device.js';
import { DeviceBuilder } from './modules/class/deviceBuilder.js';
import TestAnimation from './modules/animation/testAnimation.js';

let animation;
let service;

const sketch = (p5) => {
  // Variables scoped within p5
  const canvasWidth = p5.windowWidth;
  const canvasHeight = p5.windowHeight;

  // make library globally available
  window.p5 = p5;

  // Setup function
  // ======================================
  p5.setup = () => {
    let canvas = p5.createCanvas(canvasWidth, canvasHeight);
    canvas.parent('sketch');
    p5.frameRate(24);

    animation = new TestAnimation(p5.loadImage("assets/img/floor0.png"));

    service = new MqttSerivce(
      data => {
        let maxX = Math.max(...data.map(d => d.x))
        let maxY = Math.max(...data.map(d => d.y))

        animation.setData(
          data.map(d => new DeviceBuilder()
            .setX(p5.map(d.x, 0, maxX, 0, canvasWidth))
            .setY(p5.map(d.y, 0, maxY, 0, canvasHeight))
            .setFloor(d.z)
            .setUserType(d.userType)
            .setHash(d.hash)
            .build()
          )
        );
      },
      error => console.error(error)
    );

    service.subscribe("assets/json/dummy-data.json");
  }

  // Draw function
  // ======================================
  p5.draw = () => {
    if (animation) {
      animation.update();
      animation.show();
    }
  }
}

export default sketch;
