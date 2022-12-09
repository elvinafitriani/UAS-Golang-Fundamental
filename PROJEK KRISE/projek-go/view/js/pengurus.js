const Output = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getpengurus")
    .then(function (response) {
      return response.json();
    })
    .then(function (json) {
      console.log(json);
      Output.innerHTML = "";

      for (let b in json) {
        let trows2 = document.createElement("tr");
        Output.appendChild(trows2);

        let tdata = document.createElement("td");
        tdata.textContent = json[b].idKepengurusan;
        trows2.appendChild(tdata);

        let tdata9 = document.createElement("td");
        let tdata10 = document.createElement("img");
        tdata9.appendChild(tdata10);
        tdata10.src = json[b].foto;
        trows2.appendChild(tdata10);

        let tdata2 = document.createElement("td");
        tdata2.textContent = json[b].namaPengurus;
        trows2.appendChild(tdata2);

        let tdata7 = document.createElement("td");
        tdata7.textContent = json[b].jabatan;
        trows2.appendChild(tdata7);

        let tdata8 = document.createElement("td");
        tdata8.textContent = json[b].tglMasuk;
        trows2.appendChild(tdata8);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].jkPengurus;
        trows2.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[b].ttlPengurus;
        trows2.appendChild(tdata4);

        let tdata5 = document.createElement("td");
        tdata5.textContent = json[b].alamatPengurus;
        trows2.appendChild(tdata5);

        let tdata6 = document.createElement("td");
        tdata6.textContent = json[b].telpPengurus;
        trows2.appendChild(tdata6);
      }
    });
}
