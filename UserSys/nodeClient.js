var PROTO_PATH = __dirname + '/proto/user.proto';

var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });

var proto = grpc.loadPackageDefinition(packageDefinition).grpc_user;
var client = new proto.UserService('localhost:8080', grpc.credentials.createInsecure());

function main() {

    client.register({name: "jiwon", email: "jiohning00@hanmail.net", password: "123456"}, function(err, response) {
        console.log("~Register~");
        console.log(response.response);
    });

    client.login({email: "jiohning00@hanmail.net", password: "12345"}, function(err, response) {
        console.log("~Login~");
        console.log(response.response);
    });

}

main();