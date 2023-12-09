import { Telegraf, Scenes, session, Markup } from 'telegraf'
import axios from 'axios'
import 'dotenv/config'
import * as schedule from './funcs/schedule.js'
import * as userControl from './funcs/userControl.js'
import authScene from './scenes/auth.js'

const bot = new Telegraf(process.env.TG_BOT_TOKEN);
axios.defaults.baseURL = `${process.env.DB_URL}`;
const pathToUsers = '/users'

const stage = new Scenes.Stage([authScene])
bot.use(session())
bot.use(stage.middleware())

// Начало
bot.start(async ctx => {
  await ctx.reply('Добро пожаловать, для начала работы, вам необходимо зарегестрироваться. Введите команду /auth')
})

// Начало авторизации
bot.command('auth', async ctx => {
  await userControl.checkAuth(
    ctx,
    async () => {
      await ctx.reply('Вы уже авторизованы')
    },
    async () => {
      await ctx.scene.enter('authWizard')
    }
  )
})

// Подтверждение существования пользователя
bot.command('join', async ctx => {
  await userControl.checkAuth(
    ctx,
    async () => {
      await userControl.checkRole(
        ctx,
        async () => {
          await ctx.reply('Здравствуй хозяин...', 
            Markup.keyboard(
              [
                // ['/toadmin']
                ['Посмотреть расписание', 'Найти преподавателя'],
                ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
              ]
            )
          )
        },
        async () => {
          await ctx.reply('Поздравляю с успешной авторизацией, выберите, что вы хотите сделать...', 
            Markup.keyboard(
              [
                ['Посмотреть расписание', 'Найти преподавателя'],
                ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
              ]
            )
          )
        },
        async () => {
          await ctx.reply('Поздравляю с успешной авторизацией, выберите, что вы хотите сделать...', 
            Markup.keyboard(
              [
                ['Посмотреть расписание', 'Оставить комментарий к паре'],
                ['Узнать, где следующая пара', 'Где ученики']
              ]
            )
          )
        }
      )
      
    }
  )
})

// При нажатии на кнопку Назад
bot.hears('Назад', async ctx => {
  await userControl.checkRole(
    ctx,
    async () => {
      await ctx.reply('Здравствуй хозяин...', 
        Markup.keyboard(
          [
            // ['/toadmin']
            [
              ['Посмотреть расписание', 'Найти преподавателя'],
              ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
            ]
          ]
        )
      )
    },
    async () => {
      await ctx.reply('Поздравляю с успешной авторизацией, выберите, что вы хотите сделать...', 
        Markup.keyboard(
          [
            ['Посмотреть расписание', 'Найти преподавателя'],
            ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
          ]
        )
      )
    },
    async () => {
      await ctx.reply('Поздравляю с успешной авторизацией, выберите, что вы хотите сделать...', 
        Markup.keyboard(
          [
            ['Посмотреть расписание', 'Оставить комментарий к паре'],
            ['Узнать, где следующая пара', 'Где ученики']
          ]
        )
      )
    }
  )
})

// Узнать, где следующая пара
bot.hears('Узнать, где следующая пара', async ctx => {
  await schedule.getNextLesson(ctx);
})


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