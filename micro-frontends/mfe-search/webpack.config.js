const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const { ModuleFederationPlugin } = require("webpack").container;
const Dotenv = require('dotenv-webpack');
const MiniCssExtractPlugin = require("mini-css-extract-plugin");

module.exports = {
  entry: "./src/index.js",
  mode: "development",
  devServer: {
    port: 3002,
    static: path.join(__dirname, "dist"),
    hot: true,
    historyApiFallback: true,
  },
  output: {
    publicPath: "auto",
    filename: "bundle.js",
    clean: true,
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        use: "babel-loader",
        exclude: /node_modules/,
      },
      {
        test: /\.css$/i,
        use: [MiniCssExtractPlugin.loader, "css-loader", "postcss-loader"],
      },
    ],
  },
  resolve: {
    extensions: [".js", ".jsx"],
  },
  plugins: [
    new ModuleFederationPlugin({
      name: "searchApp",
      filename: "remoteEntry.js",
      exposes: {
        "./Dashboard": "./src/components/Dashboard",
        "./styles.css": "./src/styles.css",
      },
      shared: {
        react: {
          singleton: true,
          requiredVersion: "^19.1.0",
        },
        "react-dom": {
          singleton: true,
          requiredVersion: "^19.1.0",
        },
        "react-router-dom": { 
          singleton: true, 
          requiredVersion: "^7.5.3" 
        },
      },
    }),
    new HtmlWebpackPlugin({
      template: "./public/index.html",
    }),
    new Dotenv({
      path: path.resolve(__dirname, '.env'),
    }),
    new MiniCssExtractPlugin({
      filename: '[name].[contenthash].css',
    }),
  ],
};
