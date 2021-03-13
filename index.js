var TelegramBot = require("node-telegram-bot-api");
var rabbit = require("./rabbit.min");
var token = process.env.TELEGRAM_MMCONVERTER_TOKEN;
var bot = new TelegramBot("1691273019:AAF1KVed7Ne4s0kSuDytAVdQYi3Ews7icVw", {
   polling: true,
});

var help =
   "အသုံးပြုပုံကတော့ \n" +
   "'/z မဂၤလာပါ' ဆိုရင် 'မင်္ဂလာပါ' \n" +
   "'/u မင်္ဂလာပါ' ဆိုရင် 'မဂၤလာပါ' \n" +
   "Bot source code ကတော့  https://github.com/yelinaung/telegram-mmconverter-bot မှာပါ \n" +
   "Font convert လုပ်တာ လွဲနေတယ်ဆိုရင် https://github.com/Rabbit-Converter/Rabbit/issues မှာ report လုပ်နိုင်ပါတယ်။ ";

var counter = 0;

bot.on("text", function (msg) {
   var chatId = msg.chat.id;

   // Zg to Uni
   zg2uni = /\/z */.test(msg.text);

   // Uni to Zg
   uni2zg = /\/u */.test(msg.text);

   // Help
   if (msg.text == "/help") {
      bot.sendMessage(chatId, help);
   } else if (zg2uni) {
      x = msg.text.replace(/\/z*/, "").trim();
      bot.sendMessage(chatId, rabbit.zg2uni(x));
   }
   if (uni2zg) {
      x = msg.text.replace(/\/u*/, "").trim();
      bot.sendMessage(chatId, rabbit.uni2zg(x));
   }

   counter++;
   console.log("Total request " + counter);
});
