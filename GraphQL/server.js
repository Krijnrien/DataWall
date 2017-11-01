
const express = require('express');
const expressGraphQL = require('express-graphql');
const schema = require('./schema.js');
const mqttServer = require('./mqtt')

const app = express();

app.use('/graphql', expressGraphQL({
  schema: schema,
  graphiql: true,
  context: { mqttServer }
}));


mqttServer.listen(1885, () => console.log('MQTT is listening on port 1885!'));
app.listen(4000, () => {
  console.log('starting GraphQL');
  console.log('server is running on port 4000');
});
