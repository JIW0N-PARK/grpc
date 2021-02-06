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

    // client.register({name: "yoonha", age: 23}, function(err, response) {
    //     console.log(response.res);
    // });

    client.search({name: "jia", age: 23}, function(err, response) {
        console.log(response.res);
    });
}

main();