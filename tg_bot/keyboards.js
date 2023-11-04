import Markup from "telegraf/markup.js";

export function getMainMenu() {
  return Markup.keyboard([
    ['Меню' , 'Посмотреть расписания'],
    ["Изменение роли"]
  ]).resize().extra()
}