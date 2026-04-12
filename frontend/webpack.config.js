const path = require('path');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const sourceDir = path.resolve(__dirname, 'src');
const buildDir = path.resolve(__dirname, 'build');

module.exports = {
	entry: {
		index: path.resolve(sourceDir, 'main.js')
	},
	output: {
		path: buildDir,
		filename: 'main.js',
		clean: true
	},
	optimization: {
		splitChunks: false
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
