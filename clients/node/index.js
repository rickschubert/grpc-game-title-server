const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDefinition = protoLoader.loadSync(
    __dirname + '../../../gametitles/game_titles.proto', {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
const {gametitles} = grpc.loadPackageDefinition(packageDefinition);

const client = new gametitles.GameTitles('localhost:8000', grpc.credentials.createInsecure());
client.getGame({
    title: 'uncharted 4'
}, function(err, response) {
    if (err) {
        throw err
    }
    console.log('Game found:', response);
});
