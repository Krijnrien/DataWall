import Listener from './Listener';

export default class DeviceService {

  constructor() {
    this.listeners = [];
    this.data = [];
    this.previousData = [];
    this.interval = 1 * 1000;
    this.url = '../../../assets/json/dummy-data.json';
    this.dummyData = [
      '../../../assets/json/dummy-data.json',
      '../../../assets/json/dummy-data-2.json',
    ];
    this.dummyIndex = 0;
  }

  subscribe(listener) {
    this.listeners.push(listener);
  }

  update() {
    fetch(this.dummyData[this.dummyIndex])
      .then(response => response.json())
      .then((response) => {
        this.previousData = this.data;
        this.data = response;
      })
      .then(() => {
        this.listeners.forEach((listener) => {
          listener.onUpdate(this, { data: this.data, previousData: this.previousData });
        });
      })
      .then(this.dummyIndex = (this.dummyIndex + 1) % this.dummyData.length);
  }
}

export const initializeService = (onUpdate) => {
  const service = new DeviceService();
  const listener = new Listener();

  listener.onUpdate = onUpdate;

  service.subscribe(listener);

  service.update();

  setInterval(() => {
    service.update();
  }, service.interval);

  return service;
};
