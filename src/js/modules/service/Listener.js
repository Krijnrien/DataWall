export default class Listener {
  onUpdate(sender, args) {
    console.warn(`${this.constructor.name}: Please implement 'onUpdate'.`, args);
  }
}
