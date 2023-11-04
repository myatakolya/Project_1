import { getMainMenu } from "./keyboards.js"

require('dotenv').config();

const { Telegraf } = require('telegraf');
const { message } = require('telegraf/filters')
const express = require('express');

const app = express();
const bot = new Telegraf(process.env.BOTKEY);

// app.get('/', (req, res) => {
//   res.send("Hello World")
// })

bot.start(ctx => {
  return ctx.reply(`Привет ${ctx.update.message.from.first_name}! Меня зовут Расписалово, я предназначен для работы с учебным расписанием (удивительно, не так ли?!)`);
});
bot.hears("show schedule", ctx => { 
  return ctx.reply("Not allow");
});
bot.on('edited_message', ctx => {
  ctx.reply('Вы успешно изменили сообщение');
});
bot.command("show_pisun", ctx => {
  ctx.reply("А ничё больше тее не показать?")
})
bot.launch();

app.listen(process.env.PORT, () => console.log(`My server is running on port ${process.env.PORT}`))

process.once("SIGINT", () => bot.stop("SIGINT"));
process.once("SIGTERM", () => bot.stop("SIGTERM"));