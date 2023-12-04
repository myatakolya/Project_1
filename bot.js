import { Telegraf, Scenes, session, Composer } from 'telegraf'
import { message } from 'telegraf/filters'
import 'dotenv/config'
import axios from 'axios'
import User from './control-user.js'
// import debug from './helpers.js'

const bot = new Telegraf(process.env.TG_API_TOKEN);
const pathToUsers = '/users'
const pathToSchedule = '/schedule'

axios.defaults.baseURL = `${process.env.DB_URL}`

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
    await ctx.reply(`Приветствую, ${userFirstName} ${userLastName}`)
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

bot.command('toadmin' , ctx => {
  ctx.reply('Данная функция пока недоступна')
})

bot.command('whoami', ctx => {
  const user = getUserName();
  ctx.reply(user)
})

bot.hears(/Расписание на (.*)/, async (ctx) => {
  let curDay = ctx.match[1];
  console.log(curDay)
  await getSchedule(ctx, curDay)
})

bot.command('showusers', async ctx => {
  console.log(getUserById('653ebb213c4c46bb0101cd7a'))
})
bot.launch()

async function getUserById(userId) {
  try {
    await axios({
      method: 'get',
      url: `${pathToUsers}`
    }).then(response => {
      console.log(response.data.find( item => {
        return item.ID === `${userId}`;
      }))
    }).catch(err => {
      console.log(err)
    });
  } catch( err ) {
    console.log(err)
  }
}

async function getSchedule(ctx, curDay = "Понедельник", userGroup = "231") {
  try{
    await axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then(
      async response => {
        let result = await response.data.find( item => { return item.Day === curDay && item.Group === +userGroup });
        if (result) {
          await ctx.replyWithHTML(`
<b>Первая пара:</b> ${result.Lessons.Lesson1.Subject} 
<b>Преподаватель:</b> ${result.Lessons.Lesson1.Teacher} 
<b>Тип:</b> ${result.Lessons.Lesson1.Type} 
<b>Аудитория:</b> ${result.Lessons.Lesson1.Room} 
<b>Время:</b> ${result.Lessons.Lesson1.Time} 
          `)
          console.log(result)
        }else {
          await ctx.reply("Такого дня недели не существует")
        }
      }
    ).catch(err => {
        console.log(err)
    })
  }catch(err){
    console.log(err)
  }
}

async function getUserName() {
  try {
    return userFirstName + userLastName
  } catch ( err ) {
    console.log(err)
  }
}