const Output2 = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getsumbangan")
    .then(function (response) {
      return response.json();
    })
    .then(function (json) {
      Output2.innerHTML = "";

      for (let b in json) {
        let trows2 = document.createElement("tr");
        Output2.appendChild(trows2);

        let tdata11 = document.createElement("td");
        tdata11.className = "tdata";
        tdata11.textContent = json[b].idSumbangan;
        trows2.appendChild(tdata11);

        let tdata5 = document.createElement("td");
        tdata5.textContent = json[b].idDonatur;
        trows2.appendChild(tdata5);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].tanggal;
        trows2.appendChild(tdata3);

        let tdata6 = document.createElement("td");
        tdata6.textContent = json[b].jumlah;
        trows2.appendChild(tdata6);
      }
    });
}
