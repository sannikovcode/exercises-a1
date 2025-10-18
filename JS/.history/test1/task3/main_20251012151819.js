let age = +prompt("Ваш возраст:");
let isStudent = confirm("Вы студент?");
let totalAmount = +prompt("Сумма покупки:");

if (age < 18 || isStudent || totalAmount > 1000) {
  console.log("Вам положена скидка!");
} else {
  console.log("Скидка не предусмотрена");
}
