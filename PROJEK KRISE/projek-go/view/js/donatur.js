const Output2 = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getdonatur")
    .then(function (response) {
      return response.json();
    })
    .then(function (json) {
      console.log(json);
      Output2.innerHTML = "";
      for (let b in json) {
        let trows2 = document.createElement("tr");
        Output2.appendChild(trows2);

        let tdata2 = document.createElement("td");
        tdata2.className = "tdata";
        tdata2.textContent = json[b].idDonatur;
        trows2.appendChild(tdata2);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].namaDonatur;
        trows2.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[b].pekerjaanDonatur;
        trows2.appendChild(tdata4);

        let tdata5 = document.createElement("td");
        tdata5.textContent = json[b].telpDonatur;
        trows2.appendChild(tdata5);
      }
    });
}
