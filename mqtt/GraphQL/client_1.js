'use strict'

const mqtt = require('mqtt')
const client = mqtt.connect('mqtt://localhost:1885')

const sub1 = 'subscription{subscribeUserUpdate(id:"abcd"){loc{x}}}'
const topic1 = `subscribeUserUpdate_abcd`

client.on('message', (topic, message) => {
  console.log(message.toString())
  console.log(Date.now());
})

client.on('connect', (connack) => {
  console.log('connect')
})

client.subscribe(`${topic1}__${sub1}`, (err, granted) => {
  if (err) console.error('Error subscribing...')
})
