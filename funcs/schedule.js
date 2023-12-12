import axios from "axios";
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
// const getCurTime = `${getHours}:${getMinutes}`;
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
        return item.Day.toLowerCase() === curDay.toLowerCase() && item.Group === userGroup
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

        for(let i = 0; i <= timeList.length; i++){
          const entry = timeList[i];

          const mStart = entry.start.h * 60 + entry.start.m;
          const mEnd = entry.end.h * 60+ entry.end.m;

          if(curTime >= mStart && curTime <= mEnd) {
            curLesson = i + 1;
            break;
          }
        }

        if(curDay !== "Суббота" || curDay !== "Воскресенье"){
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