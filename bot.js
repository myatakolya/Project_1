import TelegramBot from 'node-telegram-bot-api'
// import express from 'express'
import 'dotenv/config'

// const app = express();
const bot = new TelegramBot(process.env.TG_API_TOKEN, { polling: true });

bot.on('text', msg => {
  console.log(msg)
});

bot.on('text', msg => {
  bot.sendMessage(msg.chat.id, `Hi, ${msg.from.first_name}`)
});

bot.on('text', async msg => {

  const msgWait = await bot.sendMessage(msg.chat.id, 'Wait 5 sec, pls');

  setTimeout( async () => {
    bot.editMessageText('Yappy', {
      chat_id: msgWait.chat.id,
      message_id: msgWait.message_id
    });

  }, 5000);
});


// app.listen(process.env.PORT, () => console.log(`My server is running on port ${process.env.PORT}`))
// process.once("SIGINT", () => bot.stop("SIGINT"));
// process.once("SIGTERM", () => bot.stop("SIGTERM"));