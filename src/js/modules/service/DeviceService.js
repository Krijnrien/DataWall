import Listener from './Listener';

export default class DeviceService {

  constructor() {
    this.listeners = [];
    this.data = [];
  }

  subscribe(listener) {
    this.listeners.push(listener);

    fetch('../../../assets/json/dummy-data.json')
      .then(response => response.json())
      .then(response => (this.data = response))
      .then(() => this.update());
  }

  update() {
    this.listeners.forEach((listener) => {
      listener.onUpdate(this, { data: this.data, previousData: this.data });
    });
  }
}

export const initializeService = (onUpdate) => {
  const service = new DeviceService();
  const listener = new Listener();

  listener.onUpdate = onUpdate;

  service.subscribe(listener);

  return service;
};
