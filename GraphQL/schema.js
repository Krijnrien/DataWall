'use strict'

const db = require('./db')
const axios = require('axios')

const {
  graphql,
  GraphQLSchema,
  GraphQLObjectType,
  GraphQLList,
  GraphQLString,
  GraphQLNonNull,
  GraphQLFloat,
  GraphQLInt
} = require('graphql')

const location = new GraphQLObjectType({
  name: 'Location',
  fields: () => ({
    x: { type: GraphQLFloat },
    y: { type: GraphQLFloat },
    z: { type: GraphQLFloat },
    userType: { type: GraphQLInt },
    hash: { type: GraphQLString },
  })
})

const userType = new GraphQLObjectType({
  name: 'UserType',
  fields: () => ({
    id: {type: GraphQLString},
    firstname: {type: GraphQLString},
    lastname: {type: GraphQLString},
    loc: {
      type: new GraphQLList(location),
      resolve (parent, args, context, ast) {
        return axios.get('http://localhost:3000/array').then(res => res.data);
      }
    },
    createdAt: {type: GraphQLString},
    updatedAt: {type: GraphQLString}
  })
})

const schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: 'QueryType',
    fields: {
      hello: {
        type: GraphQLString,
        resolve (parent, args, context, ast) {
          return 'world'
        }
      },
      user: {
        type: userType,
        args: {
          id: {type: new GraphQLNonNull(GraphQLString)}
        },
        resolve (parent, args, context, ast) {
          const userId = args.id
          const user = db.users[userId]
          if (!user) throw new Error('Not Found')
          else return user
        }
      }
    }
  }),
  mutation: new GraphQLObjectType({
    name: 'MutationType',
    fields: {
      updateUser: {
        type: userType,
        args: {
          id: {type: new GraphQLNonNull(GraphQLString)},
          firstname: {type: GraphQLString},
          lastname: {type: GraphQLString}
        },
        resolve (_, args, context, ast) {
          const mqttServer = context.mqttServer
          const userId = args.id
          const user = db.users[userId]
          if (!user) {
            throw new Error('Not Found')
          }

          // SAVE IN THE DATABASE
          user.firstname = args.firstname ? args.firstname : user.firstname
          user.lastname = args.lastname ? args.lastname : user.lastname
          user.updatedAt = `${Date.now()}`
          console.log(Date.now());
          mqttServer.publish(`subscribeUserUpdate_${userId}`, user)
          .catch((err) => console.error(err))

          return user
        }
      }
    }
  }),
  subscription: new GraphQLObjectType({
    name: 'SubscriptionType',
    fields: {
      subscribeUserUpdate: {
        type: userType,
        args: {
          id: {type: new GraphQLNonNull(GraphQLString)}
        },
        resolve (user) {
          return user // user
        }
      }
    }
  })
})

module.exports = schema
