module.exports = function(grunt) {
	grunt.initConfig({
			log: {
				foo: [1, 2, 3],
				bar: 'Hello WWorld',
				baz: false
			}
	});


	grunt.registerMultiTask('log', 'log some stuff', function() {
		grunt.log.write(this.target + ":" + this.data + " " + this);
		// return false;
	});

	grunt.registerTask('log2', 'log some stuff', function() {
		// grunt.log.write(this.target + ":" + this.data + " ");
		return false;
	});

	grunt.registerTask('foo', 'A sample task that logs stuff.', function(arg1, arg2){
		if(arguments.length === 0) {
			grunt.log.write(this.name + ", no args.");
		} else {
			grunt.log.write(this.name + ", arg1=" + arg1 + ", arg2=" + arg2);
		}
	});
};
