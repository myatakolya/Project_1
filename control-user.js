export default class User {
  getUserName = (userFirstName, userLastName) => {
    return (`${userFirstName}` + ' ' + `${userLastName}`)
  }
}