import axios from "axios"

const pathToUsers = '/users'

export async function getUserGroup(ctx) {
  try {
    return axios({
      method: 'get',
      url: `${pathToUsers}`
    }).then( async response => {
      const result = await response.data.find( item => {
        return item.Tg_id === `${ctx.message.from.id}`
      })

      if(result) {
        return result.Group
      }
    }).catch(err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}
export async function getUserSubgroup(ctx) {
  try {
    return axios({
      method: 'get',
      url: `${pathToUsers}`
    }).then( async response => {
      const result = await response.data.find( item => {
        return item.Tg_id === `${ctx.message.from.id}`
      })

      if(result) {
        return result.Subgroup
      }
    }).catch(err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}