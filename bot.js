import TelegramBot from 'node-telegram-bot-api'
// import mongoose from 'mongoose'
import 'dotenv/config'
import debug from './helpers.js'

// mongoose.connect(process.env.DB_URL, {
//   useMongoClient: true
// }).then( () => {
//   console.log("MongoDB connected")
// }).catch( (err) => {
//   console.log(err)
// })

const bot = new TelegramBot(process.env.TG_API_TOKEN, { polling: true });

bot.on('message', msg => {
  console.log(msg)

});
bot.on('message', msg => {
  const chatId = msg.chat.id;
  bot.sendMessage(chatId, 'some text', {
    reply_markup: {
      inline_keyboard: [
        [{
          text: 'button',
          callback_data: '1'
        }]
      ]
    }
  })
});
