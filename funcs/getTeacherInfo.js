import axios from "axios"
const pathToUsers = '/users';

export async function getTeacherName(ctx) {
  try {
    return axios({
      method: 'get',
      url: `${pathToUsers}`
    }).then( async response => {
      const result = await response.data.find( item => {
        return item.Tg_id === `${ctx.message.from.id}` && item.Role === "Учитель"
      })

      if(result){
        return result.Name
      }
    }).catch(err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}
export async function getTeacherSurname(ctx) {
  try {
    return axios({
      method: 'get',
      url: `${pathToUsers}`
    }).then( async response => {
      const result = await response.data.find( item => {
        return item.Tg_id === `${ctx.message.from.id}` && item.Role === "Учитель"
      })

      if(result){
        return result.Surname
      }
    }).catch(err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}