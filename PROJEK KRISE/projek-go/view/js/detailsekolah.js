const btn = document.querySelector("#button2");
const inp = document.querySelector(".input-keyword");

btn.addEventListener("click", function () {
  fetch(`http://localhost:5011/detailsekolah/${inp.value}`)
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
        tr.innerText = "ID ";
        buff.appendChild(tr);

        const trow = document.createElement("td");
        trow.textContent = Response[index].id_sekolah;
        tr.appendChild(trow);

        const tr3 = document.createElement("th");
        tr3.innerText = "Nama Siswa";
        buff.appendChild(tr3);

        const trow2 = document.createElement("td");
        trow2.textContent = Response[index].nama_anak;
        tr3.appendChild(trow2);

        const tr4 = document.createElement("th");
        tr4.innerText = "Kelas ";
        buff.appendChild(tr4);

        const trow3 = document.createElement("td");
        trow3.textContent = Response[index].kelas;
        tr4.appendChild(trow3);

        const tr2 = document.createElement("th");
        tr2.innerText = "Nama Sekolah ";
        buff.appendChild(tr2);

        const trow4 = document.createElement("td");
        trow4.textContent = Response[index].nama_sekolah;
        tr2.appendChild(trow4);

        const tr5 = document.createElement("th");
        tr5.innerText = "Alamat Sekolah ";
        buff.appendChild(tr5);

        const trow5 = document.createElement("td");
        trow5.textContent = Response[index].alamat_sekolah;
        tr5.appendChild(trow5);

        const tr6 = document.createElement("th");
        tr6.innerText = "Telp Sekolah ";
        buff.appendChild(tr6);

        const trow6 = document.createElement("td");
        trow6.textContent = Response[index].telp_sekolah;
        tr6.appendChild(trow6);
      }
    });
  inp.value = null;
});
