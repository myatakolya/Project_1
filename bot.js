import { Scenes, session, Telegraf } from 'telegraf'
import { message } from 'telegraf/filters'
import 'dotenv/config'
import axios from 'axios'
// import debug from './helpers.js'

const bot = new Telegraf(process.env.TG_API_TOKEN);
const db = process.env.DB_URL
const db_users = process.env.DB_USERS
const db_schedule = process.env.DB_SCHEDULE
const db_roles = process.env.DB_ROLES

bot.use(session())

const question = new Scenes.BaseScene("question");
question.enter(ctx => ctx.reply('Enter name'));
question.on(message('text'), async ctx => {
  const userName = ctx.message.text;
  await ctx.reply(`Your name is ${userName}`);
  return ctx.scene.leave();
})

bot.start(ctx => {
  console.log(ctx)
})
bot.launch()


// async function getAllUsers() {
//   try {
//     await axios.get(process.env.DB_URL + process.env.DB_USERS + `?_limit=2`).then(response => {
//       console.log(response)
//     })
//   } catch (err) {
//     console.log(err)
//   }
// }