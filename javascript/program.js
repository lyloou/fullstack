console.log("Hello world");

// setTimeout(function(){console.log("time now");}, 1000);

var http = require('http');

var server = new http.Server();
server.listen(8080);
server.on('request', function(request, response){
  console.log(request.url);
  // response.writeHead(200, {"Content-Type":"text/plain; charset=UTF-8"});
  response.write("Hello World2222");
  response.end();
});
