import TelegramBot from 'node-telegram-bot-api'
import 'dotenv/config'

const bot = new TelegramBot(process.env.TG_API_TOKEN, { polling: true });

bot.on('text', msg => {
  console.log(msg)
});

bot.on('text', msg => {
  const chatId = msg.chat.id;
  bot.sendMessage(chatId, 'some text', {
    reply_markup: {
      keyboard: [
        ['1' ,'2'],
        ['3']
      ]
    }
  })
});