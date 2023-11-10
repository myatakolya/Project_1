import { Scenes, session, Telegraf, Composer } from 'telegraf'
import { message } from 'telegraf/filters'
import 'dotenv/config'
import axios from 'axios'
// import debug from './helpers.js'

const bot = new Telegraf(process.env.TG_API_TOKEN);
// const {startWizard, firstName, lastName} = new Composer()
const db = process.env.DB_URL
const db_users = process.env.DB_USERS
const db_schedule = process.env.DB_SCHEDULE
const db_roles = process.env.DB_ROLES

const startWizard = new Composer();
startWizard.on('text', async ctx => {
  ctx.wizard.state.data = {};
  await ctx.reply('Введите своё имя')
  return ctx.wizard.next()
})

const firstName = new Composer();
firstName.on('text', async ctx => {
  ctx.wizard.state.data.firstName = ctx.message.text;
  await ctx.reply('Введите сою Фамилию')
  return ctx.wizard.next()
})

const lastName = new Composer();
lastName.on('text', async ctx => {
  ctx.wizard.state.data.lastName = ctx.message.text;
  await ctx.reply(`Молодец, тебя зовут ` + ctx.wizard.state.data.firstName + ' ' + ctx.wizard.state.data.lastName)
  return ctx.scene.leave()
})

const newScene = new Scenes.WizardScene('sceneWizard', startWizard, firstName, lastName);
const stage = new Scenes.Stage([newScene])
bot.use(session())
bot.use(stage.middleware())

bot.start(ctx => {
  ctx.scene.enter('sceneWizard')
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