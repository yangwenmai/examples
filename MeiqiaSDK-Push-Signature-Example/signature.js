// ECMAScript 5
// the method to read raw body in Express
// app.configure(function() {
// 	//. . .
// 	app.use(express.bodyParser());
// 	app.use(function(req, res, next) {
// 	        var data = '';
// 	        req.setEncoding('utf8');
// 	        req.on('data', function(chunk) { 
// 	            data += chunk;
// 	        });
// 	        req.on('end', function() {
// 	            req.rawBody = data;
// 	        });
// 	        next();
// 	    });
	
// 	//. . .
// });
// 
// Example:
// if(DTsiner(req.rawBody) != req.getHeaders('Authorization')){
//
// }
var crypto = require('crypto')

function DTSigner(raw_body, key){
	var hash_str = crypto.createHmac('sha1', key).update(raw_body).digest('hex');
	var sign_str = new Buffer(hash_str).toString('base64')
	.replace(/\+/g, '-')
    .replace(/\//g, '_');
	return 'meiqia_sign:' + sign_str;
}
