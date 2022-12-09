const elemenOutput = document.querySelector(".isi");
view();

function view(){
    fetch("http://localhost:5006/getDosen")
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
            j2.textContent = json[a].nip;
            j1.appendChild(j2);

            let j3 = document.createElement("td");
            j3.textContent = json[a].nama_dosen;
            j1.appendChild(j3);

            let j4 = document.createElement("td");
            j4.textContent = json[a].alamat_dosen;
            j1.appendChild(j4);
            
            let j5 = document.createElement("td");
            j5.textContent = json[a].telp;
            j1.appendChild(j5);

        }

    });
}