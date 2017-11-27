import Device from './device';

export default class DeviceBuilder {

  constructor(x, y, floor, userType, hash) {
    this.x = x || undefined;
    this.y = y || undefined;
    this.floor = floor || undefined;
    this.userType = userType || undefined;
    this.hash = hash || undefined;
  }

  setX(x) {
    this.x = x;
    return this;
  }

  setY(y) {
    this.y = y;
    return this;
  }

  setFloor(floor) {
    this.floor = floor;
    return this;
  }

  setUserType(userType) {
    this.userType = userType;
    return this;
  }

  setHash(hash) {
    this.hash = hash;
    return this;
  }

  build() {
    if (this.allParamtersComplete()) {
      return new Device(
        this.x,
        this.y,
        this.floor,
        this.userType,
        this.hash,
      );
    }

    throw new Error(`${this.constuctor.name}: Incomplete paramter(s) in new Device`);
  }

  allParamtersComplete() {
    return this.floor !== undefined &&
      this.hash !== undefined &&
      this.userType !== undefined &&
      this.x !== undefined &&
      this.y !== undefined;
  }
}
