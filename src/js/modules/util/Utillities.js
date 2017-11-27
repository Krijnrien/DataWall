/* global p5: true*/

import DeviceBuilder from '../class/deviceBuilder';

export default class Utillities {
  static normalizeData(data) {
    const maxX = Math.max(...data.map(d => d.x));
    const maxY = Math.max(...data.map(d => d.y));
    const canvasWidth = p5.windowWidth;
    const canvasHeight = p5.windowHeight;

    return data.map(d => new DeviceBuilder()
      .setX(p5.map(d.x, 0, maxX, 0, canvasWidth))
      .setY(p5.map(d.y, 0, maxY, 0, canvasHeight))
      .setFloor(d.z)
      .setUserType(d.userType)
      .setHash(d.hash)
      .build(),
    );
  }
}
