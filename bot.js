const TelegramBot = require('node-telegram-bot-api');
const express = require('express');
require('dotenv').config();

// const app = express();
const bot = new TelegramBot(process.env.BOTKEY, {polling: true});

bot.onText(/\/start/, msg => {
  const { chat: {id} } = msg;
  bot.sendMessage(id, "Начинаем")
});


// app.listen(process.env.PORT, () => console.log(`My server is running on port ${process.env.PORT}`))
process.once("SIGINT", () => bot.stop("SIGINT"));
process.once("SIGTERM", () => bot.stop("SIGTERM"));