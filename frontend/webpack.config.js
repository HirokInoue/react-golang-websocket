const path = require('path');
const NODE_ENV = process.env.NODE_ENV || "development";

module.exports = {
    mode: NODE_ENV,

    entry: './src/index.tsx',

    output: {
        path: path.resolve(__dirname, "./public"),
        filename: "index.js",
        publicPath: "/"
    },

    module: {
        rules: [{
            test: /\.tsx?$/,
            use: 'ts-loader',
            exclude: /node_modules/
        }]
    },
    resolve: {
        modules: [
            "node_modules",
        ],
        extensions: [
            '.ts',
            '.tsx',
            '.js',
            '.jsx'
        ]
    },
    devServer: {
        static: {
            directory: path.resolve(__dirname, "./public"),
        }
    }
};