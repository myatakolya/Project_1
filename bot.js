import { Telegraf, Scenes, session, Composer } from 'telegraf'
import { message } from 'telegraf/filters'
import 'dotenv/config'
import axios from 'axios'
import methods from 'methods';
// import debug from './helpers.js'

const bot = new Telegraf(process.env.TG_API_TOKEN);
const pathToUsers = '/users'
const pathToSchedule = '/shedule'

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
// bot.hears('Расписание на сегодня', async ctx => {
//   await getSchedule(ctx)
// })

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

// async function getSchedule(ctx) {
//   try{
//     await axios.get(
//       pathToSchedule
//     ).then( response => {
//         ctx.reply(response.data.find(item => {
//           JSON.stringify(item.Lessons, " ", 2)
//         }))
//       }
//     ).catch(err => {
//         console.log(err)
//     })
//   }catch(err){
//     console.log(err)
//   }
// }