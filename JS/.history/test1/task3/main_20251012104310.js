let userText = prompt("Введите ваш текст:");

let mood = prompt(
  "Выберите настроение:\n1 - КРИЧАТЬ\n2 - шептать\n3 - ЗдОрОвО"
);

// 3. Обрабатываем выбор
if (mood === "1") {
  // Используйте .toUpperCase() для переменной userText
  let result = userText.toUpperCase();
  alert(result);
} else if (mood === "2") {
  // Используйте .toLowerCase() для переменной userText
  let result = userText.toLowerCase();
  alert(result);
} else if (mood === "3") {
  // Бонусная часть
  alert("Пока не реализовано!");
} else {
  alert("Неизвестное настроение!");
}
