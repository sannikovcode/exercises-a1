const table = document.querySelector("table");

table.addEventListener("mouseover", function (e) {
  if (e.target.tagName === "TD" || e.target.tagName === "TH") {
    const cell = e.target;
    const row = cell.parentNode;
    const table = cell.closest("table");
    const cellIndex = cell.cellIndex;
    const rowIndex = row.rowIndex;

    clearHighlight(table);

    // Подсвечиваем столбец
    for (let r of table.rows) {
      const c = r.cells[cellIndex];
      if (c) c.style.backgroundColor = "#ffd";
    }

    // Подсвечиваем строку
    for (let c of row.cells) {
      c.style.backgroundColor = "#ffd";
    }

    // Особая подсветка для выделенной ячейки
    cell.style.backgroundColor = "#ffa";
  }
});

table.addEventListener("mouseout", function (e) {
  if (e.target.tagName === "TD" || e.target.tagName === "TH") {
    clearHighlight(table);
  }
});

function clearHighlight(table) {
  for (let r of table.rows) {
    for (let c of r.cells) {
      c.style.backgroundColor = "";
    }
  }
}
