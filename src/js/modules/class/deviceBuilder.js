import Device from './device.js';

export class DeviceBuilder {

    constructor(x, y, floor, userType, hash) {
        this.x = undefined;
        this.y = undefined;
        this.floor = undefined;
        this.userType = undefined;
        this.hash = undefined;
    };

    setX(x) {
        this.x = x;
        return this;
    };

    setY(y) {
        this.y = y;
        return this;
    };

    setFloor(floor) {
        this.floor = floor;
        return this;
    };

    setUserType(userType) {
        this.userType = userType;
        return this;
    };

    setHash(hash) {
        this.hash = hash;
        return this;
    };

    build() {
        if (this.allParamtersComplete()) {
            return new Device(
                this.x,
                this.y,
                this.floor,
                this.userType,
                this.hash
            );
        } else {
            console.warn("Incomplete paramter(s) in new Device");
        };
    };

    allParamtersComplete() {
        return this.floor != undefined &&
            this.hash != undefined &&
            this.userType != undefined &&
            this.x != undefined &&
            this.y != undefined;
    };
};
