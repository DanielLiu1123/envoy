// start a server on port 8080
var http = require('http');
var server = http.createServer(function (req, res) {
    res.writeHead(200, {'Content-Type': 'text/plain'});
    res.end('Hello World: 8080');
});
console.log('Server running at ' + 8080)
server.listen(8080);
