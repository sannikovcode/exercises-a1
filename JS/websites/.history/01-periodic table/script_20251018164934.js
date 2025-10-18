const table = document.querySelector("table");

table.addEventListener("mouseover", function (e) {
  if (e.target.tagName === "TD" || e.target.tagName === "TH") {
    const cell = e.target;
    const row = cell.parentNode;
    const table = cell.closest("table");
    const cellIndex = cell.cellIndex;

    clearHighlight(table);

    // Подсветка столбца темным цветом
    for (let r of table.rows) {
      const c = r.cells[cellIndex];
      if (c) c.style.backgroundColor = "#444"; // темно-серый фон
    }

    // Подсветка строки темным цветом
    for (let c of row.cells) {
      c.style.backgroundColor = "#444";
    }

    // Подсветка самой ячейки более светлым цветом
    cell.style.backgroundColor = "#666";
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
