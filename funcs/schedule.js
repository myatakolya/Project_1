import axios from "axios"
import * as user from './getUserInfo.js'

const pathToSchedule = '/schedule'
const pathToExams = '/exams'

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
const timeList = [
  {start: {h: 8,m: 0}, end: {h: 9, m: 50}},
  {start: {h: 9,m: 50}, end: {h: 11, m: 30}},
  {start: {h: 11,m: 30}, end: {h: 13, m: 20}},
  {start: {h: 13,m: 20}, end: {h: 15, m: 0}},
  {start: {h: 15,m: 0}, end: {h: 16, m: 30}},
  {start: {h: 16,m: 30}, end: {h: 18, m: 20}},
  {start: {h: 18,m: 20}, end: {h: 20, m: 0}},
  {start: {h: 20,m: 0}, end: {h: 21, m: 30}},
];
const curTime = getHours * 60 + getMinutes;

export async function getSchedule(ctx, userDay = "Понедельник", userGroup) {
  try{
    await axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then(async response => {
      await user.getUserGroup(ctx).then( response => {
        userGroup = response;
      })

      if(userDay.toLowerCase() === 'сегодня') userDay = days.at(getCurDay - 1)
      else if (userDay.toLowerCase() === 'завтра') userDay = days.at(getCurDay)

      const result = await response.data.find( item => {
        return item.Day?.toLowerCase() === userDay?.toLowerCase() && item.Group === userGroup 
      });
      let lessons = [];
      
      if (result) {
        for(let i in result.Lessons) {
          lessons.push(result.Lessons[i]) 
        }

        let lessonInfo = lessons.map(item => {
          let message;

          if(item.Subject) {
            message =`<b>Пара</b>: ${item.Subject}\n` + 
            `<b>Преподаватель</b>: ${item.Teacher}\n` + 
            `<b>Аудитория</b>: ${item.Room}\n` + 
            `<b>Тип</b>: ${item.Type}\n` +
            `<b>Время</b>: ${item.Time}` 
          }
          
          return message
        });

        ctx.replyWithHTML(lessonInfo.join('\n\n'))
      }else if(userDay === "Суббота" || userDay === "Воскресенье") {
        await ctx.reply("Это выходной")
      } else {
        await ctx.reply("Такого дня недели не существует")
      }
    }).catch(err => {
        console.error(err)
    })
  }catch(err){
    console.error(err)
  }
}
export async function getNextLesson(ctx, curDay = days.at(getCurDay - 1), userGroup) {
  try {
    axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then(async response => {
      await user.getUserGroup(ctx).then( response => {
        userGroup = response;
      })

      const result = await response.data.find( item => {
        return item.Day === curDay && item.Group === userGroup
      })
      let lessons = [];

      if(result) {
        for(let i in result.Lessons){
          lessons.push(result.Lessons[i])
        }

        let lessonInfo = lessons.map(item => {
          let message;

          if(item.Subject) {
            message = `<b>Пара</b>: ${item.Subject}\n` + 
            `<b>Преподаватель</b>: ${item.Teacher}\n` + 
            `<b>Аудитория</b>: ${item.Room}\n` + 
            `<b>Тип</b>: ${item.Type}\n` +
            `<b>Время</b>: ${item.Time}` 
          }

          return message
        })

        let curLesson;

        for(let i = 0; i < timeList.length; i++){
          const entry = timeList[i];

          const mStart = entry.start.h * 60 + entry.start.m;
          const mEnd = entry.end.h * 60+ entry.end.m;

          if(curTime >= mStart && curTime <= mEnd) {
            curLesson = i + 1;
            break;
          }
        }

        if(lessonInfo[curLesson]){
          await ctx.replyWithHTML(lessonInfo[curLesson])
        } else if (!lessonInfo[curLesson] && lessonInfo[curLesson + 1]) {
          await ctx.reply("Окно")
        } else {
          await ctx.reply("Пары кончились")
        }
      } else {
        ctx.reply("Это выходной")
      }
    }).catch( err => {
      console.error(err)
    })
  } catch (err) {
    console.error(err)
  }
}
export async function getExam(ctx, subject,  userGroup, userSubgroup) {
  try {
    axios({
      method: 'get',
      url: `${pathToExams}`
    }).then( async response => {
      await user.getUserGroup(ctx).then( response => {
        userGroup = response;
      })
      await user.getUserSubgroup(ctx).then( response => {
        userSubgroup = response;
      })

      const result = await response.data.find(item => {
        return item.Subject.toLowerCase() === subject.toLowerCase() && item.Group === userGroup && item.Subgroup === userSubgroup
      })

      const message = `<b>Преподаватель:</b> ${result.Teacher}\n`
        + `<b>Предмет:</b> ${result.Subject}\n`
        + `<b>Дата:</b> ${result.Date}\n` 
        + `<b>Время:</b> ${result.Time}\n` 
        + `<b>Аудитория:</b> ${result.Room}`;

      if(result) {
        ctx.replyWithHTML(message)
      }
    })
  } catch (err) {
    console.error(err)
  }
}