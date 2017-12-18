import Listener from './Listener';

export default class DeviceService {

  constructor() {
    this.listeners = [];
    this.data = [];
    this.previousData = [];
    this.interval = 20 * 1000;
    this.url = '../../../assets/json/dummy-data.json';
  }

  subscribe(listener) {
    this.listeners.push(listener);
  }

  update() {
    fetch(this.url)
      .then(response => response.json())
      .then((response) => {
        this.previousData = this.data;
        this.data = response;
      })
      .then(() => {
        this.listeners.forEach((listener) => {
          listener.onUpdate(this, { data: this.data, previousData: this.previousData });
        });
      });
  }
}

export const initializeService = (onUpdate) => {
  const service = new DeviceService();
  const listener = new Listener();

  listener.onUpdate = onUpdate;

  service.subscribe(listener);

  service.update();

  setTimeout(() => {
    service.update();
  }, service.interval);

  return service;
};
