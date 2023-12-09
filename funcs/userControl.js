import axios from 'axios'

axios.defaults.baseURL = `${process.env.DB_URL}`;
const pathToUsers = '/users'

export async function checkAuth(ctx, ifTrue, ifElse) {
  await axios({
    method: 'get',
    url: `${pathToUsers}`
  }).then(async response => {
    const result = await response.data.find( item => {
      return item.Tg_id === `${ctx.message.from.id}`
    })

    if(result) {
      await ifTrue()
    } else {
      await ifElse(ifElse ?? ctx.reply("Сначала зарегистрируйся")) 
    }
  }).catch(err => {
    console.error(err)
  })
}

export async function checkRole(ctx, ifAdmin, ifStudent, ifTeacher) {
  await axios({
    method: 'get',
    url: `${pathToUsers}`
  }).then(async response => {
    const result = await response.data.find( item => {
      return item.Tg_id === `${ctx.message.from.id}`
    })

    if(result) {
      if(result.Role === 'Админ'){
        ifAdmin()
      } else if (result.Role === 'Ученик') {
        ifStudent()
      } else if (result.Role === 'Учитель') {
        ifTeacher()
      }
    } else {
      ctx.reply("Сначала зарегистрируйся")
    }
  }).catch(err => {
    console.error(err)
  })
}