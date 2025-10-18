var students = ["Anna", "Bob", "Charlie", "Diana"];
var searchName = "Charlie";
var found = false;

for (var i = 0; i < students.length; i++) {
  if (students[i] === searchName) {
    found = true;
    break;
  }
}

if (found) {
  console.log(searchName + " найден в списке");
} else {
  console.log(searchName + "не найден");
}
