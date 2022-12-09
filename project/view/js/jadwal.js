const elemenOutput = document.querySelector(".isi");
view();

function view(){
    fetch("http://localhost:5006/viewJadwal")
    .then(function(response){
        return response.json();
    })
    .then(function(json){
        console.log(json);
        elemenOutput.innerHTML = "";
        for(let a in json){
            let j1 = document.createElement("tr");
            elemenOutput.appendChild(j1);

            let j2 = document.createElement("td");
            j2.textContent = json[a].id_jadwal;
            j1.appendChild(j2);

            let j3 = document.createElement("td");
            j3.textContent = json[a].nama_dosen;
            j1.appendChild(j3);

            let j4 = document.createElement("td");
            j4.textContent = json[a].nama_matkul;
            j1.appendChild(j4);
            
            let j5 = document.createElement("td");
            j5.textContent = json[a].ruangan;
            j1.appendChild(j5);

            let j6 = document.createElement("td");
            j6.textContent = json[a].hari;
            j1.appendChild(j6);

            let j7 = document.createElement("td");
            j7.textContent = json[a].semester;
            j1.appendChild(j7);
        }

    });
}