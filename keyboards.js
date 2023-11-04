import Markup from "telegraf";

export function getMainMenu() {
  return Markup.keyboard([
    ['Меню' , 'Посмотреть расписания'],
    ["Изменение роли"]
  ]).resize().extra()
}