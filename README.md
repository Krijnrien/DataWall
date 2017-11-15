# DataWall 2017 Testing

## Project setup
1. Make sure you clone the project in your ```src``` folder in your ````GOPATH````. If you don't , the project won't be 
able to link the different packages properly. 
If you don't know where your ```GOPATH``` is, 
open a terminal and run ```go env``` and search for ```GOPATH```.

2. Now install all dependencies by going to the root of the repository and running ```go get .```.

## GraphQL Server Setup
### Prerequisites
1. Make sure you installed ``` node.js ``` on your server. If you haven't make sure to follow the instructons mentioned in node.js [installer download](https://nodejs.org/en/download/)
2. Make sure you have cloned the repository of Testing into your ```src``` folder(Follow ```Project setup instructions above```). If your having problems with that make sure to see the full explanation in [Backend](https://github.com/Krijnrien/DataWall/tree/Backend)
3. Write the command to install all dependencies in the Graphql folder in the test version ```npm install```. This is mandatory because other wise exceptions will be thrown. 
4. Make sure you have started the ```Production server``` located in the [production](https://github.com/Krijnrien/DataWall/tree/Production) make sure you follow the instructions in the production branch 
5. Make sure you have a testing machine (another computer besides the server)
6. Local Area Network connected by a router to the external network is requiered fot the test to be performed. ```Note : Make sure all machines have connection with the router and have a private ip address!```. This depends on your ```router settings```. If `Problems occur` [Tutorial](https://www.pcworld.com/article/249185/networking/how-to-set-up-a-wireless-router.html)  
### GraphQL script setup
```There are 2 Client scripts ``` located in the GraphQL folder. 
1. Open them in a text editor example([notedpad++](https://notepad-plus-plus.org/download/v7.5.1.html), notedpad, [Atom](https://atom.io/), [Sublime](https://www.sublimetext.com/3) etc.)
2. Change the Ip address by default to your servers local ip address in both ```client_1.js and client_2.js``` files.
```js 
'use strict'
const mqtt = require('mqtt')
const client = mqtt.connect('mqtt://ip address : mqtt port on server')
'Example'
const client = mqtt.connect('mqtt://10.0.0.3:1885')
```
3. 
