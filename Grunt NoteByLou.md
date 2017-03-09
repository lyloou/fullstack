

| importing external data
- `grunt.file.readJSON()`
- `grunt.file.readYAML()`

| Template
- `<%= prop.subprop %>` Expand to the value of `prop.subprop` in the config,
regardless of type. Templates like this can be used to reference not only  
String values, but also arrays and other objects.
- `<% %>` Execute arbitrary inline JavaScript code. This is useful with control
flow or looping.
https://gruntjs.com/configuring-tasks#templates

| Global Patterns.
All most people need to know is that `foo/*.js` will match all files
ending with `.js` in the `foo/` subdirectory, but `foo/**/*.js` will match
all files ending with `.js` in the `foo/` subdirectory and all of its
subdirectories.

| Specifying both a task and target like `grunt concat:foo` or `grunt concat:bar`
will process just the specified target's configuration, while run `grunt concat`
will iterate over all targets, processing each in turn.
```js
grunt.initConfig({
  concat: {
    foo: {
    },
    bar: {
    }
  }
});
```

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

| A GruntFile is comprised of the following parts:
- The `wrapper` function
- Project and task configuration
- Loading Grunt plugins and tasks
- Custom tasks


| Because `<% %>` template strings may reference any config properties,
configuration data like filepaths and file lists may be specified this way
to reduce repetition.
- https://gruntjs.com/getting-started
