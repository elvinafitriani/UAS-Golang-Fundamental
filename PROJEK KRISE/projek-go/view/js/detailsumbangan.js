const btn = document.querySelector("#button2");
const inp = document.querySelector(".input-keyword");

btn.addEventListener("click", function () {
  fetch(`http://localhost:5011/detailsumbangan/${inp.value}`)
    .then((Response) => Response.json())
    .then((Response) => {
      const out = document.querySelector(".result");
      console.log(Response);
      console.log(inp.value);

      const buff = document.createElement("thead");
      buff.className = "detail";
      out.appendChild(buff);

      for (let index in Response) {
        const tr = document.createElement("th");
        tr.innerText = "ID";
        buff.appendChild(tr);

        const trow = document.createElement("td");
        trow.textContent = Response[index].id_sumbangan;
        tr.appendChild(trow);

        const tr2 = document.createElement("th");
        tr2.innerText = "Tanggal";
        buff.appendChild(tr2);

        const trow2 = document.createElement("td");
        trow2.textContent = Response[index].tanggal;
        tr2.appendChild(trow2);

        const tr3 = document.createElement("th");
        tr3.innerText = "Nama Donatur";
        buff.appendChild(tr3);

        const trow3 = document.createElement("td");
        trow3.textContent = Response[index].nama_donatur;
        tr3.appendChild(trow3);

        const tr5 = document.createElement("th");
        tr5.innerText = "Jumlah";
        buff.appendChild(tr5);

        const trow5 = document.createElement("td");
        trow5.textContent = Response[index].jumlah;
        tr5.appendChild(trow5);
      }
    });
  inp.value = null;
});
