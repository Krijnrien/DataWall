export default class MqttService {
    
    constructor(onData, onError) {
        this.onData = onData;
        this.onError = onError;
    }

    subscribe(url) {
        fetch(url)
            .then(response => response.json()
            .catch(error => this.onError(error))
            .then(data => this.onData(data))
            .catch(error => this.onError(error)))
    }

    unsubscribe() {

    }
}
