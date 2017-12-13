export default class Visualization {

  constructor(manager) {
    this.hasShownUpdateWarning = false;
    this.hasShownDrawWarning = false;
    this.manager = manager;
    this.name = 'undefined';
  }

  update() {
    if (!this.hasShownUpdateWarning) {
      console.warn(`${this.constructor.name}: Please implement 'update'`);
      this.hasShownUpdateWarning = true;
    }
  }

  show() {
    if (!this.hasShownDrawWarning) {
      console.warn(`${this.constructor.name}: Please implement 'update'`);
      this.hasShownDrawWarning = true;
    }
  }
}
