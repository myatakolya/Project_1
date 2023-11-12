import Scenes from 'telegraf'

export class ScenesGenerator {

  getUserNameScene() {
    const name = new Scenes.WizardScene('name');
    name.enter(async ctx => {
      await ctx.reply('Добро пожаловать в сцену, введите ваше имя')
    })
    name.on('text', async ctx => {
      const userFirstName = ctx.message.text;
      if ( name ) {
        await ctx.reply('Отлично, тебя зовут ' + userFirstName)
        name.scene.enter('age')
      } else {
        await ctx.reply('ИМЯЯЯЯЯЯЯ')
        ctx.reenter()
      }
    })
    return name;
  }

  getUserAgeScene() {
    const age = new Scenes.WizardScene('age');
    age.enter(async ctx => {
      await ctx.reply('Добро пожаловать в сцену, введите ваш возраст')
    })
    age.on('text', async ctx => {
      const userAge = Number(ctx.mesasge.text);
      if (userAge && userAge > 0) {
        await ctx.reply('Ваш возраст' + userAge)
        console.log(ctx.scene)
        await ctx.scene.leave()
      } else { 
        await ctx.reply('Введите возраст')
        ctx.reenter()
      }
    })
    return age;
  }
}