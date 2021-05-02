const { EmojiRequest, EmojiResponse } = require('./emoji_pb.js');
const { EmojiServiceClient } = require('./emoji_grpc_web_pb.js');

var client = new EmojiServiceClient('https://preprod-voucherkit.gonuclei.com');

var request = new EmojiRequest();
request.setInputText("Hi :D");

client.insertEmojis(request, {}, function (err, response) {
  if (err != null) {
    alert(err.name, err.message);
    console.log("failed");
  } else {
    console.log(response.getMessage());
  }
});