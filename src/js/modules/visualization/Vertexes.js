/* global p5: true*/

import Visualization from './Visualization';
import Circle from './vertex/circle';

export default class Vertexes extends Visualization {
  constructor(manager) {
    super(manager);
    this.name = 'Vertexes';
    this.Circles = [];
  }

  show() {
    p5.clear();
    // this.refresh();

    // this.fetchToUnits(manager.currentData);
    this.fetchToCircles(manager.currentData);

    for (let j = 0; j < this.Circles.length; j++) {
      for (let k = 0; k < this.Circles.length; k++) {
        if (
          this.Circles[j].x - this.Circles[k].x < 70 &&
          this.Circles[j].x - this.Circles[k].x > -70 &&
          this.Circles[j].y - this.Circles[k].y < 70 &&
          this.Circles[j].y - this.Circles[k].y > -70
        ) {
          const r = p5.random(255);
          const g = p5.random(255);
          const b = p5.random(255);
          p5.stroke(r, g, b);
          p5.line(
            this.Circles[k].x,
            this.Circles[k].y,
            this.Circles[j].x,
            this.Circles[j].y,
          );

          // break;
        }
      }
    }

    this.Circles.forEach((circle) => {
      const r = p5.random(255);
      const g = p5.random(255);
      const b = p5.random(255);
      p5.noStroke();
      p5.fill(r, g, b);
      p5.ellipse(circle.x, circle.y, 10);
    });
    // p5.text(`${p5.frameRate().toFixed(2)} FPS`, 40, 40);
    // this.Circles.clear();
  }
  fetchToCircles(data) {
    this.Circles = [];

    for (let i = manager.currentData.length - 1; i >= 0; i--) {
      const currentCircle = manager.currentData[i];
      const temp = new Circle(currentCircle.x * 2.3, currentCircle.y * 2.3, 4);
      this.Circles.push(temp);
    }
  }
}
