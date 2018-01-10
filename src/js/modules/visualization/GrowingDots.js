/* global p5: true*/

import Visualization from './Visualization';

export default class GrowingDots extends Visualization {

  constructor(manager) {
    super(manager);
    this.name = 'growingdots';
    this.objectHistory = [];
  }

  update() {
    this.objectHistory = this.manager.currentData.map(current => {
      return {
        previous: this.manager.previousData.filter(previous => previous.hash === current.hash)[0],
        current: current
      }
    });
  }

  show() {

    p5.clear();
    this.objectHistory.forEach((device) => {

      if (device.current.userType === 1) {
        p5.fill(255, 204, 0);
        p5.ellipse(device.current.x, device.current.y, 10);
      } else {
        p5.fill(64, 224, 208, 75);
        p5.ellipse(device.current.x, device.current.y, 7);
      }
      if (device.previous !== undefined){
        const differenceX = device.current.x - device.previous.x;
        const differenceY = device.current.y - device.previous.y;
        if (differenceX > -6 && differenceX < 6 && differenceY > -6 && differenceY < 6){
          // Instead of p5.frameCount an actual counter that tracks the size could be used. 
          // Since there is no API available that receives real data, this is being used as a replacement for the demo.
          var sizeCounter = p5.frameCount / 4;
          if (sizeCounter>100){
            sizeCounter = 100;
          }

          p5.ellipse(device.current.x, device.current.y, sizeCounter);
        }
      }
    });
  }
}
