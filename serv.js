var http = require('http');
var key = "";

var server = http.createServer(function(req, res) {
    var auth = req.headers['authorization'];
    if (!auth) {
        res.statusCode = 401;
        res.setHeader('WWW-Authenticate', 'Basic realm="Test"');
        res.end('no auth header received');
    } else if (auth == 'Basic ' + key) {
        switch (req.url) {
            case "/hello":
                res.statusCode = 200;
                res.end('Hello');
                break;
            case "/world":
                res.statusCode = 200;
                res.end('World');
                break;
            default:
                res.statusCode = 200;
                res.end('OK');
                break;
        }
    } else {
        res.statusCode = 401; // Force them to retry authentication
        res.setHeader('WWW-Authenticate', 'Basic realm="Test"');
        res.end('not authenticated');
    }
});

if (process.argv.length != 4) {
    console.log("Usage: " + process.argv[1] + " [port] [username:password]");
} else {
    port = process.argv[2]
    usr_pwd = process.argv[3]
    key = new Buffer(usr_pwd).toString('base64');
    server.listen(port, function() {
        console.log("Server Listening on http://localhost:" + port);
    });
}
