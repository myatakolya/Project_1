// import { Composer, Scenes, Markup, Telegraf } from 'telegraf'
// import axios from 'axios'
// import * as schedule from '../schedule.js'

// const bot = new Telegraf(process.env.TG_BOT_TOKEN)

// const entering = new Composer();
// entering.on('text', async ctx => {
//   try {
    

//     return ctx.wizard.next()
//   } catch (err) {
//     console.error(err)
//   }
// })

// const chose = new Composer();
// chose.hears(/Расписание на (.+)/, async ctx => {
//   if(ctx.match[1] === '...') {
//     await ctx.reply('Выбери день недели', Markup.keyboard(
//       [
//         ['Расписание на Понедельник', 'Расписание на Вторник'],
//         ['Расписание на Среда', 'Расписание на Четверг'],
//         ['Расписание на Пятница']
//       ]
//     ))
//   } else {
//     await schedule.getSchedule(ctx, ctx.match[1])
//   }
// })

// const scheduleScene = new Scenes.WizardScene('scheduleWizard', entering, chose)

// export default scheduleScene