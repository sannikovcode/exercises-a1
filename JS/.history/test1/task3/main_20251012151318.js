let age = +prompt("Ваш возраст:");
let hasSubscription = confirm("У вас есть подписка?");

if (age >= 18 && hasSubscription) {
  console.log("Доступ к контенту разрешен");
} else {
  console.log("Доступ запрещен");
}
