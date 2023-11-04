const { Telegraf } = require('telegraf');
const { message } = require('telegraf/filters')
const express = require('express');

// express
const app = express();
const PORT = 3000; 

// telegraf
const bot = new Telegraf("6891459485:AAHuNv-37ZfPC9TDJFxZ_LHPLUCKxcm9yDE");

app.get('/', (req, res) => {
  res.send("Hello World")
})

app.listen(PORT, () => console.log(`My server is running on port ${PORT}`))

bot.start(ctx => {
  return ctx.reply(`Привет ${ctx.update.message.from.first_name}! Меня зовут Расписалово, я предназначен для работы с учебным расписанием (удивительно, не так ли?!)`);
});
bot.hears("show schedule", ctx => { return })
bot.launch();

process.once("SIGINT", () => bot.stop("SIGINT"));
process.once("SIGTERM", () => bot.stop("SIGTERM"));