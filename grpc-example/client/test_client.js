var PROTO_PATH = __dirname + '/../proto/test/test.proto';

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

var test_proto = grpc.loadPackageDefinition(packageDefinition).test;

function main() {
  var client = new test_proto.Web('localhost:50051', grpc.credentials.createInsecure());

  client.register({name: "Jay", age: 30}, function(err, response) {
    console.log(response.message);
  });

  client.check({name: "JIWON"}, function(err, response) {
    console.log(response.message);
  });
}

main();