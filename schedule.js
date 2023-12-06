import axios from "axios";
const pathToUsers = '/users'
const pathToSchedule = '/schedule'

const days = [
  'Понедельник',
  'Вторник',
  'Среда',
  'Четверг',
  'Пятница',
  'Суббота',
  'Воскресенье'
]
const d = new Date();
const getCurDay = d.getDay();
const getHours = d.getHours();
const getMinutes = d.getMinutes();
const getCurTime = `${getHours}:${getMinutes}`;

export async function getSchedule(ctx, userDay = "Понедельник", userGroup = 231) {
  try{
    await axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then(async response => {
        if(userDay.toLowerCase() === 'сегодня') userDay = days[getCurDay - 1]
        else if (userDay.toLowerCase() === 'завтра') userDay = days[getCurDay]

        let result = await response.data.find( item => {
          return item.Day.toLowerCase() === userDay.toLowerCase() && item.Group  === userGroup 
        });
        let lessons = [];
        
        if (result) {
          for(let i in result.Lessons) {
            lessons.push(result.Lessons[i]) 
          }

          let lessonInfo = lessons.map(lessonInfo => {
            let message;
            if(lessonInfo.Subject) {
              message =`<b>Пара</b>: ${lessonInfo.Subject}\n` + 
              `<b>Преподаватель</b>: ${lessonInfo.Teacher}\n` + 
              `<b>Аудитория</b>: ${lessonInfo.Room}\n` + 
              `<b>Тип</b>: ${lessonInfo.Type}\n` +
              `<b>Время</b>: ${lessonInfo.Time}` 
            }
            return message
          });

          ctx.replyWithHTML(lessonInfo.join('\n\n'))
        }else if(userDay === "Суббота" || userDay === "Воскресенье") {
          await ctx.reply("Это выходной")
        } else {
          await ctx.reply("Такого дня недели не существует")
        }
      }
    ).catch(err => {
        console.error(err)
    })
  }catch(err){
    console.error(err)
  }
}

export async function getNextLesson(ctx, curDay = days[getCurDay - 1], userGroup = 231) {
  try {
    axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then(async response => {
      let result = await response.data.find( item => {
        return item.Day === curDay && item.Group === userGroup
      })
      let lessons = [];

      if(result) {
        for(let i in result.Lessons){
          lessons.push(result.Lessons[i])
        }

        let lessonInfo = lessons.map(lessonInfo => {
          let message;
          if(lessonInfo.Subject) {
            message = `<b>Пара</b>: ${lessonInfo.Subject}\n` + 
            `<b>Преподаватель</b>: ${lessonInfo.Teacher}\n` + 
            `<b>Аудитория</b>: ${lessonInfo.Room}\n` + 
            `<b>Тип</b>: ${lessonInfo.Type}\n` +
            `<b>Время</b>: ${lessonInfo.Time}` 
          }
          return message
        })

        if(getCurTime >= '8:00' && getCurTime < '9:50') {
          await ctx.replyWithHTML(lessonInfo[0])
        } else if (getCurTime >= '9:50' && getCurTime < '11:30') {
          await ctx.replyWithHTML(lessonInfo[1])
        } else if (getCurTime >= '11:30' && getCurTime < '13:20') {
          await ctx.replyWithHTML(lessonInfo[2])
        } else if (getCurTime >= '13:20' && getCurTime < '15:00') {
          await ctx.replyWithHTML(lessonInfo[3])
        } else if (getCurTime >= '15:00' && getCurTime < '16:40') {
          await ctx.replyWithHTML(lessonInfo[4])
        } else if (getCurTime >= '16:40' && getCurTime < '18:20') {
          await ctx.replyWithHTML(lessonInfo[5])
        } else if (getCurTime >= '18:20' && getCurTime < '20:00') {
          await ctx.replyWithHTML(lessonInfo[6])
        } else if (getCurTime >= '20:00' && getCurTime < '21:30') {
          await ctx.replyWithHTML(lessonInfo[7])
        } else {
          await ctx.reply("Пары кончились")
        }
      } else {
        ctx.reply("Ошибка")
      }
    }).catch( err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}