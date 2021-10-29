const webpack = require("webpack");
const path = require("path");

const source = "view/app"

module.exports = {
    entry: path.resolve(__dirname, path.join(source, "index.js")),
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {loader: "babel-loader"}
            },
            {
                test: /\.html$/,
                use: [{loader: "html-loader"}]
            },
            {
                test: /\.css$/,
                use: ["style-loader", "css-loader"]
            },
            {
                test: /\.(png|jpe?g|gif)$/i,
                use: ['file-loader']
            }
        ],
    },
    resolve: {
        extensions: ["*", ".js", ".jsx"],
    },
    output: {
        path: path.resolve(__dirname, "./dist"),
        filename: "bundle.js",
    },
    plugins: [new webpack.HotModuleReplacementPlugin()],
    devServer: {
        contentBase: path.resolve(__dirname, "./dist"),
        hot: true,
    },
};
