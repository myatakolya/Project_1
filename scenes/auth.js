import { Scenes, Composer } from 'telegraf'
import axios from 'axios'
axios.defaults.baseURL = `${process.env.DB_URL}`;
const pathToUsers = '/users'
const pathToUser = '/user'

let userGit;
let userFirstName;
let userLastName;
let userGroup;

const entering = new Composer();
entering.on('text', async ctx => {
  await ctx.reply('Отправьте ссылку на свой гитхаб')
  return ctx.wizard.next()
})

const git = new Composer();
git.on('text', async ctx => {
  userGit = ctx.message.text;
  await ctx.reply('Отлично, введите ваше имя')
  return ctx.wizard.next()
})

const firstName = new Composer();
firstName.on('text', async ctx => {
  userFirstName = ctx.message.text;
  await ctx.reply('Введите вашу фамилию')
  return ctx.wizard.next()
})

const lastName = new Composer();
lastName.on('text', async ctx => {
  userLastName = ctx.message.text;
  await ctx.reply('Введите вашу группу(Цифрами, без подгруппы)')
  return ctx.wizard.next()
})

const group = new Composer();
group.on('text', async ctx => {
  userGroup = +ctx.message.text;
  await ctx.reply('Почти всё готово, добавляю ваши данные в базу')
  await axios({
    method: 'POST',
    url: `${pathToUser}`,
    data: {
      // Git: userGit,
      Tg_id: `${ctx.message.from.id}`,
      Name: userFirstName,
      Surname: userLastName,
      Username: ctx.message.from.username,
      Groupe: +userGroup,
      Role: 'Ученик'
    },
    headers: {
      "Content-type": "application/json; charset=UTF-8"
    }
  })
  await ctx.reply('Регистрация завершена успешно, введите команду /join')
  await ctx.scene.leave()
})
const authScene = new Scenes.WizardScene('authWizard', entering, git, firstName, lastName, group)

export default authScene