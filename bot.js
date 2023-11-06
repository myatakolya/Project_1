import TelegramBot from 'node-telegram-bot-api'
import 'dotenv/config'

const bot = new TelegramBot(process.env.TG_API_TOKEN, { polling: true });

bot.on('text', msg => {
  console.log(msg)
});

bot.on('text', msg => {
  if (msg.text == "home") {
    bot.sendMessage(msg.chat.id, "Motherfucker")
  }
});

bot.onText(/\/home/, msg => {
  bot.sendMessage(msg.chat.id, `Hi ${msg.from.first_name}`)
})