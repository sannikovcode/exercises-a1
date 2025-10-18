function addTask() {
  const input = document.getElementById("taskInput");
  const value = input.value.trim();
  if (!value) return;

  const li = document.createElement("li");
  li.textContent = value;
  li.onclick = () => li.classList.toggle("done");

  document.getElementById("taskList").appendChild(li);
  input.value = "";
}
