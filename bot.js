import { Telegraf, Scenes, session, Markup } from 'telegraf'
import axios from 'axios'
import 'dotenv/config'
import * as schedule from './schedule.js'
// import scheduleScene from './scenes/scheduleScene.js'

const bot = new Telegraf(process.env.TG_BOT_TOKEN);
axios.defaults.baseURL = `${process.env.DB_URL}`;

// const stage = new Scenes.Stage([scheduleScene])
// bot.use(session())
// bot.use(stage.middleware())

// Начало
bot.start(async ctx => {
  await ctx.reply('Добро пожаловать, для начала работы, вам необходимо авторизироваться. Введите команду /auth')
})

// Начало авторизации
bot.command('auth', async ctx => {
  await ctx.reply("Пока что я нахожусь в процессе разработки, поэтому каждый пользователь \"авторизирован\"")
})

// При нажатии на кнопку Назад
bot.hears('Назад', async ctx => [
  await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
    [
      ['Посмотреть расписание', 'Найти преподавателя'],
      ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
    ]
  ))
])


// Посмотреть расписание
bot.hears('Посмотреть расписание', async ctx => {
  await ctx.reply('Что именно интересует?', Markup.keyboard(
    [
      ['Расписание на завтра', 'Расписание на сегодня'],
      ['Расписание на ...', 'Назад']
    ]
  ))
})

// Выбор расписания
bot.hears(/Расписание на (.+)/, async ctx => {
  if(ctx.match[1] === '...') {
    await ctx.reply('Выбери день недели', Markup.keyboard(
      [
        ['Расписание на Понедельник', 'Расписание на Вторник'],
        ['Расписание на Среда', 'Расписание на Четверг'],
        ['Расписание на Пятница', 'Назад']
      ]
    ))
  } else {
    await schedule.getSchedule(ctx, ctx.match[1])
  }
})

// Переход к панели администрирования
bot.command('toadmin', async ctx => {
  await ctx.reply('Временно недоступно...')
})

bot.launch()