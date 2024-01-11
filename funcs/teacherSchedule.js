import axios from "axios"
import * as teacher from './getTeacherInfo.js'
const pathToSchedule = '/schedule';
const pathToTeacherSchedule = '/teachers_schedule'

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

export async function getTeacherSchedule(ctx, userDay, teacherName, teacherSurname){
  try {
    axios({
      method: 'get',
      url: `${pathToTeacherSchedule}`
    }).then( async response => {
      await teacher.getTeacherName(ctx).then( response => {
        teacherName = response;
      }).catch(err => {
        console.error(err)
      })
      await teacher.getTeacherSurname(ctx).then( response => {
        teacherSurname = response;
      }).catch(err => {
        console.error(err)
      })
      
      let day = days.findIndex( item => item === userDay )
      if(userDay === 'сегодня') day = getCurDay -1;
      if(userDay === 'завтра') day = getCurDay;
      let lessons = [];
      let lesson = [];

      let result = await response.data.find( item => {
        return item.Name?.toLowerCase() === teacherName?.toLowerCase() && item.Surname?.toLowerCase() === teacherSurname?.toLowerCase()
      })

      

      if (result) {
        if(day >= 0 && day != 6) {
          for(let i in result.Lessons) {
            lessons.push(result.Lessons[i])
          }

          for(let i = 0; i <= lessons.length; i++) {
            if(i === day) {
              for(let k in lessons[i]) {
                lesson.push(lessons[i][k])
              }
            }
          }

          let lessonInfo = lesson.map(item => {
            let message;

            if(result.Subject) {
              message =`<b>Пара</b>: ${result.Subject}\n` + 
              `<b>Аудитория</b>: ${item.Room}\n` + 
              `<b>Тип</b>: ${item.Type}\n` +
              `<b>Время</b>: ${item.Time}\n` +
              `<b>Группа</b>: ${item.Group}`
            }

            return message;
          })

          ctx.replyWithHTML(lessonInfo.join('\n\n'))
        } else if (day < 0  || day === 6) {
          ctx.reply("Сегодня выходной")
        } else {
          ctx.reply("Что-то не так")
        }
      } else {
        await ctx.reply("Не удаётся получить данные")
      }
    })
  } catch (err) {
    console.error(err)
  }
}

export async function getTeacherNextLesson(ctx, teacherName, teacherSurname, curDay = getCurDay - 1){
  try {
    axios({
      method: 'get',
      url: `${pathToTeacherSchedule}`
    }).then( async response => {
      if(!teacherName) {
        await teacher.getTeacherName(ctx).then( response => {
          teacherName = response;
        }).catch(err => {
          console.error(err)
        })
      }
      if(!teacherSurname) {
        await teacher.getTeacherSurname(ctx).then( response => {
          teacherSurname = response;
        }).catch(err => {
          console.error(err)
        })
      }

      let lessons = [];
      let lesson = [];

      let result = await response.data.find( item => {
        return item.Name?.toLowerCase() === teacherName?.toLowerCase() && item.Surname?.toLowerCase() === teacherSurname?.toLowerCase()
      })
      
      if (result) {
        if(curDay >= 0 && curDay != 6) {
          for(let i in result.Lessons) {
            lessons.push(result.Lessons[i])
          }

          for(let i = 0; i <= lessons.length; i++) {
            if(i === curDay) {
              for(let k in lessons[i]) {
                lesson.push(lessons[i][k])
              }
            }
          }

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

          let lessonInfo = lesson.map(item => {
            let message;

            if(result.Subject) {
              message =`<b>Пара</b>: ${result.Subject}\n` + 
              `<b>Аудитория</b>: ${item.Room}\n` + 
              `<b>Тип</b>: ${item.Type}\n` +
              `<b>Время</b>: ${item.Time}\n` +
              `<b>Группа</b>: ${item.Group}`
            }

            return message;
          })

          if(lessonInfo[curLesson]){
            await ctx.replyWithHTML(lessonInfo[curLesson])
          } else if (!lessonInfo[curLesson] && lessonInfo[curLesson + 1]) {
            await ctx.reply("Окно")
          } else {
            await ctx.reply("Пары кончились")
          }
        } else if (curDay < 0  || curDay === 6) {
          ctx.reply("Сегодня выходной")
        } else {
          ctx.reply("Что-то не так")
        }
      } else {
        await ctx.reply("Не удаётся получить данные")
      }
    })
  } catch (err) {
    console.error(err)
  }
}
export async function getStudents(ctx, studentsGroup, curDay = days[getCurDay]) {
  try {
    axios({
      method: 'get',
      url: `${pathToSchedule}`
    }).then( async response => {
      const result = await response.data.find( item => {
        return item.Group === studentsGroup && item.Day?.toLowerCase() === curDay?.toLowerCase()
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
          console.log(entry)

          const mStart = entry.start.h * 60 + entry.start.m;
          const mEnd = entry.end.h * 60+ entry.end.m;

          if(curTime >= mStart && curTime <= mEnd) {
            curLesson = i - 1;
            break;
          }
        }

        if(lessonInfo[curLesson]){
          await ctx.replyWithHTML(lessonInfo[curLesson])
        } else if (!lessonInfo[curLesson] && lessonInfo[curLesson + 1]) {
          await ctx.reply("У них окно")
        } else if (!lessonInfo[curLesson] && !lessonInfo[curLesson + 1]) {
          await ctx.replyWithHTML(lessonInfo.find(item => {return item === true}))
        } else {
          await ctx.reply("Скоро будут дома, если ещё не")
        }
      } else {
        ctx.reply("Сегодня выходной")
      }
    })
  } catch (err) {
    console.error(err)
  }
}