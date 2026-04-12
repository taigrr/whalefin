const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');

let sourceDir = path.resolve(__dirname, 'src');
let buildDir = path.resolve(__dirname, 'build');

module.exports = {
	entry: {
		index: path.resolve(sourceDir, 'main.js')
	},
	output: {
		path: buildDir,
		filename: 'main.js'
	},
	optimization: {
		splitChunks: false
	},
	devServer: {
		disableHostCheck: true,
		contentBase: path.join(__dirname, 'src'),
		compress: true,
		open: true,
		port: 8090
	},
	mode: 'production',
	plugins: [
		new CopyWebpackPlugin({
			patterns: [
				{
					from: path.resolve(sourceDir, 'main.css'),
					to: path.resolve(buildDir, 'main.css')
				},
				{
					from: path.resolve(sourceDir, 'index.html'),
					to: path.resolve(buildDir, 'index.html')
				},
			]
		})
	]
};
