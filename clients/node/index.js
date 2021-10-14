const grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    __dirname + '../../../gamingstats/gaming_stats.proto',
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
const hello_proto = grpc.loadPackageDefinition(packageDefinition).gamingstats;

function main() {
    var client = new hello_proto.GamingStats('localhost:8000',
                                         grpc.credentials.createInsecure());
    client.getGame({title: 'uncharted 4'}, function(err, response) {
      if (err) {
          throw err
      }
      console.log('Game found:', response);
    });
}

main()
