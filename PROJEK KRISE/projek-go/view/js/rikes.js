const Output = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getrikes")
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
        tdata.textContent = json[b].idKesehatan;
        trows2.appendChild(tdata);

        let tdata2 = document.createElement("td");
        tdata2.textContent = json[b].no_Induk;
        trows2.appendChild(tdata2);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].tglPeriksa;
        trows2.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[b].hasilPemeriksaaan;
        trows2.appendChild(tdata4);
      }
    });
}
