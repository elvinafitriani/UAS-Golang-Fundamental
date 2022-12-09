const elemenOutput = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getanak")
    .then(function (response) {
      return response.json();
    })
    .then(function (json) {
      console.log(json);
      elemenOutput.innerHTML = "";
      for (let c in json) {
        let trows = document.createElement("tr");
        elemenOutput.appendChild(trows);

        let tdata = document.createElement("td");
        tdata.textContent = json[c].noInduk;
        trows.appendChild(tdata);

        let tdata2 = document.createElement("td");
        tdata2.textContent = json[c].namaAnak;
        trows.appendChild(tdata2);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[c].jk_anak;
        trows.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[c].ttlAnak;
        trows.appendChild(tdata4);

        let tdata5 = document.createElement("td");
        tdata5.textContent = json[c].agama;
        trows.appendChild(tdata5);

        let tdata6 = document.createElement("td");
        tdata6.textContent = json[c].golDarah;
        trows.appendChild(tdata6);

        let tdata7 = document.createElement("td");
        tdata7.textContent = json[c].beratBadan;
        trows.appendChild(tdata7);

        let tdata8 = document.createElement("td");
        tdata8.textContent = json[c].tinggiBadan;
        trows.appendChild(tdata8);

        let tdata9 = document.createElement("td");
        tdata9.textContent = json[c].tglMasuk;
        trows.appendChild(tdata9);

        let tdata10 = document.createElement("td");
        tdata10.textContent = json[c].status;
        trows.appendChild(tdata10);

        let tdata11 = document.createElement("td");
        tdata11.textContent = json[c].statusPanti;
        trows.appendChild(tdata11);
      }
    });
}
