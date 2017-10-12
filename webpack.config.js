var HtmlWebpackPlugin = require('html-webpack-plugin');
var CopyWebpackPlugin = require('copy-webpack-plugin');
var ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
  entry: './src/app.js',
  output: {
    path: './dist',
    filename: 'bundle.js'
  },
  devServer: {
    inline: true,
    port: 8888,
  },
  module: {
    loaders: [{
      test: /\.js$/,
      exclude: /node_modules/,
      loader: 'babel-loader',
      query: {
        presets: ['es2015']
      }
    }, {
      test: /\.(scss|sass)$/,
      loader: ExtractTextPlugin.extract(
        'style', // The backup style loader
        'css?sourceMap!sass?sourceMap'
      )
    }]
  },
  resolve: {
    extensions: ['', '.js', '.jpg', '.png', '.gif', '.ttf']
  },
  plugins: [
    new CopyWebpackPlugin([{
      from: './src/assets',
      to: './assets'
    }]),
    new HtmlWebpackPlugin({
      title: 'p5-es6-webpack-seed',
      template: './src/index.html',
      favicon: './src/assets/img/favicon.png',
      inject: 'head'
    }),
    new ExtractTextPlugin("assets/styles.css"),
  ]
}
