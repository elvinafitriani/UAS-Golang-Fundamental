const Output = document.querySelector("#tbody");
tampil();
function tampil() {
  fetch("http://localhost:5011/getsekolah")
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
        tdata.textContent = json[b].idSekolah;
        trows2.appendChild(tdata);

        let tdata2 = document.createElement("td");
        tdata2.textContent = json[b].namaSekolah;
        trows2.appendChild(tdata2);

        let tdata3 = document.createElement("td");
        tdata3.textContent = json[b].alamatSekolah;
        trows2.appendChild(tdata3);

        let tdata4 = document.createElement("td");
        tdata4.textContent = json[b].telpSekolah;
        trows2.appendChild(tdata4);
      }
    });
}
