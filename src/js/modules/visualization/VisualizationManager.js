import TestVisualization from './TestVisualization';

export default class VisualizationManager {

  constructor() {
    this._currentData = [];
    this._previousData = [];
    this._floor = 0;
    this.currentIndex = 0;
    // List of visualizations the user wants to see
    // Defualts to all visualizations.
    this.userPreferences = ['all'];

    // Add your visualization here!
    // Also, dont forget to add 'this' as paramter as this class supplies the data to your visualization.
    this.visualizations = [
      new TestVisualization(this),
    ];
  }

  next() {
    console.log(this.filterVisualizations.length);
    this.currentIndex = (this.currentIndex + 1) % this.filterVisualizations.length;
  }

  /* Getters and Setters */
  get filterVisualizations() {
    if (this.userPreferences.includes('all')) {
      return this.visualizations;
    }

    return this.visualizations.filter(v => this.userPreferences.includes(v.name));
  }

  get current() {
    return this.filterVisualizations[this.currentIndex];
  }

  set currentData(data) {
    this._currentData = data;
  }

  get currentData() {
    return this._currentData.filter(d => d.floor === this.floor);
  }

  set previousData(data) {
    this._previousData = data;
  }

  get previousData() {
    return this._previousData.filter(d => d.floor === this.floor);
  }

  get floor() {
    return this._floor;
  }

  set floor(floor) {
    if (floor > 3) this._floor = 3;
    else this._floor = floor;
  }
}
