import TestVisualization from './TestVisualization';

export default class VisualizationManager {

  constructor() {
    this.currentIndex = 0;
    this.data = [];
    this.previousData = [];
    this.visualizations = [
      new TestVisualization(this),
    ];
  }

  get current() {
    return this.visualizations[this.currentIndex];
  }

  next() {
    this.currentIndex = this.currentIndex += 1 % this.visualizations.length;
  }
}
