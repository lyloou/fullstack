

| Define a configuration for each of the tasks we mentioned.
Note: The configuration object for a plugins lives as a property on the
  configuration object, that offen shares the same name as its plugin.
  eg. `grunt-contrib-concat` => `concat`

| 5个插件
- grunt-contrib-watch: check the code at every change you perform；
- grunt-contrib-jshint: ensure best practices, or check the code behaviors；
- grunt-contrib-uglify: 创建一个minified的version；
- grunt-contrib-qunit: 测试你的project；
- grunt-contrib-concat: 多个文件合并成一个文件；


| The typical folder structure features the following folders: src, dist, and test.
- The src folder (sometimes called app) contains the source code of the library as you author it.
- The dist folder (sometimes called build) contains the distribution, a minified version of the source code.
- the test folder contains the code to test the project.
https://gruntjs.com/sample-gruntfile

| Grunt and Grunt plugins should be defined as devDependencies in your project's
`package.json`.
please use: `npm install $package --save-dev`
https://gruntjs.com/installing-grunt


| load all grunt tasks
```js
// load all grunt tasks
require('matchdep').filterDev('grunt-*').forEach(grunt.loadNpmTasks);

// 上面的一句话就替代了很多次的类似下面的调用
grunt.loadNpmTasks('grunt-contrib-jshint');
grunt.loadNpmTasks('grunt-contrib-watch');
```
注意：`matchdep`需要在`package.json`中的`devDependencies`中添加：
`"matchdep": "latest"`


| Because `<% %>` template strings may reference any config properties,
configuration data like filepaths and file lists may be specified this way
to reduce repetition.
- https://gruntjs.com/getting-started
