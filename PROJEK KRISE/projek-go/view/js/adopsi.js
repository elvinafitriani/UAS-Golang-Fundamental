const Output = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getadopsi")
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
        tdata.textContent = json[b].idAdopsi;
        trows2.appendChild(tdata);

        let tdata2 = document.createElement("td");
        tdata2.textContent = json[b].namaPengadopsi;
        trows2.appendChild(tdata2);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].umur;
        trows2.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[b].pekerjaan;
        trows2.appendChild(tdata4);

        let tdata5 = document.createElement("td");
        tdata5.textContent = json[b].penghasilan;
        trows2.appendChild(tdata5);

        let tdata9 = document.createElement("td");
        tdata9.textContent = json[b].alamat;
        trows2.appendChild(tdata9);

        let tdata6 = document.createElement("td");
        tdata6.textContent = json[b].statusPerkawinan;
        trows2.appendChild(tdata6);

        let tdata7 = document.createElement("td");
        tdata7.textContent = json[b].nooInduk;
        trows2.appendChild(tdata7);

        let tdata8 = document.createElement("td");
        tdata8.textContent = json[b].tglAdopsi;
        trows2.appendChild(tdata8);
      }
    });
}
