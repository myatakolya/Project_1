import { Telegraf, Scenes, session, Markup } from 'telegraf'
import axios from 'axios'
import 'dotenv/config'
import * as schedule from './funcs/schedule.js'
import * as permission from './funcs/checkPermission.js'
import * as teacherSchedule from './funcs/teacherSchedule.js'
import authScene from './scenes/auth.js'

const bot = new Telegraf(process.env.TG_BOT_TOKEN);
axios.defaults.baseURL = `${process.env.DB_URL}`;

const stage = new Scenes.Stage([authScene])
bot.use(session())
bot.use(stage.middleware())

bot.start(async ctx => {
  await ctx.reply('Добро пожаловать, для начала работы, вам необходимо зарегистрироваться. Введите команду /auth')
})

bot.command('auth', async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await ctx.reply('Вы уже авторизованы, введите команду /join')
    },
    async () => {
      await ctx.scene.enter('authWizard')
    }
  )
})

bot.command('join', async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await ctx.reply('Здравствуй хозяин...', Markup.keyboard(
              [
                ["Хочу быть учеником", "Хочу быть преподом"],
                ["/toadmin"]
              ]
            )
          )
        },
        async () => {
          await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
              [
                ['Посмотреть расписание', 'Найти преподавателя'],
                ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
              ]
            )
          )
        },
        async () => {
          await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
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
bot.hears('Назад', async ctx => {
  await permission.checkRole(
    ctx,
    async () => {
      await ctx.reply('Да, хозяин, кем ты хочешь быть?', Markup.keyboard(
          [
            ["Хочу быть учеником", "Хочу быть преподом"],
            ["/toadmin"]
          ]
        )
      )
    },
    async () => {
      await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
          [
            ['Посмотреть расписание', 'Найти преподавателя'],
            ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН']
          ]
        )
      )
    },
    async () => {
      await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
          [
            ['Посмотреть расписание', 'Оставить комментарий к паре'],
            ['Узнать, где следующая пара', 'Где ученики']
          ]
        )
      )
    }
  )
})
bot.hears('Узнать, где следующая пара', async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await schedule.getNextLesson(ctx)
        },
        async () => {
          await schedule.getNextLesson(ctx)
        },
        async () => {
          await teacherSchedule.getTeacherNextLesson(ctx)
        }
      )
    }
  )
})
bot.hears("Где ученики", async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await ctx.reply("Какая группа вас интересует?", Markup.keyboard(
            [
              ["Группа ИВТ-б-о-231"],
              ["Назад"]
            ]
          ))
        },
        async () => {
          await ctx.reply("Недостаточно прав")
        },
        async () => {
          await ctx.reply("Какая группа вас интересует?", Markup.keyboard(
            [
              ["Группа ИВТ-б-о-231", "Группа ИВТ-б-о-232"],
              ["Назад"]
            ]
          ))
        },
      )
    }
  )
})
bot.hears(/Группа (.+)/, async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await teacherSchedule.getStudents(ctx, ctx.match[1])
        },
        async () => {
          await ctx.reply("Недостаточно прав")
        },
        async () => {
          await teacherSchedule.getStudents(ctx, ctx.match[1])
        }
      )
    }
  )
})

bot.hears('Посмотреть расписание', async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await ctx.reply('Что именно интересует?', Markup.keyboard(
        [
          ['Расписание на завтра', 'Расписание на сегодня'],
          ['Расписание на ...', 'Назад']
        ]
      ))
    }
  )
})

bot.hears(/Расписание на (.+)/, async ctx => {
  await permission.checkAuth(
    ctx,
    async () => { 
      if(ctx.match[1] === '...') {
        await ctx.reply('Выбери день недели', Markup.keyboard(
          [
            ['Расписание на Понедельник', 'Расписание на Вторник'],
            ['Расписание на Среда', 'Расписание на Четверг'],
            ['Расписание на Пятница', 'Назад']
          ]
        ))
    } else {
      await permission.checkRole(
        ctx,
        async () => {
          await schedule.getSchedule(ctx, ctx.match[1])
        },
        async () => {
          await schedule.getSchedule(ctx, ctx.match[1])
        },
        async () => {
          await teacherSchedule.getTeacherSchedule(ctx, ctx.match[1])
        }
      )
    }}
  )
})
bot.hears("Узнать, когда БУЭКЗАМЕН", async ctx =>{
  await permission.checkAuth(
    ctx,
    async () => {
      await ctx.reply("Выберите предмет...", Markup.keyboard(
        [
          ["Экзамен по Алгоритмизация и Программирование"],
          ["Назад"]
        ]
      ))
    }
  )
})
bot.hears(/Экзамен по (.+)/, async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await schedule.getExam(
        ctx,
        ctx.match[1]
      )
    }
  )
})
bot.hears(/Хочу быть (.+)/, async ctx => {
  if(ctx.match[1] === "учеником") {
    await ctx.reply("Теперь ты ученик", Markup.keyboard(
      [
        ['Посмотреть расписание', 'Найти преподавателя'],
        ['Узнать, где следующая пара', 'Узнать, когда БУЭКЗАМЕН'],
        ["Назад"]
      ]
    ))
  }  else if (ctx.match[1] === "преподом") {
    await ctx.reply('Выберите, что вы хотите сделать...', Markup.keyboard(
      [
        ['Посмотреть расписание', 'Оставить комментарий к паре'],
        ['Узнать, где следующая пара', 'Где ученики'],
        ["Назад"]
      ]
    ))
  }
})
bot.hears("Найти преподавателя", async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await ctx.reply("Кого вы хотите найти?", Markup.keyboard(
            [
              ['Найти Владимир Чабанов'],
              ["Назад"]
            ]
          ))
        },
        async () => {
          await ctx.reply("Кого вы хотите найти?", Markup.keyboard(
            [
              ['Найти Владимир Чабанов'],
              ["Назад"]
            ]
          ))
        },
        async () => {
          await ctx.reply("Так вы и есть препродаватель...")
        }
      )
    }
  )
})
bot.hears(/Найти (.*) (.*)/, async ctx => {
  await permission.checkAuth(
    ctx,
    async () => {
      await permission.checkRole(
        ctx,
        async () => {
          await teacherSchedule.getTeacherNextLesson(ctx, ctx.match[1], ctx.match[2])
        },
        async () => {
          await teacherSchedule.getTeacherNextLesson(ctx, ctx.match[1], ctx.match[2])
        },
        async () => {
          await ctx.reply("Можно только ученикам")
        },
      )
    }
  )
})

bot.command('toadmin', async ctx => {
  await ctx.reply('Временно недоступно...') 
})

bot.launch()