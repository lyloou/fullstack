var http = require('http');

var server = new http.Server();
server.listen(8080);
server.on('request', function(request, response){
  console.log(request.url); //打印请求网址
  response.write("Hello World!");
  response.end();
});
