import { Telegraf, Scenes, session, Composer } from 'telegraf'
import { message } from 'telegraf/filters'
import 'dotenv/config'
import axios from 'axios'
// import debug from './helpers.js'

const bot = new Telegraf(process.env.TG_API_TOKEN);
const db = process.env.DB_URL
const db_users = process.env.DB_USERS
const db_schedule = process.env.DB_SCHEDULE
const db_roles = process.env.DB_ROLES
const date = new Date();

let userFirstName;
let userLastName;
let userGitID;


const startWizard = new Composer();
startWizard.on('text', async ctx => {
  await ctx.reply('Введите ваше имя')
  return ctx.wizard.next()
})

const firstName = new Composer();
firstName.on('text', async ctx => {
  if(ctx.message.text) {
    userFirstName = ctx.message.text
    await ctx.reply('Введите вашу фамилию')
    return ctx.wizard.next()
  } else { 
    await ctx.reenter()
  }
})

const lastName = new Composer();
lastName.on('text', async ctx => {
  if (ctx.message.text) {
    userLastName = ctx.message.text
    await ctx.scene.leave()
  } else {
    await ctx.reenter()
  }
})

const getUserScene = new Scenes.WizardScene('userName', startWizard, firstName, lastName);
bot.use(session())
const stage = new Scenes.Stage([getUserScene])
bot.use(stage.middleware())

bot.start(async ctx => {
  await ctx.reply('Добро пожаловать, меня зовут Расписалово. Я - бот, который управляет расписанием. Перед началом работы необходимо аутентифицироваться')
  await ctx.scene.enter('userName')
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